package generic

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
	"strings"
)

func marshal(_ *proc.Process, iface interface{}) (b []byte, err error) {
	switch v := iface.(type) {
	case []string:
		for i := range v {
			b = append(b, []byte(v[i]+utils.NewLineString)...)
		}
		return

	case []interface{}:
		for i := range v {
			b = append(b, iface2str(&v[i])...)
		}
		return

	case map[string]string:
		for s := range v {
			b = append(b, []byte(s+": "+v[s]+utils.NewLineString)...)
		}
		return

	case map[string]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", s, fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]string:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), v[s], utils.NewLineString))...)
		}
		return

	case interface{}:
		return []byte(fmt.Sprintln(iface)), nil

	default:
		err = errors.New("I don't know how to marshal that data into a `*`. Data possibly too complex?")
		return
	}
}

func iface2str(v *interface{}) (b []byte) {
	switch t := (*v).(type) {
	case string:
		return []byte((*v).(string) + utils.NewLineString)
	case int, uint, float64:
		s := fmt.Sprintln(*v)
		return []byte(s)
	default:
		s := fmt.Sprintf("%s: %s%s", t, *v, utils.NewLineString)
		return []byte(s)
	}
}

func unmarshal(p *proc.Process) (interface{}, error) {
	s := make([]string, 0)
	/*err := p.Stdin.ReadLine(func(b []byte) {
		s = append(s, string(b))
	})

	return s, err*/
	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		s = append(s, strings.TrimSpace(scanner.Text()))
	}

	err := scanner.Err()
	return s, err
}
