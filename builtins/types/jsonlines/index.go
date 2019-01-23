package jsonlines

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
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

	return err
}
