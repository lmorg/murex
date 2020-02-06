package unicode

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
)

func index(p *lang.Process, params []string) error {
	indexes := make([]int, len(params))
	utf8str := make([]rune, len(params))

	for j := range params {
		i, err := strconv.Atoi(params[j])
		if err != nil {
			return fmt.Errorf("Parameter %d not a number '%s' - %s index expects all parameters to be numeric (%s)",
				j, params[j], dataType, err)
		}
		indexes[j] = i
	}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	u := []rune(string(b))

	for j := range indexes {
		i := indexes[j]
		switch {
		case i > -1 && i < len(u):
			utf8str[j] = u[i]

		case i < 0 && i*-1 <= len(u):
			utf8str[j] = u[len(u)+i]

		default:
			return fmt.Errorf("Index out of bounds: %d (max. values %d / -%d)", i, len(u)-1, len(u))
		}
	}

	_, err = p.Stdout.Write([]byte(string(utf8str)))
	return err
}
