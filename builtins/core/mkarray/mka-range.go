package mkarray

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	rxEngAlphaLower = regexp.MustCompile(`^[a-z]$`)
	rxEngAlphaUpper = regexp.MustCompile(`^[A-Z]$`)
	rxAltNumberBase = regexp.MustCompile(`^([a-zA-Z0-9]+)\.\.([a-zA-Z0-9]+)[\.x]([0-9]+)$`)
)

func rangeToArray(b []byte) ([]string, error) {
	split := strings.Split(string(b), "..")
	if len(split) > 2 {
		return nil, fmt.Errorf("Invalid syntax. Too many double periods, `..`, in range`%s`. Please escape periods, `\\.`, if you wish to include period in your range", string(b))
	}

	if len(split) < 2 {
		return nil, fmt.Errorf("Invalid syntax. Range periods, `..`, found but cannot determine start and end range in `%s`", string(b))
	}

	i1, e1 := strconv.Atoi(split[0])
	i2, e2 := strconv.Atoi(split[1])

	if e1 == nil && e2 == nil {
		switch {
		case i1 < i2:
			a := make([]string, i2-i1+1)
			if split[0][0] != '0' {
				for i := range a {
					a[i] = strconv.Itoa(i + i1)
				}
			} else {
				l := len(split[0])
				s := "%0" + strconv.Itoa(l) + "d"
				for i := range a {
					a[i] = fmt.Sprintf(s, i+i1)
				}
			}
			return a, nil

		case i1 > i2:
			a := make([]string, i1-i2+1)
			if split[1][0] != '0' {
				for i := range a {
					a[i] = strconv.Itoa(i1 - i)
				}
			} else {
				l := len(split[1])
				s := "%0" + strconv.Itoa(l) + "d"
				for i := range a {
					a[i] = fmt.Sprintf(s, i1-i)
				}
			}
			return a, nil

		default:
			return nil, fmt.Errorf("Invalid range. Start and end of range are the same in `%s`", string(b))
		}
	}

	if rxEngAlphaLower.MatchString(split[0]) && rxEngAlphaLower.MatchString(split[1]) {
		switch {
		case split[0] < split[1]:
			a := make([]string, 0)
			for i := []byte(split[0])[0]; i <= []byte(split[1])[0]; i++ {
				a = append(a, string([]byte{i}))
			}
			return a, nil
		case split[0] > split[1]:
			a := make([]string, 0)
			for i := []byte(split[0])[0]; i >= []byte(split[1])[0]; i-- {
				a = append(a, string([]byte{i}))
			}
			return a, nil
		default:
			return nil, fmt.Errorf("Invalid range. Start and end of range are the same in `%s`", string(b))
		}
	}

	if rxEngAlphaUpper.MatchString(split[0]) && rxEngAlphaUpper.MatchString(split[1]) {
		switch {
		case split[0] < split[1]:
			a := make([]string, 0)
			for i := []byte(split[0])[0]; i <= []byte(split[1])[0]; i++ {
				a = append(a, string([]byte{i}))
			}
			return a, nil
		case split[0] > split[1]:
			a := make([]string, 0)
			for i := []byte(split[0])[0]; i >= []byte(split[1])[0]; i-- {
				a = append(a, string([]byte{i}))
			}
			return a, nil
		default:
			return nil, fmt.Errorf("Invalid range. Start and end of range are the same in `%s`", string(b))
		}

	}

	if rxAltNumberBase.Match(b) {
		split = rxAltNumberBase.FindStringSubmatch(string(b))
		base, err := strconv.Atoi(split[3])
		if err != nil {
			return nil, errors.New("Unable to determin number base: " + err.Error())
		}
		if base < 2 || base > 36 {
			return nil, errors.New("Number base must be between 2 and 36 (inclusive)")
		}

		i1, err := strconv.ParseInt(split[1], base, 64)
		if err != nil {
			return nil, errors.New("Unable to determin start of range: " + err.Error())
		}

		i2, err := strconv.ParseInt(split[2], base, 64)
		if err != nil {
			return nil, errors.New("Unable to determin end of range: " + err.Error())
		}

		switch {
		case i1 < i2:
			a := make([]string, i2-i1+1)
			if split[0][0] != '0' {
				for i := range a {
					a[i] = strconv.FormatInt(i1+int64(i), base)
				}
			} else {
				l := len(split[1])
				s := "%0" + strconv.Itoa(l) + "s"
				for i := range a {
					a[i] = fmt.Sprintf(s, strconv.FormatInt(i1+int64(i), base))
				}
			}
			return a, nil

		case i1 > i2:
			a := make([]string, i1-i2+1)
			if split[1][0] != '0' {
				for i := range a {
					a[i] = strconv.FormatInt(i1-int64(i), base)
				}
			} else {
				l := len(split[1])
				s := "%0" + strconv.Itoa(l) + "s"
				for i := range a {
					a[i] = fmt.Sprintf(s, strconv.FormatInt(i1-int64(i), base))
				}
			}
			return a, nil
		default:
			return nil, fmt.Errorf("Invalid range. Start and end of range are the same in `%s`", string(b))
		}
	}

	// Mapped lists. See consts.go
	c := getCase(split[0])
	start := strings.ToLower(split[0])
	end := strings.ToLower(split[1])
	for i := range mapRanges {
		matched, array := mapArray(mapRanges[i][start], mapRanges[i][end], mapRanges[i], c)
		if matched {
			return array, nil
		}
	}

	return nil, fmt.Errorf("Unable to auto-detect range in `%s`", string(b))
}

func mapArray(start, end int, constMap map[string]int, c int) (matched bool, array []string) {
	if start == 0 || end == 0 {
		return
	}

	matched = true

	consts := make([]string, len(constMap))
	for s, i := range constMap {
		consts[i-1] = setCase(s, c)
	}

	switch {
	case start < end:
		array = consts[start-1 : end]

	case start >= end:
		array = consts[start-1:]
		array = append(array, consts[:end]...)
	}

	return
}
