package string

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func index(p *lang.Process, params []string) error {
	lines := make(map[int]bool)
	for i := range params {
		num, err := strconv.Atoi(params[i])
		if err != nil {
			return fmt.Errorf("Parameter, `%s`, isn't an integer. %s", params[i], err)
		}
		lines[num] = true
	}

	var (
		i   int
		err error
	)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		if lines[i] != p.IsNot {
			_, err = p.Stdout.Writeln(scanner.Bytes())
			if err != nil {
				break
			}
		}
		i++
	}

	if err != nil {
		return err
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error while indexing a %s map: %s", types.String, err.Error())
	}

	return nil
}
