package mkarray

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	rxEngAlphaLower = regexp.MustCompile(`^[a-z]$`)
	rxEngAlphaUpper = regexp.MustCompile(`^[A-Z]$`)
	rxAltNumberBase = regexp.MustCompile(`^([a-zA-Z0-9]+)\.\.([a-zA-Z0-9]+)[\.x]([0-9]+)$`)
)

func rangeToArray(b []byte) ([]string, error) {
	split := strings.Split(string(b), "..")
	if len(split) > 2 {
		return nil, fmt.Errorf("invalid syntax. Too many double periods, `..`, in range`%s`. Please escape periods, `\\.`, if you wish to include period in your range", string(b))
	}

	if len(split) < 2 {
		return nil, fmt.Errorf("invalid syntax. Range periods, `..`, found but cannot determine start and end range in `%s`", string(b))
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
			a := make([]string, 1)
			if split[1][0] != '0' {
				a[0] = strconv.Itoa(i1)
			} else {
				l := len(split[1])
				s := "%0" + strconv.Itoa(l) + "d"
				a[0] = fmt.Sprintf(s, i1)
			}
			return a, nil
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
			return []string{split[0]}, nil
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
			return []string{split[0]}, nil
		}

	}

	if rxAltNumberBase.Match(b) {
		split = rxAltNumberBase.FindStringSubmatch(string(b))
		base, err := strconv.Atoi(split[3])
		if err != nil {
			return nil, errors.New("unable to determin number base: " + err.Error())
		}
		if base < 2 || base > 36 {
			return nil, errors.New("number base must be between 2 and 36 (inclusive)")
		}

		i1, err := strconv.ParseInt(split[1], base, 64)
		if err != nil {
			return nil, errors.New("unable to determin start of range: " + err.Error())
		}

		i2, err := strconv.ParseInt(split[2], base, 64)
		if err != nil {
			return nil, errors.New("unable to determin end of range: " + err.Error())
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
			a := make([]string, 1)
			if split[1][0] != '0' {
				a[0] = strconv.FormatInt(i1, base)
			} else {
				l := len(split[1])
				s := "%0" + strconv.Itoa(l) + "s"
				a[0] = fmt.Sprintf(s, strconv.FormatInt(i1, base))
			}
			return a, nil
		}
	}

	var t1, t2 time.Time
	for i := range dateFormat {
		t1, e1 = time.Parse(dateFormat[i], split[0])
		if e1 == nil || (len(split[0]) == 0 && len(split[1]) > 0) {
			t2, e2 = time.Parse(dateFormat[i], split[1])

			if e2 == nil || (len(split[0]) > 0 && len(split[1]) == 0) {
				var c int
				switch {
				case len(split[0]) == 0:
					t1, e1 = time.Parse(dateFormat[i], time.Now().Format(dateFormat[i]))
					if e1 != nil {
						return nil, e1
					}
					c = getCase(split[1])
				case len(split[1]) == 0:
					t2, e2 = time.Parse(dateFormat[i], time.Now().Format(dateFormat[i]))
					if e2 != nil {
						return nil, e2
					}
					c = getCase(split[0])
				default:
					c = getCase(split[0])
				}

				switch {
				case t1.Before(t2):
					a := []string{setCase(t1.Format(dateFormat[i]), c)}
					for t1.Before(t2) {
						t1 = t1.AddDate(0, 0, 1)
						a = append(a, setCase(t1.Format(dateFormat[i]), c))
					}
					return a, nil

				case t1.After(t2):
					a := []string{setCase(t1.Format(dateFormat[i]), c)}
					for t1.After(t2) {
						t1 = t1.AddDate(0, 0, -1)
						a = append(a, setCase(t1.Format(dateFormat[i]), c))
					}
					return a, nil

				default:
					return []string{setCase(t1.Format(dateFormat[i]), c)}, nil
					//return nil, fmt.Errorf("invalid range. Start and end of range are the same in `%s`", string(b))
				}
			}
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

	return nil, fmt.Errorf("unable to auto-detect range in `%s`", string(b))
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
