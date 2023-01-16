package mkarray

import (
	"regexp"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

// var rxIsNumberArray = regexp.MustCompile(`^\[([0-9]+)..([0-9]+)\]$`)
var rxIsNumberArray = regexp.MustCompile(`^\[([0-9]+..[0-9]+|[0-9]+|,)+\]$`)

func (a *arrayT) isNumberArray() (bool, error) {
	var err error

	// these data types are all strings anyway. So no point making them numeric
	if a.dataType == types.String || a.dataType == types.Generic {
		return false, nil
	}

	if !rxIsNumberArray.Match(a.expression) {
		return false, nil
	}

	err = a.parseExpression()
	if err != nil {
		return false, err
	}

	if len(a.groups) != 1 {
		return false, nil
	}

	return a.writeArrayNumber()
}

func (a *arrayT) writeArrayNumber() (bool, error) {
	var array []int

	for n := range a.groups[0] {
		if a.p.HasCancelled() {
			goto cancelled
		}

		switch a.groups[0][n].Type {
		case astTypeString:
			if len(a.groups[0][n].Data) == 0 {
				continue
			}
			if len(a.groups[0][n].Data) > 1 && a.groups[0][n].Data[0] == '0' {
				// numbers prefixed with a zero should be a string
				return false, nil
			}
			i, err := strconv.Atoi(string(a.groups[0][n].Data))
			if err != nil {
				return false, err
			}
			array = append(array, i)

		case astTypeRange:
			v, isNum, err := rangeToArrayNumber(a.groups[0][n].Data)
			if !isNum || err != nil {
				return isNum, err
			}
			array = append(array, v...)
		}
	}

cancelled:
	b, err := lang.MarshalData(a.p, a.dataType, array)
	if err != nil {
		return true, err
	}
	_, err = a.p.Stdout.Write(b)
	return true, err
}
