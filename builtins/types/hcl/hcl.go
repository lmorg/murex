package hcl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	hcljson "github.com/hashicorp/hcl/v2/json"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/json"
)

const typeName = "hcl"

func init() {
	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex

	//stdio.RegisterReadArray(typeName, readArray)
	//stdio.RegisterReadArrayWithType(typeName, readArrayWithType)
	//stdio.RegisterReadMap(typeName, readMap)
	stdio.RegisterWriteArray(typeName, newArrayWriter)

	lang.Marshallers[typeName] = marshal
	lang.Unmarshallers[typeName] = unmarshal

	// These are just guessed at as I couldn't find any formally named MIMEs
	lang.SetMime(typeName,
		"application/hcl",
		"application/x-hcl",
		"text/hcl",
		"text/x-hcl",
	)

	lang.SetFileExtensions(typeName, "hcl", "tf", "tfvars")
}

func decode(data []byte, v interface{}) error {
	filename := ""
	//file, diags := hclsyntax.ParseConfig(data, filename, hcl.Pos{Line: 1, Column: 1})
	file, diags := hcljson.Parse(data, filename)
	if diags.HasErrors() {
		return diags
	}

	var ctx *hcl.EvalContext

	diags = gohcl.DecodeBody(file.Body, ctx, v)
	if diags.HasErrors() {
		return diags
	}
	return nil
}

/*func readArray(read stdio.Io, callback func([]byte)) error {
	// Create a marshaller function to pass to ArrayTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayTemplate(marshaller, unmarshaller, read, callback)
}

func readArrayWithType(read stdio.Io, callback func([]byte, string)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayWithTypeTemplate(types.Json, marshaller, unmarshaller, read, callback)
}

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var jObj interface{}
	err = unmarshaller(b, &jObj)
	if err == nil {

		switch v := jObj.(type) {
		case []interface{}:
			for i := range jObj.([]interface{}) {
				j, err := json.Marshal(jObj.([]interface{})[i], false)
				if err != nil {
					return err
				}
				callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
			}

		case map[string]interface{}, map[interface{}]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				j, err := json.Marshal(jObj.(map[string]interface{})[key], false)
				if err != nil {
					return err
				}
				callback(key, string(j), i != len(jObj.(map[string]interface{})))
				i++
			}
			return nil

		default:
			if debug.Enabled {
				panic(v)
			}
		}
		return nil
	}
	return err
}*/

func readIndex(p *lang.Process, params []string) error {
	v := make(map[string]interface{})

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = decode(b, &v)
	if err != nil {
		return err
	}

	var vWrapper interface{} = v

	marshaller := func(iface interface{}) ([]byte, error) {
		return json.Marshal(iface, p.Stdout.IsTTY())
	}

	return lang.IndexTemplateObject(p, params, &vWrapper, marshaller)
}

func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	return json.Marshal(v, p.Stdout.IsTTY())
}

func unmarshal(p *lang.Process) (interface{}, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	v := make(map[string]interface{})

	err = decode(b, &v)
	return v, err
}
