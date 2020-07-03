package lang

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

var (
	errVariableReserved     = errors.New("Cannot set a reserved variable")
	errVarNotExist          = errors.New("Variable does not exist")
	errInvalidVarProperty   = "Invalid varProperty set!"
	errNoDirectGlobalMethod = "This method cannot be invoked directly for global variables"
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

// GetValue return the value of a variable
func (v *Variables) GetValue(name string) interface{} {
	switch {
	case v.global:
		return v.getValue(name)

	case name == "SELF":
		return getVarSelf(v.process)

	case name == "ARGS":
		return getVarParams(v.process)
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
func (v *Variables) GetString(name string) string {
	switch {
	case v.global:
		val, _ := v.getString(name)
		return val
		//panic(errNoDirectGlobalMethod)

	case name == "SELF":
		b, _ := json.Marshal(getVarSelf(v.process), v.process.Stdout.IsTTY())
		return string(b)

	case name == "ARGS":
		b, _ := json.Marshal(getVarParams(v.process), v.process.Stdout.IsTTY())
		return string(b)
	}

	s, exists := v.getString(name)
	if exists {
		return s
	}

	s, exists = GlobalVariables.getString(name)
	if exists {
		return s
	}

	// variable not found so lets fallback to the environmental variables
	return os.Getenv(name)
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
	switch {
	case v.global:
		dt, _ := v.getDataType(name)
		return dt
		//panic(errNoDirectGlobalMethod)

	case name == "SELF":
		return types.Json

	case name == "ARGS":
		return types.Json
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
	case "SELF", "ARGS", "_":
		return errVariableReserved
	}

	s, err := types.ConvertGoType(value, types.String)
	if err != nil {
		return fmt.Errorf("Cannot store variable: %s", err.Error())
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
		return errVarNotExist
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

	envVars := os.Environ()
	for i := range envVars {
		split := strings.Split(envVars[i], "=")
		m[split[0]] = strings.Join(split[1:], "=")
	}

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

	return m
}
