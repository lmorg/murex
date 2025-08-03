package string

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func marshal(_ *lang.Process, iface any) (b []byte, err error) {
	switch v := iface.(type) {
	case string:
		b = []byte(v)
		return

	case []string:
		for i := range v {
			b = append(b, []byte(v[i]+utils.NewLineString)...)
		}
		return

	case []any:
		for i := range v {
			b = append(b, []byte(fmt.Sprintln(v[i]))...)
		}
		return

	case map[string]string:
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
		return []byte(fmt.Sprintln(iface)), nil

	default:
		err = errors.New("I don't know how to marshal that data into a `str`. Data possibly too complex?")
		return
	}
}

func unmarshal(p *lang.Process) (any, error) {
	s := make([]string, 0)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		s = append(s, strings.TrimSpace(scanner.Text()))
	}

	err := scanner.Err()
	if err != nil {
		return s, fmt.Errorf("error while unmarshalling %s document: %s", types.String, err.Error())
	}

	return s, nil
}
