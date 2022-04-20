package lang

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
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

// Reserved variable names. Set as constants so any typos of these names within
// the code will be raised as compiler errors
const (
	SELF       = "SELF"
	ARGS       = "ARGS"
	PARAMS     = "PARAMS"
	MUREX_EXE  = "MUREX_EXE"
	MUREX_ARGS = "MUREX_ARGS"
	HOSTNAME   = "HOSTNAME"
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
	DataType string
	Value    interface{}
	String   string
	Modify   time.Time
	FileRef  *ref.File // only needed for globals
}

// GetValue return the value of a variable. If a variable does not exist then
// GetValue will return nil. Please check if p.Config.Get("proc", "strict-vars", "bool")
// matters for your usage of GetValue because this API doesn't care. If in doubt
// use GetString instead.
func (v *Variables) GetValue(name string) interface{} {
	switch name {
	case SELF:
		return getVarSelf(v.process)

	case ARGS:
		return getVarArgs(v.process)

	case PARAMS:
		return v.process.Scope.Parameters.StringArray()

	case MUREX_EXE:
		return getVarMurexExe()

	case HOSTNAME:
		return getHostname()
	}

	if i, err := strconv.Atoi(name); err == nil && i > 0 {
		s, err := v.process.Scope.Parameters.String(i - 1)
		if err != nil {
			return nil
		}
		return s
	}

	if v.global {
		return v.getValue(name)
	}

	value := v.getValue(name)
	if value != nil {
		return value
	}

	value = GlobalVariables.getValue(name)
	if value != nil {
		return value
	}

	// variable not found so lets fallback to the environmental variables
	value = os.Getenv(name)
	if value != "" {
		return value
	}

	return nil
}

func (v *Variables) getValue(name string) (value interface{}) {
	v.mutex.Lock()
	variable := v.vars[name]
	if variable == nil {
		v.mutex.Unlock()
		return nil
	}

	value = variable.Value
	v.mutex.Unlock()
	return value
}

// GetString returns a string representation of the data stored in the requested variable
func (v *Variables) GetString(name string) (string, error) {
	switch name {
	case SELF:
		b, _ := json.Marshal(getVarSelf(v.process), v.process.Stdout.IsTTY())
		return string(b), nil

	case ARGS:
		b, _ := json.Marshal(getVarArgs(v.process), v.process.Stdout.IsTTY())
		return string(b), nil

	case PARAMS:
		b, _ := json.Marshal(v.process.Scope.Parameters.StringArray(), v.process.Stdout.IsTTY())
		return string(b), nil

	case MUREX_EXE:
		return getVarMurexExe().(string), nil

	case HOSTNAME:
		return getHostname(), nil
	}

	if i, err := strconv.Atoi(name); err == nil && i > 0 {
		s, err := v.process.Scope.Parameters.String(i - 1)
		if err != nil {
			return "", errVarNoParam(i, err)
		}
		return s, nil
	}

	if v.global {
		val, _ := v.getString(name)
		return val, nil
	}

	s, exists := v.getString(name)
	if exists {
		return s, nil
	}

	s, exists = GlobalVariables.getString(name)
	if exists {
		return s, nil
	}

	// variable not found so lets fallback to the environmental variables
	s, exists = os.LookupEnv(name)

	if v, err := v.process.Config.Get("proc", "strict-vars", "bool"); err == nil && v.(bool) && !exists {
		return "", errVarNotExist(name)
	}

	return s, nil
}

func (v *Variables) getString(name string) (string, bool) {
	v.mutex.Lock()
	variable := v.vars[name]
	if variable == nil {
		v.mutex.Unlock()
		return "", false
	}

	s := variable.String
	v.mutex.Unlock()
	return s, true
}

// GetDataType returns the data type of the variable stored in the referenced VarTable
func (v *Variables) GetDataType(name string) string {
	switch name {
	case SELF:
		return types.Json

	case ARGS:
		return types.Json

	case PARAMS:
		return types.Json

	case MUREX_EXE:
		return types.String
	}

	if i, err := strconv.Atoi(name); err == nil && i > 0 {
		if i >= v.process.Scope.Parameters.Len() {
			return ""
		}
		return types.String
	}

	if v.global {
		dt, _ := v.getDataType(name)
		return dt
	}

	s, exists := v.getDataType(name)
	if exists {
		return s
	}

	s, exists = GlobalVariables.getDataType(name)
	if exists {
		return s
	}

	// variable not found so lets fallback to the environmental variables
	value := os.Getenv(name)
	if value != "" {
		return types.String
	}

	return ""
}

func (v *Variables) getDataType(name string) (string, bool) {
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

// Set writes a variable
func (v *Variables) Set(p *Process, name string, value interface{}, dataType string) error {
	switch name {
	case SELF, ARGS, PARAMS, MUREX_EXE, MUREX_ARGS, HOSTNAME, "_":
		return errVariableReserved(name)
	}
	for _, r := range name {
		if r < '0' || r > '9' {
			goto notReserved
		}
	}
	return errVariableReserved(name)

notReserved:

	s, err := types.ConvertGoType(value, types.String)
	if err != nil {
		return fmt.Errorf("cannot store variable: %s", err.Error())
	}

	fileRef := v.process.FileRef
	if v.global {
		fileRef = p.FileRef
	}

	v.mutex.Lock()

	v.vars[name] = &variable{
		Value:    value,
		String:   s.(string),
		DataType: dataType,
		Modify:   time.Now(),
		FileRef:  fileRef,
	}

	v.mutex.Unlock()

	return nil
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

	m[SELF] = p.Variables.GetValue(SELF)
	m[ARGS] = p.Variables.GetValue(ARGS)
	m[PARAMS] = p.Variables.GetValue(PARAMS)
	m[MUREX_EXE] = p.Variables.GetValue(MUREX_EXE)
	m[MUREX_ARGS] = p.Variables.GetValue(MUREX_ARGS)
	m[HOSTNAME] = p.Variables.GetValue(HOSTNAME)
	return m
}
