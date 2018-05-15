package proc

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
)

// Variables is an object that methods out lookups against the `varTable`.
// This will need to be created on each `proc.Process`.
//
// While it might seem odd wrapping `varTable` struct up inside another struct,
// the idea behind this is `Variables` would be per process and `varTable` would
// be global. `Variables` then references the `varTable`. This allows us to do
// some clever things with variables such as have scopes that don't have any
// visibility of even the shell's global vars.
type Variables struct {
	varTable *varTable
	process  *Process
	time     time.Time
}

// NewVariables creates a new Variables object
func NewVariables() *Variables {
	vars := new(Variables)
	vars.varTable = new(varTable)
	vars.process = ShellProcess
	return vars
}

// ReferenceVariables creates a new Variables object linked to an existing varTable
func ReferenceVariables(ref *Variables) *Variables {
	//vars := new(Variables)
	//vars.varTable = ref.varTable
	//return vars
	return &Variables{varTable: ref.varTable}
}

// This is the core variable table that will be used for all vars
type varTable struct {
	vars  []*variable
	mutex sync.Mutex
}

func newVarTable() *varTable {
	vt := new(varTable)
	go garbageCollection(vt)
	return vt
}

func garbageCollection(vt *varTable) {
	for {
		time.Sleep(10 * time.Second)

		vt.mutex.Lock()
		for i := 0; i < len(vt.vars); i++ {
			vt.vars[i].mutex.Lock()
			disabled := vt.vars[i].disabled
			vt.vars[i].mutex.Unlock()

			if disabled {
				switch i {
				case 0:
					vt.vars = vt.vars[1:]
				case len(vt.vars) - 1:
					vt.vars = vt.vars[:len(vt.vars)-1]
				default:
					vt.vars = append(vt.vars[:i], vt.vars[i+1:]...)
				}
				i--
			}
		}
		vt.mutex.Unlock()
	}
}

// CloseScopedVariables marks variables ready for garbage collection
func CloseScopedVariables(p *Process) {
	p.Variables.varTable.mutex.Lock()
	for _, v := range p.Variables.varTable.vars {
		if v.owner == p.Id {
			v.mutex.Lock()
			v.disabled = true
			v.mutex.Unlock()
		}
	}
	p.Variables.varTable.mutex.Unlock()
}

func (vt *varTable) getVariable(p *Process, name string) *variable {
	var candidate *variable

	vt.mutex.Lock()

	for _, v := range vt.vars {
		v.mutex.Lock()
		disabled := v.disabled
		v.mutex.Unlock()
		if disabled || v.name != name /*|| v.creationTime.After(p.StartTime)*/ {
			continue
		}

		for i := range p.FidTree {
			if p.FidTree[i] == v.owner && (candidate == nil || v.owner > candidate.owner) {
				candidate = v
				break
			}
		}
	}

	vt.mutex.Unlock()

	return candidate
}

// This is a struct for each variable
type variable struct {
	name         string
	Value        interface{}
	DataType     string
	owner        int
	disabled     bool
	creationTime time.Time
	mutex        sync.Mutex
}

// GetValue return the value of a variable stored in the referenced VarTable
func (vars *Variables) GetValue(name string) interface{} {
	v := vars.varTable.getVariable(vars.process, name)
	if v != nil {
		v.mutex.Lock()
		value := v.Value
		v.mutex.Unlock()

		return value
	}

	// variable not found so lets fallback to the environmental variables
	value := os.Getenv(name)
	if value != "" {
		return value
	}

	return nil
}

// GetDataType returns the data type of the variable stored in the referenced VarTable
func (vars *Variables) GetDataType(name string) string {
	v := vars.varTable.getVariable(vars.process, name)
	if v != nil {
		v.mutex.Lock()
		dt := v.DataType
		v.mutex.Unlock()

		return dt
	}

	// variable not found so lets fallback to the environmental variables
	value := os.Getenv(name)
	if value != "" {
		return types.String
	}

	return ""
}

// GetString returns a string representation of the data stored in the requested variable
func (vars *Variables) GetString(name string) string {
	v := vars.varTable.getVariable(vars.process, name)
	if v != nil {
		v.mutex.Lock()
		value := v.Value
		v.mutex.Unlock()

		s, err := types.ConvertGoType(value, types.String)
		if err != nil {
			if debug.Enable {
				panic(err.Error())
			}
			return fmt.Sprint(value) // silent fallback for stability
		}

		return s.(string)
	}

	// variable not found so lets fallback to the environmental variables
	value := os.Getenv(name)
	return value
}

// this is rather pointless!!!
func convDataType(value interface{}, dataType string) (val interface{}, err error) {
	switch dataType {
	case types.Number, types.Float:
		val, err = types.ConvertGoType(value, dataType)

	case types.Integer:
		val, err = types.ConvertGoType(value, dataType)

	case types.Boolean:
		val, err = types.ConvertGoType(value, dataType)

	default:
		// this is literally the only time we are overriding the default for
		// ConvertGoType!!
		val, err = types.ConvertGoType(value, types.String)
	}

	return
}

// Set checks if a variable already exists, if it does it updates the value, if
// it doesn't it creates a new one.
func (vars *Variables) Set(name string, value interface{}, dataType string) error {
	//debug.Json("vars set", vars.process)
	val, err := convDataType(value, dataType)
	if err != nil {
		return err
	}

	v := vars.varTable.getVariable(vars.process, name)
	if v != nil {
		v.mutex.Lock()
		v.Value = val
		v.DataType = dataType
		v.mutex.Unlock()

		return nil
	}

	vars.varTable.mutex.Lock()
	vars.varTable.vars = append(vars.varTable.vars, &variable{
		name:         name,
		Value:        val,
		DataType:     dataType,
		owner:        vars.process.Id,
		creationTime: time.Now(),
	})
	vars.varTable.mutex.Unlock()

	return nil
}

// Unset removes a variable from the table
func (vars *Variables) Unset(name string) error {
	v := vars.varTable.getVariable(vars.process, name)
	if v == nil {
		return errors.New("No variables match the name.")
	}

	v.mutex.Lock()
	v.disabled = true
	v.mutex.Unlock()
	return nil
}

// Dump returns a map of the structure of all variables in scope
func (vars *Variables) Dump() map[string]*variable {
	m := make(map[string]*variable)

	vars.varTable.mutex.Lock()

	for i := range vars.varTable.vars {

		vars.varTable.vars[i].mutex.Lock()
		if !vars.varTable.vars[i].disabled {

			for _, proc := range vars.process.FidTree {

				if vars.varTable.vars[i].owner == proc {
					vcopy := &vars.varTable.vars[i]
					m[vars.varTable.vars[i].name] = *vcopy
					continue
				}

			}

		}
		vars.varTable.vars[i].mutex.Unlock()

	}
	vars.varTable.mutex.Unlock()

	return m
}

// DumpMap returns a map of the variables and values for all variables in scope.
// This isn't recommended for general consumption but is needed for the `eval`
// function.
func (vars *Variables) DumpMap() map[string]interface{} {
	m := make(map[string]interface{})
	dump := vars.Dump()
	for v := range dump {
		m[v] = dump[v].Value
	}

	envVars := os.Environ()
	for i := range envVars {
		split := strings.Split(envVars[i], "=")
		m[split[0]] = strings.Join(split[1:], "=")
	}

	return m
}

// DumpEntireTable is a temporary function which is used for debugging.
// It should be killed soon
func (vars *Variables) DumpEntireTable() interface{} {
	m := make([]map[string]interface{}, 0)

	for _, v := range vars.varTable.vars {
		mv := map[string]interface{}{
			"name":     v.name,
			"value":    v.Value,
			"datatype": v.DataType,
			"owner":    v.owner,
			"enabled":  !v.disabled,
			"created":  v.creationTime,
		}

		m = append(m, mv)
	}
	return m
}
