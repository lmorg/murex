package define

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/lmorg/murex/lang/proc"
)

const (
	byRowNumber = iota + 1
	byColumnNumber
	byColumnName
)

var (
	rxColumnPrefix = regexp.MustCompile(`^:[0-9]+$`)
	rxRowSuffix    = regexp.MustCompile(`^[0-9]+:$`)
)

// IndexTemplateTable is a handy standard indexer you can use in your custom data types for tabulated / streamed data.
// The point of this is to minimize code rewriting and standardising the behavior of the indexer.
func IndexTemplateTable(p *proc.Process, params []string, cRecords chan []string, marshaller func([]string) []byte) error {
	if p.IsNot {
		return ittNot(p, params, cRecords, marshaller)
	}
	return ittIndex(p, params, cRecords, marshaller)
}

func ittIndex(p *proc.Process, params []string, cRecords chan []string, marshaller func([]string) []byte) error {
	var (
		mode     int
		matchStr []string
		matchInt []int
	)

	for i := range params {
		switch {
		case rxRowSuffix.MatchString(params[i]):
			if mode != 0 && mode != byRowNumber {
				return errors.New("You cannot mix and match matching modes")
			}
			mode = byRowNumber
			num, _ := strconv.Atoi(params[i][:len(params[i])-1])
			matchInt = append(matchInt, num)

		case rxColumnPrefix.MatchString(params[i]):
			if mode != 0 && mode != byColumnNumber {
				return errors.New("You cannot mix and match matching modes")
			}
			mode = byColumnNumber
			num, _ := strconv.Atoi(params[i][1:])
			matchInt = append(matchInt, num)

		default:
			if mode != 0 && mode != byColumnName {
				return errors.New("You cannot mix and match matching modes")
			}
			matchStr = append(matchStr, params[i])
			mode = byColumnName

		}
	}

	switch mode {
	case byRowNumber:
		var (
			ordered = true
			last    int
			max     int
		)
		// check order
		for _, i := range matchInt {
			if i < last {
				ordered = false
			}
			if i > max {
				max = i
			}
			last = i
		}

		if ordered {
			// ordered matching - for this we can just read in the records we want sequentially. Low memory overhead
			var i int
			for {
				recs, ok := <-cRecords
				if !ok {
					return nil
				}
				if i == matchInt[0] {
					_, err := p.Stdout.Writeln(marshaller(recs))
					if err != nil {
						p.Stderr.Writeln([]byte(err.Error()))
					}
					if len(matchInt) == 1 {
						matchInt[0] = -1
						return nil
					}
					matchInt = matchInt[1:]
				}
				i++
			}

		} else {
			// unordered matching - for this we load the entire data set into memory - up until the maximum value
			var (
				i     int
				lines = make([][]string, max+1)
			)
			for {
				recs, ok := <-cRecords
				if !ok {
					break
				}
				if i <= max {
					lines[i] = recs
				}
				i++
			}

			for _, j := range matchInt {
				p.Stdout.Writeln(marshaller(lines[j]))
			}

			return nil
		}

	case byColumnNumber:
		for {
			recs, ok := <-cRecords
			if !ok {
				return nil
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
		}

	case byColumnName:
		var (
			lineNum  int
			headings = make(map[string]int)
		)

		for {
			var line []string
			recs, ok := <-cRecords
			if !ok {
				return nil
			}

			if lineNum == 0 {
				for i := range recs {
					headings[recs[i]] = i + 1
				}
				for i := range matchStr {
					if headings[matchStr[i]] != 0 {
						line = append(line, matchStr[i])
					}
				}
				if len(line) != 0 {
					p.Stdout.Writeln(marshaller(line))
				}

			} else {
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
		}

	default:
		return errors.New("You haven't selected any rows / columns")
	}
}

func ittNot(p *proc.Process, params []string, cRecords chan []string, marshaller func([]string) []byte) error {
	var (
		mode     int
		matchStr = make(map[string]bool)
		matchInt = make(map[int]bool)
	)

	for i := range params {
		switch {
		case rxRowSuffix.MatchString(params[i]):
			if mode != 0 && mode != byRowNumber {
				return errors.New("You cannot mix and match matching modes")
			}
			mode = byRowNumber
			num, _ := strconv.Atoi(params[i][:len(params[i])-1])
			matchInt[num] = true

		case rxColumnPrefix.MatchString(params[i]):
			if mode != 0 && mode != byColumnNumber {
				return errors.New("You cannot mix and match matching modes")
			}
			mode = byColumnNumber
			num, _ := strconv.Atoi(params[i][1:])
			matchInt[num] = true

		default:
			if mode != 0 && mode != byColumnName {
				return errors.New("You cannot mix and match matching modes")
			}
			matchStr[params[i]] = true
			mode = byColumnName

		}
	}

	switch mode {
	case byRowNumber:
		i := -1
		for {
			recs, ok := <-cRecords
			if !ok {
				return nil
			}

			if !matchInt[i] {
				_, err := p.Stdout.Writeln(marshaller(recs))
				if err != nil {
					p.Stderr.Writeln([]byte(err.Error()))
				}
			}
			i++
		}

	case byColumnNumber:
		for {
			recs, ok := <-cRecords
			if !ok {
				return nil
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
		}

	case byColumnName:
		var (
			lineNum  int
			headings = make(map[int]string)
		)

		for {
			var line []string
			recs, ok := <-cRecords
			if !ok {
				return nil
			}

			if lineNum == 0 {
				for i := range recs {
					headings[i] = recs[i]
					if !matchStr[headings[i]] {
						line = append(line, recs[i])
					}
				}
				if len(line) != 0 {
					p.Stdout.Writeln(marshaller(line))
				}

			} else {
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
		}

	default:
		return errors.New("You haven't selected any rows / columns")
	}
}
