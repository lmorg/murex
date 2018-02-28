package variables

import (
	"fmt"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
	"sync"
)

type VarTable struct {
	vars  []*variable
	mutex sync.Mutex
}

type variable struct {
	name     string
	value    interface{}
	dataType string
}

// GetValue return the value of a variable stored in the referenced VarTable
func (v *VarTable) GetValue(name string) interface{} {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	for i := range v.vars {
		if v.vars[i].name == name {
			return v.vars[i].value
		}
	}

	return nil
}

// GetDataType returns the dataType of the variable stored in the referenced VarTable
func (v *VarTable) GetDataType(name string) string {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	for i := range v.vars {
		if v.vars[i].name == name {
			return v.vars[i].dataType
		}
	}

	return ""
}

// GetString returns a string representation of the data stored in the requested variable
func (v *VarTable) GetString(name string) string {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	for i := range v.vars {
		if v.vars[i].name == name {
			s, err := types.ConvertGoType(v.vars[i].value, types.String)
			if err != nil {
				if debug.Enable {
					panic(err.Error())
				}
				return fmt.Sprint(v.vars[i].value)
			}
			return s.(string)
		}
	}

	return ""
}

// Set checks if a variable already exists, if it does it updates the value, if
// it doesn't it creates a new one.
func (v *VarTable) Set(name string, value interface{}, dataType string) (err error) {
	var val interface{}

	switch dataType {
	case types.Integer:
		val, err = types.ConvertGoType(value, dataType)
		if err != nil {
			return err
		}

	case types.Float, types.Number:
		val, err = types.ConvertGoType(value, dataType)
		if err != nil {
			return err
		}

	case types.Boolean:
		val, err = types.ConvertGoType(value, types.Boolean)
		if err != nil {
			return err
		}

	default:
		val, err = types.ConvertGoType(value, types.String)
		if err != nil {
			return err
		}
	}

	v.mutex.Lock()
	defer v.mutex.Unlock()

	for i := range v.vars {
		if v.vars[i].name == name {
			v.vars[i].value = val
			v.vars[i].dataType = dataType
			return nil
		}
	}

	v.vars = append(v.vars, &variable{
		name:     name,
		value:    val,
		dataType: dataType,
	})

	return nil
}

// Copy is used to clone the VarTable
func (v *VarTable) Copy() *VarTable {
	vt := new(VarTable)
	vt.vars = make([]*variable, len(v.vars))
	copy(vt.vars, v.vars)
	//vt.vars = append([]*variable{}, v.vars...)
	return vt
}
