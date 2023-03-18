package lang

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/alter"
	"github.com/lmorg/murex/utils/envvars"
	"github.com/lmorg/murex/utils/json"
)

func errVariableReserved(name string) error {
	return fmt.Errorf("cannot set a reserved variable: %s", name)
}

func errVarNotExist(name string) error {
	return fmt.Errorf("variable '%s' does not exist", name)
}

func errVarNoParam(i int, err error) error {
	return fmt.Errorf("variable '%d' cannot be defined: %s", i, err.Error())
}

var (
	errZeroLengthPath = errors.New("zero length path")
)

// Reserved variable names. Set as constants so any typos of these names within
// the code will be raised as compiler errors
const (
	DOT        = ""
	SELF       = "SELF"
	ARGS       = "ARGS"
	PARAMS     = "PARAMS"
	MUREX_EXE  = "MUREX_EXE"
	MUREX_ARGS = "MUREX_ARGS"
	HOSTNAME   = "HOSTNAME"
	PWD        = "PWD"
	ENV        = "ENV"
	GLOBAL     = "GLOBAL"
)

// Variables is a table of all the variables. This will be local to the scope's
// process
type Variables struct {
	process *Process // only needed for variables
	vars    map[string]*variable
	mutex   sync.Mutex
	global  bool
}

// NewVariables creates a new variable table
func NewVariables(p *Process) *Variables {
	v := new(Variables)
	v.vars = make(map[string]*variable)
	v.process = p
	return v
}

// NewGlobals creates a new global variable table
func NewGlobals() *Variables {
	v := new(Variables)
	v.vars = make(map[string]*variable)
	v.process = ShellProcess
	v.global = true
	return v
}

// variable is an individual variable or global variable
type variable struct {
	DataType    string
	Value       interface{}
	String      string
	IsInterface bool
	Modify      time.Time
	FileRef     *ref.File // only needed for globals
}

// GetValue return the value of a variable. If a variable does not exist then
// GetValue will return nil. Please check if p.Config.Get("proc", "strict-vars", "bool")
// matters for your usage of GetValue because this API doesn't care. If in doubt
// use GetString instead.
func (v *Variables) GetValue(path string) (interface{}, error) {
	if path == "." {
		return v.getValue(DOT)
	}

	split := strings.Split(path, ".")
	switch len(split) {
	case 0:
		return nil, errors.New("zero length path")
	case 1:
		return v.getValue(split[0])
	default:
		val, err := v.getValue(split[0])
		if err != nil {
			return nil, err
		}

		return ElementLookup(val, "."+strings.Join(split[1:], "."))
	}
}

/*func (v *Variables) isObject(name string) bool {
	if v.global {
		isObject := v.isObjectValue(name)
		if isObject == nil {
			return false
		}
		return isObject.(bool)
	}

	isObject := v.isObjectValue(name)
	if isObject != nil {
		return isObject.(bool)
	}

	isObject = GlobalVariables.getValueValue(name)
	if isObject == nil {
		return false
	}
	return isObject.(bool)
}

// Return values:
// * true:  var exists and is object
// * false: var exists and not an object
// * nil:   var does not exist
func (v *Variables) isObjectValue(name string) interface{} {
	v.mutex.Lock()
	variable := v.vars[name]
	if variable == nil {
		v.mutex.Unlock()
		return nil
	}

	isObject := variable.IsObject
	v.mutex.Unlock()

	return isObject
}*/

func (v *Variables) getValue(name string) (interface{}, error) {
	switch name {
	case ENV:
		return getEnvVarValue(), nil

	case GLOBAL:
		return getGlobalValues(), nil

	case SELF:
		return getVarSelf(v.process), nil

	case ARGS:
		return getVarArgs(v.process), nil

	case PARAMS:
		return v.process.Scope.Parameters.StringArray(), nil

	case MUREX_EXE:
		return getVarMurexExeValue()

	case HOSTNAME:
		return getHostname(), nil

	case PWD:
		return getPwdValue()

	case "0":
		return v.process.Scope.Name.String(), nil
	}

	if i, err := strconv.Atoi(name); err == nil && i > 0 {
		s, err := v.process.Scope.Parameters.String(i - 1)
		if err != nil {
			return nil, nil
		}
		return s, nil
	}

	if v.global {
		return v.getValueValue(name), nil
	}

	value := v.getValueValue(name)
	if value != nil {
		return value, nil
	}

	value = GlobalVariables.getValueValue(name)
	if value != nil {
		return value, nil
	}

	// variable not found so lets fallback to the environmental variables
	value = os.Getenv(name)
	if value != "" {
		return value, nil
	}

	strictVars, err := v.process.Config.Get("proc", "strict-vars", "bool")
	if err != nil || strictVars.(bool) {
		return nil, errVarNotExist(name)
	}
	return nil, nil
}

func (v *Variables) getValueValue(name string) interface{} {
	v.mutex.Lock()
	variable := v.vars[name]
	if variable == nil {
		v.mutex.Unlock()
		return nil
	}

	if variable.IsInterface {
		value := variable.Value.(MxInterface).GetValue()

		v.mutex.Unlock()
		return value
	}

	value := variable.Value

	v.mutex.Unlock()
	return value
}

// GetString returns a string representation of the data stored in the requested variable
func (v *Variables) GetString(path string) (string, error) {
	if path == "." {
		return v.getString(DOT)
	}

	split := strings.Split(path, ".")
	switch len(split) {
	case 0:
		return "", errZeroLengthPath
	case 1:
		return v.getString(split[0])
	default:
		val, err := v.getValue(split[0])
		if err != nil {
			return "", err
		}

		val, err = ElementLookup(val, "."+strings.Join(split[1:], "."))
		if err != nil {
			return "", err
		}

		switch val.(type) {
		case int, float64, string, bool, nil:
			s, err := types.ConvertGoType(val, types.String)
			if err != nil {
				return "", err
			}
			return s.(string), nil
		default:
			dataType := v.GetDataType(split[0])
			b, err := MarshalData(v.process, dataType, val)
			return string(b), err
		}
	}
}

func (v *Variables) getString(name string) (string, error) {
	switch name {
	case ENV:
		b, err := json.Marshal(getEnvVarValue(), v.process.Stdout.IsTTY())
		return string(b), err

	case GLOBAL:
		b, err := json.Marshal(getGlobalValues(), v.process.Stdout.IsTTY())
		return string(b), err

	case SELF:
		b, err := json.Marshal(getVarSelf(v.process), v.process.Stdout.IsTTY())
		return string(b), err

	case ARGS:
		b, err := json.Marshal(getVarArgs(v.process), v.process.Stdout.IsTTY())
		return string(b), err

	case PARAMS:
		b, err := json.Marshal(v.process.Scope.Parameters.StringArray(), v.process.Stdout.IsTTY())
		return string(b), err

	case MUREX_EXE:
		return os.Executable()

	case HOSTNAME:
		return getHostname(), nil

	case PWD:
		return os.Getwd()

	case "0":
		return v.process.Scope.Name.String(), nil
	}

	if i, err := strconv.Atoi(name); err == nil && i > 0 {
		s, err := v.process.Scope.Parameters.String(i - 1)
		if err != nil {
			return "", errVarNoParam(i, err)
		}
		return s, nil
	}

	if v.global {
		val, _ := v.getStringValue(name)
		return val, nil
	}

	s, exists := v.getStringValue(name)
	if exists {
		return s, nil
	}

	s, exists = GlobalVariables.getStringValue(name)
	if exists {
		return s, nil
	}

	// variable not found so lets fallback to the environmental variables
	s, exists = os.LookupEnv(name)

	strictVars, err := v.process.Config.Get("proc", "strict-vars", "bool")
	if (err != nil || strictVars.(bool)) && !exists {
		return "", errVarNotExist(name)
	}

	return s, nil
}

func (v *Variables) getStringValue(name string) (string, bool) {
	v.mutex.Lock()
	variable := v.vars[name]
	if variable == nil {
		v.mutex.Unlock()
		return "", false
	}

	if variable.IsInterface {
		s := variable.Value.(MxInterface).GetString()

		v.mutex.Unlock()
		return s, true
	}

	s := variable.String
	v.mutex.Unlock()
	return s, true
}

// GetDataType returns the data type of the variable stored in the referenced VarTable
func (v *Variables) GetDataType(path string) string {
	if path == "." {
		return v.getDataType(DOT)
	}

	split := strings.Split(path, ".")
	switch len(split) {
	case 0:
		return ""
	case 1:
		return v.getDataType(split[0])
	default:
		val, err := v.getValue(split[0])
		if err != nil {
			return v.getDataType(split[0])
		}

		val, err = ElementLookup(val, "."+strings.Join(split[1:], "."))
		if err != nil {
			return v.getDataType(split[0])
		}

		switch val.(type) {
		case int:
			return types.Integer
		case float64:
			return types.Number
		case string, []byte, []rune:
			return types.String
		case bool:
			return types.Boolean
		case nil:
			return types.Null
		default:
			return v.getDataType(split[0])
		}
	}
}

func (v *Variables) getDataType(name string) string {
	switch name {
	case ENV:
		return types.Json

	case GLOBAL:
		return types.Json

	case SELF:
		return types.Json

	case ARGS:
		return types.Json

	case PARAMS:
		return types.Json

	case MUREX_EXE:
		return types.Path

	case PWD:
		return types.Path

	case "0":
		return types.String
	}

	if i, err := strconv.Atoi(name); err == nil && i > 0 {
		if i >= v.process.Scope.Parameters.Len() {
			return ""
		}
		return types.String
	}

	if v.global {
		dt, _ := v.getDataTypeValue(name)
		return dt
	}

	s, exists := v.getDataTypeValue(name)
	if exists {
		return s
	}

	s, exists = GlobalVariables.getDataTypeValue(name)
	if exists {
		return s
	}

	// variable not found so lets fallback to the environmental variables
	value := os.Getenv(name)
	if value != "" {
		return types.String
	}

	return types.Null
}

func (v *Variables) getDataTypeValue(name string) (string, bool) {
	v.mutex.Lock()
	variable := v.vars[name]
	if variable == nil {
		v.mutex.Unlock()
		return "", false
	}

	dt := variable.DataType
	v.mutex.Unlock()
	return dt, true
}

func errCannotUpdateNested(name string, err error) error {
	return fmt.Errorf("cannot update element inside %s: %s", name, err.Error())
}

func (v *Variables) Set(p *Process, path string, value interface{}, dataType string) error {
	split := strings.Split(path, ".")
	switch len(split) {
	case 0:
		return errZeroLengthPath
	case 1:
		return v.set(p, split[0], value, dataType, nil)
	default:
		variable, err := v.getValue(split[0])
		if err != nil {
			return errCannotUpdateNested(split[0], err)
		}

		variable, err = alter.Alter(p.Context, variable, split[1:], value)
		if err != nil {
			return errCannotUpdateNested(split[0], err)
		}
		err = v.set(p, split[0], variable, v.GetDataType(split[0]), split[1:])
		if err != nil {
			return errCannotUpdateNested(split[0], err)
		}
		return nil
	}
}

// Set writes a variable
func (v *Variables) set(p *Process, name string, value interface{}, dataType string, changePath []string) error {
	switch name {
	case SELF, ARGS, PARAMS, MUREX_EXE, MUREX_ARGS, HOSTNAME, PWD, "_":
		return errVariableReserved(name)
	case ENV:
		return setEnvVar(value, changePath)
	case DOT:
		goto notReserved
	}
	for _, r := range name {
		if r < '0' || r > '9' {
			goto notReserved
		}
	}
	return errVariableReserved(name)

notReserved:

	fileRef := v.process.FileRef
	if v.global {
		fileRef = p.FileRef
	}

	mxi := MxInterfaces[dataType]
	if mxi != nil {
		mxvar := v.vars[name]
		if mxvar != nil && mxvar.IsInterface {

			v.mutex.Lock()

			err := mxvar.Value.(MxInterface).Set(value, changePath)
			if err != nil {
				v.vars[name].Modify = time.Now()
			}

			v.mutex.Unlock()

			return err
		}

		s, _, err := convertDataType(value, dataType)
		if err != nil {
			return err
		}

		mxi, err := mxi.New(s)
		if err != nil {
			return err
		}

		v.mutex.Lock()

		v.vars[name] = &variable{
			Value:       mxi,
			DataType:    dataType,
			Modify:      time.Now(),
			FileRef:     fileRef,
			IsInterface: true,
		}

		v.mutex.Unlock()

		return nil
	}

	s, iface, err := convertDataType(value, dataType)
	if err != nil {
		return err
	}

	v.mutex.Lock()

	v.vars[name] = &variable{
		Value:    iface,
		String:   s,
		DataType: dataType,
		Modify:   time.Now(),
		FileRef:  fileRef,
	}

	v.mutex.Unlock()

	return nil
}

func setEnvVar(v interface{}, changePath []string) (err error) {
	var value interface{}

	if len(changePath) == 0 {
		return fmt.Errorf("invalid use of $%s. Expecting an environmental variable name, eg `$ENV.EXAMPLE`", ENV)
	}

	switch t := v.(type) {
	case map[string]interface{}:
		value, err = types.ConvertGoType(t[changePath[0]], types.String)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("expecting a map of environmental variables. Instead got a %T", t)
	}

	return os.Setenv(changePath[0], value.(string))
}

const errCannotStoreVariable = "cannot store variable"

func convertDataType(value interface{}, dataType string) (string, interface{}, error) {
	var (
		s     string
		iface interface{}
		err   error
	)

	switch v := value.(type) {
	case float64, int, bool, nil:
		s, err = varConvertPrimitive(value)
		iface = value
	case string:
		s = v
		if dataType != types.String && dataType != types.Generic {
			iface, err = varConvertString([]byte(v), dataType)
		} else {
			iface = s
		}
	case []byte:
		s = string(v)
		if dataType != types.String && dataType != types.Generic {
			iface, err = varConvertString(v, dataType)
		} else {
			iface = s
		}
	case []rune:
		s = string(v)
		if dataType != types.String && dataType != types.Generic {
			iface, err = varConvertString([]byte(string(v)), dataType)
		} else {
			iface = s
		}
	default:
		s, err = varConvertInterface(v, dataType)
		iface = value
	}
	return s, iface, err
}

func varConvertPrimitive(value interface{}) (string, error) {
	s, err := types.ConvertGoType(value, types.String)
	if err != nil {
		return "", fmt.Errorf("%s: %s", errCannotStoreVariable, err.Error())
	}
	return s.(string), nil
}

func varConvertString(value []byte, dataType string) (interface{}, error) {
	UnmarshalData := Unmarshallers[dataType]

	// no unmarshaller exists so lets just return the bare string
	if UnmarshalData == nil {
		return string(value), nil
	}

	p := new(Process)
	p.Stdin = streams.NewStdin()
	_, err := p.Stdin.Write([]byte(value))
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errCannotStoreVariable, err.Error())
	}
	v, err := UnmarshalData(p)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errCannotStoreVariable, err.Error())
	}
	return v, nil
}

func varConvertInterface(value interface{}, dataType string) (string, error) {
	MarshalData := Marshallers[dataType]

	// no marshaller exists so lets just return the bare string
	if MarshalData == nil {
		s, err := types.ConvertGoType(value, types.String)
		if err != nil {
			return "", fmt.Errorf("%s: %s", errCannotStoreVariable, err.Error())
		}
		return s.(string), nil
	}

	b, err := MarshalData(ShellProcess, value)
	if err != nil {
		return "", fmt.Errorf("%s: %s", errCannotStoreVariable, err.Error())
	}
	s, err := types.ConvertGoType(b, types.String)
	if err != nil {
		return "", fmt.Errorf("%s: %s", errCannotStoreVariable, err.Error())
	}
	return s.(string), nil
}

// Unset removes a variable from the table
func (v *Variables) Unset(name string) error {
	v.mutex.Lock()
	variable := v.vars[name]
	if variable == nil {
		v.mutex.Unlock()
		return errVarNotExist(name)
	}

	delete(v.vars, name)
	v.mutex.Unlock()
	return nil
}

// Dump returns a map of the structure of all variables in scope
func (v *Variables) Dump() interface{} {
	v.mutex.Lock()
	vars := v.vars // TODO: This isn't thread safe
	v.mutex.Unlock()

	return vars
}

// DumpVariables returns a map of the variables and values for all variables
// in scope.
func DumpVariables(p *Process) map[string]interface{} {
	m := make(map[string]interface{})

	envvars.All(m)

	GlobalVariables.mutex.Lock()
	for name, v := range GlobalVariables.vars {
		m[name] = v.Value
	}
	GlobalVariables.mutex.Unlock()

	p.Variables.mutex.Lock()
	for name, v := range p.Variables.vars {
		m[name] = v.Value
	}
	p.Variables.mutex.Unlock()

	m[SELF], _ = p.Variables.GetValue(SELF)
	m[ARGS], _ = p.Variables.GetValue(ARGS)
	m[PARAMS], _ = p.Variables.GetValue(PARAMS)
	m[MUREX_EXE], _ = p.Variables.GetValue(MUREX_EXE)
	m[MUREX_ARGS], _ = p.Variables.GetValue(MUREX_ARGS)
	m[HOSTNAME], _ = p.Variables.GetValue(HOSTNAME)
	m[PWD], _ = p.Variables.GetValue(PWD)
	m[ENV], _ = p.Variables.GetValue(ENV)
	m[GLOBAL] = nil
	return m
}
