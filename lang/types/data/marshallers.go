package data

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/csv"
)

func marshalCsv(*proc.Process, interface{}) ([]byte, error) {
	/*switch v := t.(type) {
	case []string:
		//return csv.ArrayToCsv(v), nil
	case map[string]string:
	case map[string][]string:
	case []map[string]string:
	case []map[string]interface{}:
	case []interface{}:
	default:
	}
	return nil, nil*/
	return nil, errors.New("TODO: marsheller not yet written!")
}

func unmarshalCsv(p *proc.Process) (interface{}, error) {
	csvReader, err := csv.NewParser(p.Stdin, &proc.GlobalConf)
	if err != nil {
		return nil, err
	}

	table := make([]map[string]string, 0)
	csvReader.ReadLine(func(recs []string, heads []string) {
		record := make(map[string]string)
		for i := range recs {
			record[heads[i]] = recs[i]
		}
		table = append(table, record)
	})

	return table, nil
}

func marshalJson(p *proc.Process, v interface{}) ([]byte, error) {
	return utils.JsonMarshal(v, p.Stdout.IsTTY())
}

func unmarshalJson(p *proc.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &v)
	return
}

func marshalString(p *proc.Process, v interface{}) ([]byte, error) {
	return nil, errors.New("TODO: marsheller not yet written!")
}

func unmarshalString(p *proc.Process) (interface{}, error) {
	s := make([]string, 0)
	err := p.Stdin.ReadLine(func(b []byte) {
		s = append(s, string(b))
	})

	return s, err
}
