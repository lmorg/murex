# _murex_ Shell Docs

## API Reference: `lang.IndexTemplateTable()` (template API)

> Returns element(s) from a table

## Description

This is a template API you can use for your custom data types.

It should only be called from `ReadIndex()` and `ReadNotIndex()` functions.

This function ensures consistency with the index, `[`, builtin when used with
different _murex_ data types. Thus making indexing a data type agnostic
capability.



## Examples

Example calling `lang.IndexTemplateTable()` function:

```go
package generic

import (
	"bytes"
	"strings"

	"github.com/lmorg/murex/lang"
)

func index(p *lang.Process, params []string) error {
	cRecords := make(chan []string, 1)

	go func() {
		err := p.Stdin.ReadLine(func(b []byte) {
			cRecords <- rxWhitespace.Split(string(bytes.TrimSpace(b)), -1)
		})
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
		}
		close(cRecords)
	}()

	marshaller := func(s []string) (b []byte) {
		b = []byte(strings.Join(s, "\t"))
		return
	}

	return lang.IndexTemplateTable(p, params, cRecords, marshaller)
}
```

## Detail

### API Source:

```go
package lang

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lmorg/murex/utils"
)

const (
	byRowNumber = iota + 1
	byColumnNumber
	byColumnName
)

var (
	rxColumnPrefix = regexp.MustCompile(`^:[0-9]+$`)
	rxRowSuffix    = regexp.MustCompile(`^[0-9]+:$`)
	errMixAndMatch = errors.New("you cannot mix and match matching modes")
)

// IndexTemplateTable is a handy standard indexer you can use in your custom data types for tabulated / streamed data.
// The point of this is to minimize code rewriting and standardising the behavior of the indexer.
func IndexTemplateTable(p *Process, params []string, cRecords chan []string, marshaller func([]string) []byte) error {
	if p.IsNot {
		return ittNot(p, params, cRecords, marshaller)
	}
	return ittIndex(p, params, cRecords, marshaller)
}

func ittIndex(p *Process, params []string, cRecords chan []string, marshaller func([]string) []byte) (err error) {
	var (
		mode      int
		matchStr  []string
		matchInt  []int
		unmatched []string
	)

	defer func() {
		if len(unmatched) != 0 {
			p.ExitNum = 1
			err = fmt.Errorf("some records did not contain all of the requested fields:%s%s",
				utils.NewLineString,
				strings.Join(unmatched, utils.NewLineString))
		}
	}()

	for i := range params {
		switch {
		case rxRowSuffix.MatchString(params[i]):
			if mode != 0 && mode != byRowNumber {
				return errMixAndMatch
			}
			mode = byRowNumber
			num, _ := strconv.Atoi(params[i][:len(params[i])-1])
			matchInt = append(matchInt, num)

		case rxColumnPrefix.MatchString(params[i]):
			if mode != 0 && mode != byColumnNumber {
				return errMixAndMatch
			}
			mode = byColumnNumber
			num, _ := strconv.Atoi(params[i][1:])
			matchInt = append(matchInt, num)

		default:
			if mode != 0 && mode != byColumnName {
				return errMixAndMatch
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
					_, err = p.Stdout.Writeln(marshaller(recs))
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
				_, err = p.Stdout.Writeln(marshaller(lines[j]))
				if err != nil {
					p.Stderr.Writeln([]byte(err.Error()))
				}
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
					if len(recs) == 0 || (len(recs) == 1 && recs[0] == "") {
						continue
					}
					unmatched = append(unmatched, strings.Join(recs, "\t"))
				}
			}
			if len(line) != 0 {
				_, err = p.Stdout.Writeln(marshaller(line))
				if err != nil {
					p.Stderr.Writeln([]byte(err.Error()))
				}
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
					_, err = p.Stdout.Writeln(marshaller(line))
					if err != nil {
						p.Stderr.Writeln([]byte(err.Error()))
					}
				}

			} else {
				for i := range matchStr {
					col := headings[matchStr[i]]
					if col != 0 && col < len(recs)+1 {
						line = append(line, recs[col-1])
					} else {
						if len(recs) == 0 || (len(recs) == 1 && recs[0] == "") {
							continue
						}
						unmatched = append(unmatched, strings.Join(recs, "\t"))
					}
				}
				if len(line) != 0 {
					_, err = p.Stdout.Writeln(marshaller(line))
					if err != nil {
						p.Stderr.Writeln([]byte(err.Error()))
					}
				}
			}
			lineNum++
		}

	default:
		return errors.New("you haven't selected any rows / columns")
	}
}

func ittNot(p *Process, params []string, cRecords chan []string, marshaller func([]string) []byte) error {
	var (
		mode     int
		matchStr = make(map[string]bool)
		matchInt = make(map[int]bool)
	)

	for i := range params {
		switch {
		case rxRowSuffix.MatchString(params[i]):
			if mode != 0 && mode != byRowNumber {
				return errMixAndMatch
			}
			mode = byRowNumber
			num, _ := strconv.Atoi(params[i][:len(params[i])-1])
			matchInt[num] = true

		case rxColumnPrefix.MatchString(params[i]):
			if mode != 0 && mode != byColumnNumber {
				return errMixAndMatch
			}
			mode = byColumnNumber
			num, _ := strconv.Atoi(params[i][1:])
			matchInt[num] = true

		default:
			if mode != 0 && mode != byColumnName {
				return errMixAndMatch
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
		return errors.New("you haven't selected any rows / columns")
	}
}
```

## Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p` 
2. `[]string`: slice of parameters used in `[` / `![` 
3. `chan []string`: a channel for rows (each element in the slice is a column within the row). This allows tables to be stream-able
4. `func(interface{}) ([]byte, error)`: data type marshaller function

## See Also

* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [apis/`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [apis/`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [apis/`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [apis/`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [apis/`lang.IndexTemplateObject()` (template API)](../apis/lang.IndexTemplateObject.md):
  Returns element(s) from a data structure