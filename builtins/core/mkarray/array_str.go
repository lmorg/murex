package mkarray

import (
	"strings"
)

func (a *arrayT) isStringArray() (err error) {
	a.writer, err = a.p.Stdout.WriteArray(a.dataType)
	if err != nil {
		return err
	}

	err = a.parseExpression()
	if err != nil {
		return err
	}

	return a.writeArrayString()
}

func (a *arrayT) writeArrayString() error {
	var (
		marker = string([]byte{0})
		open   bool
	)

	for g := range a.groups {
		var (
			template string
			variable = make(map[int][]string)
			l        = -1
		)

		for n := range a.groups[g] {
			if a.p.HasCancelled() {
				goto cancelled
			}

			switch a.groups[g][n].Type {
			case astTypeString:
				if open {
					variable[l] = append(variable[l], string(a.groups[g][n].Data))
					continue
				}
				template += string(a.groups[g][n].Data)

			case astTypeRange:
				a, err := rangeToArrayString(a.groups[g][n].Data)
				if err != nil {
					return err
				}
				variable[l] = append(variable[l], a...)
				continue

			case astTypeOpen:
				template += marker
				l++
				variable[l] = make([]string, 0)
				open = true

			case astTypeClose:
				open = false
			}
		}

		counter := make([]int, len(variable))

		for {
		nextIndex:
			if a.p.HasCancelled() {
				goto cancelled
			}

			s := template
			for t := 0; t < len(counter); t++ {
				c := counter[t]
				s = strings.Replace(s, marker, variable[t][c], 1)
			}
			a.writer.WriteString(s)

			i := len(counter) - 1
			if i < 0 {
				goto nextGroup
			}

			counter[i]++
			if counter[i] == len(variable[i]) {
			nextCounter:
				counter[i] = 0
				i--
				if i < 0 {
					goto nextGroup
				}
				counter[i]++
				if counter[i] < len(variable[i]) {
					goto nextIndex
				} else {
					goto nextCounter
				}
			} else {
				goto nextIndex
			}

		}
	nextGroup:
	}

cancelled:
	return a.writer.Close()
}
