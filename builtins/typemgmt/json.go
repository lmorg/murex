package typemgmt

/*import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["->"] = proc.GoFunction{Func: indexJson, TypeIn: types.Die, TypeOut: types.Generic}
}

func indexJson(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	var jInterface interface{}

	if err = json.Unmarshal(p.Stdin.ReadAll(), &jInterface); err != nil {
		return
	}

	for _, field := range p.Parameters.StringArray() {
		switch t := jInterface.(type) {
		case map[string]interface{}:
			jInterface = t[field]

		case string:
			jInterface = t

		default:
			errors.New("Unable to find " + p.Parameters.StringAll() + " in JSON.")
			return
		}
	}

	b, err := json.MarshalIndent(jInterface, "", "\t")
	p.Stdout.Write(b)
	return err
}
*/
