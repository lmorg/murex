package define

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/ansi"
	"regexp"
	"strconv"
)

var (
	rxColumnPrefix *regexp.Regexp = regexp.MustCompile(`^:[0-9]+$`)
	rxRowSuffix    *regexp.Regexp = regexp.MustCompile(`^[0-9]+:$`)
)

// IndexTemplateTable is a handy standard indexer you can use in your custom data types for tabulated / streamed data.
// The point of this is to minimize code rewriting and standardising the behavior of the indexer.
func IndexTemplateTable(p *proc.Process, params []string, unmarshaller func([]byte) ([]string, error), marshaller func([]string) []byte) error {
	if p.IsNot {
		return ittNot(p, params, unmarshaller, marshaller)
	}
	return ittIndex(p, params, unmarshaller, marshaller)
}

func ittIndex(p *proc.Process, params []string, unmarshaller func([]byte) ([]string, error), marshaller func([]string) []byte) error {
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
		// Match row numbers
		var (
			unordered bool
			last      int
			max       int
		)
		// check order
		for _, i := range matchInt {
			if i < last {
				unordered = true
			}
			if i > max {
				max = i
			}
			last = i
		}

		if !unordered {
			// ordered matching - for this we can just read in the records we want sequentially. Low memory overhead
			var i int
			err := p.Stdin.ReadLine(func(b []byte) {
				if i == matchInt[0] {
					_, err := p.Stdout.Write(b)
					if err != nil {
						p.Stderr.Writeln([]byte(err.Error()))
					}
					if len(matchInt) == 1 {
						matchInt[0] = -1
						return
					}
					matchInt = matchInt[1:]
				}
				i++
			})
			if err != nil {
				return err
			}

		} else {
			// unordered matching - for this we load the entire data set into memory - up until the maximum value
			var (
				i    int
				recs map[int][]byte = make(map[int][]byte)
			)
			err := p.Stdin.ReadLine(func(b []byte) {
				if i <= max {
					recs[i] = b
				}
				i++
			})
			if err != nil {
				return err
			}
			for _, i = range matchInt {
				p.Stdout.Write(recs[i])
			}

		}

	case 2:
		// Match column numbers
		p.Stdin.ReadLine(func(b []byte) {
			recs, err := unmarshaller(b)
			if err != nil {
				ansi.Stderrln(ansi.FgRed, err.Error())
				return
			}

			var line []string
			for _, i := range matchInt {
				if i < len(recs) {
					line = append(line, recs[i])
				} else {
					p.ExitNum = 1
				}
			}
			if len(line) != 0 {
				p.Stdout.Writeln(marshaller(line))
			}
		})

	case 3:
		// Match column names
		var (
			lineNum  int
			headings map[string]int = make(map[string]int)
		)

		p.Stdin.ReadLine(func(b []byte) {
			recs, err := unmarshaller(b)
			if err != nil {
				ansi.Stderrln(ansi.FgRed, err.Error())
				return
			}

			if lineNum == 0 {
				for i := range recs {
					headings[recs[i]] = i + 1
				}

			} else {
				var line []string
				for i := range matchStr {
					col := headings[matchStr[i]]
					if col != 0 && col < len(recs) {
						line = append(line, recs[col-1])
					} else {
						p.ExitNum = 1
					}
				}
				if len(line) != 0 {
					p.Stdout.Writeln(marshaller(line))
				}
			}
			lineNum++
		})

	default:
		return errors.New("You haven't selected any rows / columns.")
	}

	return nil
}

func ittNot(p *proc.Process, params []string, unmarshaller func([]byte) ([]string, error), marshaller func([]string) []byte) error {
	var (
		mode     int
		matchStr map[string]bool = make(map[string]bool)
		matchInt map[int]bool    = make(map[int]bool)
	)

	for i := range params {
		switch {
		case rxRowSuffix.MatchString(params[i]):
			if mode != 0 && mode != 1 {
				return errors.New("You cannot mix and match matching modes.")
			}
			mode = 1
			num, _ := strconv.Atoi(params[i][:len(params[i])-1])
			matchInt[num] = true

		case rxColumnPrefix.MatchString(params[i]):
			if mode != 0 && mode != 2 {
				return errors.New("You cannot mix and match matching modes.")
			}
			mode = 2
			num, _ := strconv.Atoi(params[i][1:])
			matchInt[num] = true

		default:
			if mode != 0 && mode != 3 {
				return errors.New("You cannot mix and match matching modes.")
			}
			matchStr[params[i]] = true
			mode = 3

		}
	}

	switch mode {
	case 1:
		var i int
		err := p.Stdin.ReadLine(func(b []byte) {
			if !matchInt[i] {
				_, err := p.Stdout.Write(b)
				if err != nil {
					p.Stderr.Writeln([]byte(err.Error()))
				}
			}
			i++
		})
		if err != nil {
			return err
		}

	case 2:
		// Match column numbers
		p.Stdin.ReadLine(func(b []byte) {
			recs, err := unmarshaller(b)
			if err != nil {
				ansi.Stderrln(ansi.FgRed, err.Error())
				return
			}

			var line []string
			for i := range recs {
				if !matchInt[i] {
					line = append(line, recs[i])
				}
			}
			if len(line) != 0 {
				p.Stdout.Writeln(marshaller(line))
			}
		})

	case 3:
		// Match column names
		var (
			lineNum  int
			headings map[int]string = make(map[int]string)
		)

		p.Stdin.ReadLine(func(b []byte) {
			recs, err := unmarshaller(b)
			if err != nil {
				ansi.Stderrln(ansi.FgRed, err.Error())
				return
			}

			if lineNum == 0 {
				for i := range recs {
					headings[i] = recs[i]
				}

			} else {
				var line []string
				for i := range recs {
					if !matchStr[headings[i]] {
						line = append(line, recs[i])
					}
				}

				if len(line) != 0 {
					p.Stdout.Writeln(marshaller(line))
				}
			}
			lineNum++
		})

	default:
		return errors.New("You haven't selected any rows / columns.")
	}

	return nil
}
