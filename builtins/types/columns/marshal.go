package columns

import (
	"bytes"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func marshal(_ *lang.Process, iface any) (b []byte, err error) {
	switch v := iface.(type) {
	case []string:
		marshalSliceStr(v, b)
		b = bytes.TrimRight(b, "\t")
		return

	/*case [][]string:
	for i := range v {
		b = append(b, []byte(strings.Join(v[i], "\t")+utils.NewLineString)...)
	}
	return*/

	case []any:
		for i := range v {
			b = append(b, iface2str(&v[i])...)
		}
		b = bytes.TrimRight(b, "\t")
		return

	/*case map[string]string:
		for s := range v {
			b = append(b, []byte(s+": "+v[s]+utils.NewLineString)...)
		}
		return
	case map[string]any:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", s, fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return
	case map[any]any:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return
	case map[any]string:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), v[s], utils.NewLineString))...)
		}
		return
	case any:
		return []byte(fmt.Sprintln(iface)), nil*/

	default:
		err = fmt.Errorf("i don't know how to marshal that data into a `%s`. Data possibly too complex?", types.Columns)
		return
	}
}

func marshalSliceStr(a []string, b []byte) {
	for i := range a {
		b = append(b, []byte(a[i]+"\t")...)
	}
}

func iface2str(v *any) (b []byte) {
	return []byte(fmt.Sprintf("%v\t", *v))
}

/*func unmarshal(p *lang.Process) (any, error) {
	table := make([][]string, 1)
	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		table = append(table, rxWhitespace.Split(scanner.Text(), -1))
	}
	if len(table) > 1 {
		table = table[1:]
	}
	err := scanner.Err()
	return table, err
}*/
