package lang

import (
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
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func errVariableReserved(name string) error {
	return fmt.Errorf("cannot set a reserved variable: %s", name)
}

const ErrDoesNotExist = "does not exist"

func errVarNotExist(name string) error {
	return fmt.Errorf("variable '%s' %s", name, ErrDoesNotExist)
}

func errVarCannotUpdateNested(name string, err error) error {
	return fmt.Errorf("cannot update element inside %s: %s", name, err.Error())
}

func errVarCannotUpdateIndexOrElement(name string) error {
	return fmt.Errorf("cannot update `[ indexes ]` nor `[[ elements ]]`, these are immutable objects.\nPlease reference values using dot notation instead, eg $variable_name.path.to.element\nVariable  : %s", name)
}

func errVarNoParam(i int, err error) error {
	return fmt.Errorf("variable '%d' is not set because the scope returned the following error when querying parameter %d: %s", i, i, err.Error())
}

func errVarZeroLengthPath(name string) error {
	return fmt.Errorf("zero length path in variable name `%s`", name)
}

func errVarCannotGetProperty(name, path string, err error) error {
	return fmt.Errorf("cannot get property '%s' for variable '%s'.\n%s.\nif this wasn't intended to be a property then try surrounding your variable name with parenthesis, eg:\n`  ...  $(%s).%s  ...  `",
		path, name, err.Error(), name, path)
}

func errVarCannotStoreVariable(name string, err error) error {
	return fmt.Errorf("cannot store variable '%s': %s", name, err.Error())
}

// Reserved variable names. Set as constants so any typos of these names within
// the code will be raised as compiler errors
const (
	_VAR_DOT        = ""
	_VAR_SELF       = "SELF"       // Scope: function
	_VAR_ARGV       = "ARGV"       // Scope: function
	_VAR_ARGS       = "ARGS"       // Scope: function
	_VAR_PARAMS     = "PARAMS"     // Scope: function
	_VAR_MUREX_EXE  = "MUREX_EXE"  // Scope: global
	_VAR_MUREX_ARGS = "MUREX_ARGS" // Scope: global
	_VAR_MUREX_ARGV = "MUREX_ARGV" // Scope: global
	_VAR_HOSTNAME   = "HOSTNAME"   // Scope: global
	_VAR_MODULE     = "MOD"        // Scope: module
	_VAR_GLOBAL     = "GLOBAL"     // Scope: global
	_VAR_ENV        = "ENV"        // Scope: session (environmental variable)
	_VAR_PWD        = "PWD"        // POSIX
	_VAR_OLDPWD     = "OLDPWD"     // POSIX (not tested)
	_VAR_HOME       = "HOME"       // POSIX (not tested)
	_VAR_TMPDIR     = "TMPDIR"     // POSIX (not tested)
	_VAR_COLUMNS    = "COLUMNS"    // POSIX
	_VAR_LINES      = "LINES"      // POSIX
	_VAR_RANDOM     = "RANDOM"     // POSIX
	_VAR_USER       = "USER"       // POSIX
	_VAR_LOGNAME    = "LOGNAME"    // POSIX
)

var ReservedVariableNames = []string{
	_VAR_SELF, _VAR_ARGV, _VAR_ARGS, _VAR_PARAMS,
	_VAR_MUREX_EXE, _VAR_MUREX_ARGS, _VAR_MUREX_ARGV,
	_VAR_HOSTNAME,
	_VAR_MODULE, _VAR_GLOBAL, _VAR_ENV,
	_VAR_PWD, _VAR_OLDPWD, _VAR_HOME, _VAR_TMPDIR,
	_VAR_COLUMNS, _VAR_LINES, _VAR_RANDOM,
	_VAR_USER, _VAR_LOGNAME,
}

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
		return v.getValue(_VAR_DOT)
	}

	split := strings.Split(path, ".")
	switch len(split) {
	case 0:
		return nil, errVarZeroLengthPath(path)
	case 1:
		return v.getValue(split[0])
	default:
		val, err := v.getValue(split[0])
		if err != nil {
			return nil, err
		}

		propertyPath := strings.Join(split[1:], ".")
		value, err := ElementLookup(val, "."+propertyPath, v.getDataType(split[0]))
		if err != nil {
			err = errVarCannotGetProperty(split[0], propertyPath, err)
		}
		return value, err
	}
}

func (v *Variables) getValue(name string) (interface{}, error) {
	switch name {
	case _VAR_ENV:
		return getEnvVarValue(v), nil

	case _VAR_GLOBAL:
		return getGlobalValues(), nil

	case _VAR_MODULE:
		return ModuleVariables.GetValues(v.process), nil

	case _VAR_SELF:
		return getVarSelf(v.process), nil

	case _VAR_ARGV, _VAR_ARGS:
		return getVarArgs(v.process), nil

	case _VAR_PARAMS:
		return v.process.Scope.Parameters.StringArray(), nil

	case _VAR_MUREX_EXE:
		return getVarMurexExeValue()

	case _VAR_MUREX_ARGV, _VAR_MUREX_ARGS:
		return getVarMurexArgs(), nil

	case _VAR_HOSTNAME:
		return getHostname(), nil

	case _VAR_PWD:
		return getPwdValue()

	case _VAR_OLDPWD:
		return getVarOldPwdValue()

	case _VAR_HOME:
		return home.MyDir, nil

	case _VAR_TMPDIR:
		return getVarTmpDirValue(), nil

	case _VAR_COLUMNS:
		return getVarColumnsValue(), nil

	case _VAR_LINES:
		return getVarLinesValue()

	case _VAR_RANDOM:
		return getVarRandomValue(), nil

	case _VAR_USER, _VAR_LOGNAME:
		return getVarUserNameValue()

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
	s, exists := os.LookupEnv(name)
	if exists {
		return v.getEnvValueValue(name, s)
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

func (v *Variables) getEnvValueValue(name, str string) (interface{}, error) {
	dt := getEnvVarDataType(name)
	if dt == types.String {
		return str, nil
	}

	value, err := UnmarshalDataBuffered(v.process, []byte(str), dt)
	return value, err
}

// GetString returns a string representation of the data stored in the requested variable
func (v *Variables) GetString(path string) (string, error) {
	if path == "." {
		return v.getString(_VAR_DOT)
	}

	split := strings.Split(path, ".")
	switch len(split) {
	case 0:
		return "", errVarZeroLengthPath(path)
	case 1:
		return v.getString(split[0])
	default:
		val, err := v.getValue(split[0])
		if err != nil {
			return "", err
		}

		propertyPath := strings.Join(split[1:], ".")
		val, err = ElementLookup(val, "."+propertyPath, v.getDataType(split[0]))
		if err != nil {
			return "", errVarCannotGetProperty(split[0], propertyPath, err)
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
	case _VAR_ENV:
		b, err := json.Marshal(getEnvVarString(), v.process.Stdout.IsTTY())
		return string(b), err

	case _VAR_GLOBAL:
		b, err := json.Marshal(getGlobalValues(), v.process.Stdout.IsTTY())
		return string(b), err

	case _VAR_MODULE:
		b, err := json.Marshal(ModuleVariables.GetValues(v.process), v.process.Stdout.IsTTY())
		return string(b), err

	case _VAR_SELF:
		b, err := json.Marshal(getVarSelf(v.process), v.process.Stdout.IsTTY())
		return string(b), err

	case _VAR_ARGV, _VAR_ARGS:
		b, err := json.Marshal(getVarArgs(v.process), v.process.Stdout.IsTTY())
		return string(b), err

	case _VAR_PARAMS:
		b, err := json.Marshal(v.process.Scope.Parameters.StringArray(), v.process.Stdout.IsTTY())
		return string(b), err

	case _VAR_MUREX_EXE:
		return os.Executable()

	case _VAR_MUREX_ARGV, _VAR_MUREX_ARGS:
		b, err := json.Marshal(getVarMurexArgs(), v.process.Stdout.IsTTY())
		return string(b), err

	case _VAR_HOSTNAME:
		return getHostname(), nil

	case _VAR_PWD:
		return os.Getwd()

	case _VAR_OLDPWD:
		return getVarOldPwdValue()

	case _VAR_HOME:
		return home.MyDir, nil

	case _VAR_TMPDIR:
		return getVarTmpDirValue(), nil

	case _VAR_COLUMNS:
		return strconv.Itoa(getVarColumnsValue()), nil

	case _VAR_LINES:
		i, err := getVarLinesValue()
		return strconv.Itoa(i), err

	case _VAR_RANDOM:
		return strconv.Itoa(getVarRandomValue()), nil

	case _VAR_USER, _VAR_LOGNAME:
		return getVarUserNameValue()

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
		return v.getDataType(_VAR_DOT)
	}

	split := strings.Split(path, ".")
	switch len(split) {
	case 0:
		return ""
	case 1:
		return v.getDataType(split[0])
	default:
		switch split[0] {
		case _VAR_ENV:
			return getEnvVarDataType(split[1])
		case _VAR_GLOBAL:
			return getGlobalDataType(split[1])
		case _VAR_MODULE:
			return ModuleVariables.GetDataType(v.process, split[1])
		}

		val, err := v.getValue(split[0])
		if err != nil {
			return v.getDataType(split[0])
		}

		val, err = ElementLookup(val, "."+strings.Join(split[1:], "."), v.getDataType(split[0]))
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
	case _VAR_ENV, _VAR_GLOBAL, _VAR_MODULE:
		return types.Json

	case _VAR_SELF:
		return types.Json

	case _VAR_ARGV, _VAR_ARGS:
		return types.Json

	case _VAR_PARAMS:
		return types.Json

	case _VAR_MUREX_EXE:
		return types.Path

	case _VAR_MUREX_ARGV, _VAR_MUREX_ARGS:
		return types.Json

	case _VAR_PWD, _VAR_OLDPWD, _VAR_TMPDIR, _VAR_HOME:
		return types.Path

	case _VAR_COLUMNS, _VAR_LINES, _VAR_RANDOM:
		return types.Integer

	case _VAR_USER, _VAR_LOGNAME:
		return types.String

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
	_, exists = os.LookupEnv(name)
	if exists {
		return getEnvVarDataType(name)
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

func getGlobalDataType(name string) (dt string) {
	GlobalVariables.mutex.Lock()
	v := GlobalVariables.vars[name]
	if v != nil {
		dt = v.DataType
	}
	GlobalVariables.mutex.Unlock()
	return
}

func getEnvVarDataType(name string) string {
	for dt, v := range envDataTypes {
		if lists.Match(v, name) {
			return dt
		}
	}

	return types.String
}

func (v *Variables) Set(p *Process, path string, value interface{}, dataType string) error {
	if strings.Contains(path, "[") {
		return errVarCannotUpdateIndexOrElement(path)
	}

	split := strings.Split(path, ".")
	switch len(split) {
	case 0:
		return errVarZeroLengthPath(path)
	case 1:
		return v.set(p, split[0], value, dataType, nil)
	default:
		variable, err := v.getValue(split[0])
		if err != nil {
			return errVarCannotUpdateNested(split[0], err)
		}

		variable, err = alter.Alter(p.Context, variable, split[1:], value)
		if err != nil {
			return errVarCannotUpdateNested(split[0], err)
		}
		err = v.set(p, split[0], variable, v.getNestedDataType(split[0], dataType), split[1:])
		if err != nil {
			return errVarCannotUpdateNested(split[0], err)
		}
		return nil
	}
}

func (v *Variables) getNestedDataType(name string, dataType string) string {
	switch name {
	case _VAR_GLOBAL, _VAR_MODULE:
		return dataType
	default:
		return v.GetDataType(name)
	}
}

// Set writes a variable
func (v *Variables) set(p *Process, name string, value interface{}, dataType string, changePath []string) error {
	switch name {
	case _VAR_SELF, _VAR_ARGV, _VAR_ARGS, _VAR_PARAMS,
		_VAR_MUREX_EXE, _VAR_MUREX_ARGS, _VAR_MUREX_ARGV,
		_VAR_HOSTNAME, _VAR_PWD, _VAR_OLDPWD, _VAR_HOME, _VAR_TMPDIR,
		_VAR_COLUMNS, _VAR_LINES, _VAR_RANDOM,
		_VAR_USER, _VAR_LOGNAME,
		"_":
		return errVariableReserved(name)
	case _VAR_ENV:
		return setEnvVar(value, changePath)
	case _VAR_GLOBAL:
		return setGlobalVar(value, changePath, dataType)
	case _VAR_MODULE:
		return ModuleVariables.Set(p, value, changePath, dataType)
	case _VAR_DOT:
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

		s, _, err := convertDataType(p, value, dataType, &name)
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

	s, iface, err := convertDataType(p, value, dataType, &name)
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

func setGlobalVar(v interface{}, changePath []string, dataType string) (err error) {
	if len(changePath) == 0 {
		return fmt.Errorf("invalid use of $%s. Expecting a global variable name, eg `$%s.example`", _VAR_GLOBAL, _VAR_GLOBAL)
	}

	switch t := v.(type) {
	case map[string]interface{}:
		return GlobalVariables.Set(ShellProcess, changePath[0], t[changePath[0]], dataType)

	default:
		return fmt.Errorf("expecting a map of global variables. Instead got a %T", t)
	}
}

func setEnvVar(v interface{}, changePath []string) (err error) {
	var value interface{}

	if len(changePath) == 0 {
		return fmt.Errorf("invalid use of $%s. Expecting an environmental variable name, eg `$%s.EXAMPLE`", _VAR_ENV, _VAR_ENV)
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

func convertDataType(p *Process, value interface{}, dataType string, name *string) (string, interface{}, error) {
	var (
		s     string
		iface interface{}
		err   error
	)

	switch v := value.(type) {
	case float64, int, bool, nil:
		s, err = varConvertPrimitive(value, name)
		iface = value
	case string:
		s = v
		if dataType != types.String && dataType != types.Generic {
			iface, err = varConvertString(p, []byte(v), dataType, name)
		} else {
			iface = s
		}
	case []byte:
		s = string(v)
		if dataType != types.String && dataType != types.Generic {
			iface, err = varConvertString(p, v, dataType, name)
		} else {
			iface = s
		}
	case []rune:
		s = string(v)
		if dataType != types.String && dataType != types.Generic {
			iface, err = varConvertString(p, []byte(string(v)), dataType, name)
		} else {
			iface = s
		}
	default:
		s, err = varConvertInterface(p, v, dataType, name)
		iface = value
	}
	return s, iface, err
}

func varConvertPrimitive(value interface{}, name *string) (string, error) {
	s, err := types.ConvertGoType(value, types.String)
	if err != nil {
		return "", errVarCannotStoreVariable(*name, err)
	}
	return s.(string), nil
}

func varConvertString(parent *Process, value []byte, dataType string, name *string) (interface{}, error) {
	UnmarshalData := _unmarshallers[dataType]

	// no unmarshaller exists so lets just return the bare string
	if UnmarshalData == nil {
		return string(value), nil
	}

	p := new(Process)
	p.Config = parent.Config
	p.Stdin = streams.NewStdin()
	_, err := p.Stdin.Write([]byte(value))
	if err != nil {
		return nil, errVarCannotStoreVariable(*name, err)
	}
	v, err := UnmarshalData(p)
	if err != nil {
		return nil, errVarCannotStoreVariable(*name, err)
	}
	return v, nil
}

/*func varConvertString(parent *Process, value []byte, dataType string, name *string) (interface{}, error) {
	// no unmarshaller exists so lets just return the bare string
	if _unmarshallers[dataType] == nil {
		return string(value), nil
	}

	v, err := UnmarshalDataBuffered(parent, value, dataType)
	if err != nil {
		return nil, errVarCannotStoreVariable(*name, err)
	}
	return v, nil
}*/

/*func varConvertInterface(p *Process, value interface{}, dataType string, name *string) (string, error) {
	// no marshaller exists so lets just return the bare string
	if _marshallers[dataType] == nil {
		s, err := types.ConvertGoType(value, types.String)
		if err != nil {
			return "", errVarCannotStoreVariable(*name, err)
		}
		return s.(string), nil
	}

	b, err := MarshalData(p, dataType, value)
	if err != nil {
		return "", errVarCannotStoreVariable(*name, err)
	}
	s, err := types.ConvertGoType(b, types.String)
	if err != nil {
		return "", errVarCannotStoreVariable(*name, err)
	}
	return s.(string), nil
}*/

func varConvertInterface(p *Process, value interface{}, dataType string, name *string) (string, error) {
	MarshalData := _marshallers[dataType]

	// no marshaller exists so lets just return the bare string
	if MarshalData == nil {
		s, err := types.ConvertGoType(value, types.String)
		if err != nil {
			return "", errVarCannotStoreVariable(*name, err)
		}
		return s.(string), nil
	}

	b, err := MarshalData(p, value)
	if err != nil {
		return "", errVarCannotStoreVariable(*name, err)
	}
	s, err := types.ConvertGoType(b, types.String)
	if err != nil {
		return "", errVarCannotStoreVariable(*name, err)
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

	m[_VAR_SELF], _ = p.Variables.GetValue(_VAR_SELF)
	m[_VAR_ARGV], _ = p.Variables.GetValue(_VAR_ARGV)
	m[_VAR_PARAMS], _ = p.Variables.GetValue(_VAR_PARAMS)
	m[_VAR_MUREX_EXE], _ = p.Variables.GetValue(_VAR_MUREX_EXE)
	m[_VAR_MUREX_ARGV], _ = p.Variables.GetValue(_VAR_MUREX_ARGV)
	m[_VAR_HOSTNAME], _ = p.Variables.GetValue(_VAR_HOSTNAME)
	m[_VAR_PWD], _ = p.Variables.GetValue(_VAR_PWD)
	m[_VAR_OLDPWD], _ = p.Variables.GetValue(_VAR_OLDPWD)
	m[_VAR_HOME], _ = p.Variables.GetValue(_VAR_HOME)
	m[_VAR_TMPDIR], _ = p.Variables.GetValue(_VAR_TMPDIR)
	m[_VAR_GLOBAL] = ".."
	m[_VAR_MODULE] = ".."
	m[_VAR_ENV], _ = p.Variables.GetValue(_VAR_ENV)
	m[_VAR_COLUMNS], _ = p.Variables.GetValue(_VAR_COLUMNS)
	m[_VAR_LINES], _ = p.Variables.GetValue(_VAR_LINES)
	m[_VAR_RANDOM], _ = p.Variables.GetValue(_VAR_RANDOM)
	m[_VAR_USER], _ = p.Variables.GetValue(_VAR_USER)
	m[_VAR_LOGNAME], _ = p.Variables.GetValue(_VAR_LOGNAME)

	return m
}
