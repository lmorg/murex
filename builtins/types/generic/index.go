package generic

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"regexp"
	"strconv"
)

var (
	rxWhitespace   *regexp.Regexp = regexp.MustCompile(`\s+`)
	rxColumnPrefix *regexp.Regexp = regexp.MustCompile(`^:[0-9]+$`)
	rxRowSuffix    *regexp.Regexp = regexp.MustCompile(`^[0-9]+:$`)
)

func index(p *proc.Process, params []string) error {
	var (
		mode     int
		matchStr []string
		matchInt []int
	)

	for i := range params {
		switch {
		case rxRowSuffix.MatchString(params[i]):
			if mode != 0 && mode != 1 {
				return errors.New("You cannot mix and match matching modes.")
			}
			mode = 1
			num, _ := strconv.Atoi(params[i][:len(params[i])-1])
			matchInt = append(matchInt, num)

		case rxColumnPrefix.MatchString(params[i]):
			if mode != 0 && mode != 2 {
				return errors.New("You cannot mix and match matching modes.")
			}
			mode = 2
			num, _ := strconv.Atoi(params[i][1:])
			matchInt = append(matchInt, num)

		default:
			if mode != 0 && mode != 3 {
				return errors.New("You cannot mix and match matching modes.")
			}
			matchStr = append(matchStr, params[i])
			mode = 3

		}
	}

	switch mode {
	case 1:
		// TODO: match rows

	case 2:
		// Match column numbers
		p.Stdin.ReadLine(func(b []byte) {
			recs := rxWhitespace.Split(string(b), -1)
			var line string
			for _, i := range matchInt {
				if i < len(recs) {
					line += "\t" + recs[i]
				} else {
					p.ExitNum = 1
				}
			}
			if len(line) > 1 {
				p.Stdout.Writeln([]byte(line[1:]))
			}
		})

	case 3:
		// Match column names
		var (
			lineNum  int
			headings map[string]int = make(map[string]int)
		)

		p.Stdin.ReadLine(func(b []byte) {
			recs := rxWhitespace.Split(string(b), -1)
			if lineNum == 0 {
				for i := range recs {
					headings[recs[i]] = i + 1
				}

			} else {
				var line string
				for i := range matchStr {
					col := headings[matchStr[i]]
					if col != 0 && col < len(recs) {
						line += "\t" + recs[col-1]
					} else {
						p.ExitNum = 1
					}
				}
				if len(line) > 1 {
					p.Stdout.Writeln([]byte(line[1:]))
				}
			}
			lineNum++
		})

	default:
		return errors.New("You haven't selected any rows / columns.")
	}

	return nil
}
