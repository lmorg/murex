package proc

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"

	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Variables is an object that methods out lookups against the varTable.
// This will need to be created on each proc.Process.
type Variables struct {
	varTable *varTable
	process  *Process
}

// NewVariables creates a new Variables object
func NewVariables(p *Process) *Variables {
	vars := new(Variables)
	vars.process = p
	vars.varTable = masterVarTable
	return vars
}

// This is the core variable table that will be used for all vars
type varTable struct {
	vars  []*variable
	mutex sync.Mutex
}

func newVarTable() *varTable {
	vt := new(varTable)
	//go garbageCollection(vt)
	return vt
}

/*func garbageCollection(vt *varTable) {
	for {
		time.Sleep(3 * time.Second)
		vt.mutex.Lock()

		for i := 0; i < len(vt.vars); i++ {
			vt.vars[i].mutex.Lock()
			if vt.vars[i].disabled {
				switch i {
				case 0:
					vt.vars = vt.vars[1:]
				case len(vt.vars) - 1:
					vt.vars = vt.vars[:len(vt.vars)-1]
				default:
					vt.vars = append(vt.vars[:i], vt.vars[i+1:]...)
				}
				i--
				continue
			}
			vt.vars[i].mutex.Unlock()
		}

		vt.mutex.Unlock()
	}
}*/

// GetVariable is a single API that handles the logic parsing the varTable
// `readOnly` defines whether to return new variable struct with the same settings
// as the original (true) or the original but mutex locked (false). By default
// you should always return a readOnly copy (true) as that is better for concurrency.
// Also environmental variables only get checked when readOnly == true
// `nil` gets returned if no variable matches the name or owner.
func (vt *varTable) GetVariable(p *Process, name string, copy bool) *variable {
	vt.mutex.Lock()
	if copy {
		defer vt.mutex.Unlock()
	}

	//for i := range vt.vars {
	for i := len(vt.vars) - 1; i > -1; i-- {

		vt.vars[i].mutex.Lock()
		if !vt.vars[i].disabled && vt.vars[i].name == name {

			// variable exists. Check permissions (ie is it in scope?)
			//for _, proc := range p.FidTree {
			//if proc == vt.vars[i].owner {

			// return variable
			if copy {
				vcopy := &vt.vars[i]
				vt.vars[i].mutex.Unlock()
				return *vcopy
			}
			return vt.vars[i]

			//}
			//}
		}

		vt.vars[i].mutex.Unlock()
	}

	s, exists := os.LookupEnv(name)
	if !exists {
		return nil
	}

	return &variable{
		name:     name,
		Value:    s,
		DataType: types.String,
	}
}

// This is a struct for each variable
type variable struct {
	name     string
	Value    interface{}
	DataType string
	owner    int
	disabled bool
	mutex    sync.Mutex
}

// GetValue return the value of a variable stored in the referenced VarTable
func (vars *Variables) GetValue(name string) interface{} {
	v := vars.varTable.GetVariable(vars.process, name, true)
	if v == nil {
		return v
	}

	return v.Value
}

// GetDataType returns the data type of the variable stored in the referenced VarTable
func (vars *Variables) GetDataType(name string) string {
	v := vars.varTable.GetVariable(vars.process, name, true)
	if v == nil {
		return ""
	}

	return v.DataType
}

// GetString returns a string representation of the data stored in the requested variable
func (vars *Variables) GetString(name string) string {
	v := vars.varTable.GetVariable(vars.process, name, true)
	if v == nil {
		return ""
	}

	s, err := types.ConvertGoType(v.Value, types.String)
	if err != nil {
		if debug.Enable {
			panic(err.Error())
		}
		return fmt.Sprint(v.Value)
	}
	return s.(string)
}

func convDataType(value interface{}, dataType string) (val interface{}, err error) {
	switch dataType {
	case types.Integer:
		val, err = types.ConvertGoType(value, dataType)
		//if err != nil {
		//	return err
		//}

	case types.Float, types.Number:
		val, err = types.ConvertGoType(value, dataType)
		//if err != nil {
		//	return err
		//}

	case types.Boolean:
		val, err = types.ConvertGoType(value, types.Boolean)
		//if err != nil {
		//	return err
		//}

	default:
		val, err = types.ConvertGoType(value, types.String)
		//if err != nil {
		//	return err
		//}
	}

	return
}

// Set checks if a variable already exists, if it does it updates the value, if
// it doesn't it creates a new one.
func (vars *Variables) Set(name string, value interface{}, dataType string) error {
	val, err := convDataType(value, dataType)
	if err != nil {
		return err
	}

	v := vars.varTable.GetVariable(vars.process, name, false)
	if v != nil {
		v.Value = val
		v.DataType = dataType
		v.mutex.Unlock()
		vars.varTable.mutex.Unlock()
		return nil
	}

	vars.varTable.vars = append(vars.varTable.vars, &variable{
		name:     name,
		Value:    val,
		DataType: dataType,
		owner:    vars.process.Parent.Id,
	})

	vars.varTable.mutex.Unlock()

	return nil
}

func (vars *Variables) ForceNewScope(name string, value interface{}, dataType string) error {
	val, err := convDataType(value, dataType)
	if err != nil {
		return err
	}

	vars.varTable.vars = append(vars.varTable.vars, &variable{
		name:     name,
		Value:    val,
		DataType: dataType,
		owner:    vars.process.Parent.Id,
	})

	vars.varTable.mutex.Unlock()

	return nil
}

// Unset removes a variable from the table
func (vars *Variables) Unset(name string) error {
	v := vars.varTable.GetVariable(vars.process, name, false)
	if v == nil {
		return errors.New("No variables match the name.")
	}

	v.disabled = true
	v.mutex.Unlock()
	vars.varTable.mutex.Unlock()
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

// DumpEntireTable is a temporary function which is used for debugging. It still be Killed soon
func (vars *Variables) DumpEntireTable() interface{} {
	m := make([]map[string]interface{}, 0)

	for _, v := range vars.varTable.vars {
		mv := map[string]interface{}{
			"name":     v.name,
			"value":    v.Value,
			"datatype": v.DataType,
			"owner":    v.owner,
			"enabled":  !v.disabled,
		}

		m = append(m, mv)
	}
	return m
}
