package generic

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func marshal(_ *lang.Process, iface interface{}) (b []byte, err error) {
	switch v := iface.(type) {
	case string:
		b = []byte(v)
		return

	case []string:
		for i := range v {
			b = append(b, []byte(v[i]+utils.NewLineString)...)
		}
		return

	case [][]string:
		return tabWriter(iface.([][]string))
		/*for i := range v {
			b = append(b, []byte(strings.Join(v[i], "\t")+utils.NewLineString)...)
		}
		return*/

	case []interface{}:
		for i := range v {
			b = append(b, iface2str(&v[i])...)
		}
		return

	case map[string]string:
		return mapToArray(v)
	case map[string]float64:
		return mapToArray(v)
	case map[string]int:
		return mapToArray(v)
	case map[string]bool:
		return mapToArray(v)
	case map[string]interface{}:
		return mapToArray(v)

	case map[interface{}]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s\t%s%s", fmt.Sprint(s), fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]string:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s\t%s%s", fmt.Sprint(s), v[s], utils.NewLineString))...)
		}
		return

	/*case interface{}:
	return []byte(fmt.Sprintln(iface)), nil*/

	default:
		err = fmt.Errorf("I don't know how to marshal %T into a `%s`. Data possibly too complex?", v, types.Generic)
		return
	}
}

func iface2str(iface *interface{}) (b []byte) {
	switch v := (*iface).(type) {
	case []interface{}:
		if len(v) == 0 {
			return
		}

		for i := 0; i < len(v)-2; i++ {
			b = append(b, []byte(fmt.Sprintf("%v\t", v[i]))...)
		}
		return append(b, []byte(fmt.Sprintf("%v%s", v[len(v)-1], utils.NewLineString))...)

	case string:
		return []byte(v + utils.NewLineString)

	case interface{}:
		return []byte(fmt.Sprintf("%v%s", v, utils.NewLineString))

		//default:
		//	return []byte(fmt.Sprintf("%v%s", v, utils.NewLineString))
	default:
		panic(fmt.Sprintf("cannot marshal %T", v))
	}
}

func mapToArray[K comparable, V string | float64 | int | bool | any](m map[K]V) ([]byte, error) {
	var a [][]string
	for k, v := range m {
		a = append(a, []string{fmt.Sprint(k), fmt.Sprint(v)})
	}
	return tabWriter(a)
}

func tabWriter(v [][]string) ([]byte, error) {
	var (
		b   []byte
		err error
	)

	buf := bytes.NewBuffer(b)
	w := tabwriter.NewWriter(buf, twMinWidth, twTabWidth, twPadding, twPadChar, twFlags)

	for i := range v {
		_, err = fmt.Fprintln(w, strings.Join(v[i], "\t"))
		if err != nil {
			return nil, err
		}
	}

	err = w.Flush()
	return buf.Bytes(), err
}

func unmarshal(p *lang.Process) (interface{}, error) {
	table := make([][]string, 1)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		table = append(table, rxWhitespace.Split(scanner.Text(), -1))
	}

	if len(table) > 1 {
		table = table[1:]
	}

	err := scanner.Err()
	if err != nil {
		return table, fmt.Errorf("error while unmarshalling a %s array: %s", types.Generic, err.Error())
	}
	return table, err
}
