package semver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func raiseError(err error, context string, values ...any) (*Version, error) {
	if err != nil {
		return nil, errors.New(fmt.Sprintf(context, values...) + ": " + err.Error())
	}
	return nil, fmt.Errorf(context, values...)
}

// Parse takes a version string with the following pattern: major.minor.patch
// and returns its component parts
func Parse(s string) (*Version, error) {
	if len(s) == 0 {
		return nil, errors.New("empty version string")
	}

	var err error
	ver := new(Version)

	split := strings.Split(s, ".")

	switch len(split) {
	case 3:
		ver.Patch, err = strconv.Atoi(split[2])
		if err != nil {
			return raiseError(err, "patch value is not a number")
		}
		fallthrough

	case 2:
		ver.Minor, err = strconv.Atoi(split[1])
		if err != nil {
			return raiseError(err, "minor value is not a number")
		}
		fallthrough

	case 1:
		ver.Major, err = strconv.Atoi(split[0])
		if err != nil {
			return raiseError(err, "major value is not a number")
		}

	default:
		return raiseError(nil, "too many full stops / period or not a valid version string")
	}

	return ver, err
}

func Compare(version string, comparison string) (bool, error) {
	cond, comp := parseComparison(comparison)
	cond = strings.TrimSpace(cond)
	comp = strings.TrimSpace(comp)

	ver, err := Parse(version)
	if err != nil {
		return false, fmt.Errorf("cannot parse version string: '%s'", err.Error())
	}

	compV, err := parseCompVersion(comp)
	if err != nil {
		return false, fmt.Errorf("cannot parse comparison string: '%s'", err.Error())
	}

	switch cond {
	case ">":
		return compare(ver, compV) == greaterThan, nil
	case ">=":
		return compare(ver, compV) != lessThan, nil
	case "", "=", "==":
		return compare(ver, compV) == equalTo, nil
	case "<=":
		return compare(ver, compV) != greaterThan, nil
	case "<":
		return compare(ver, compV) == lessThan, nil
	default:
		return false, fmt.Errorf("unknown comparison token '%s'", cond)
	}
}

func parseComparison(comparison string) (string, string) {
	for i := range comparison {
		switch {
		case comparison[i] == ' ':
			continue
		case comparison[i] <= '9' && '0' <= comparison[i]:
			return comparison[:i], comparison[i:]
		}
	}
	return "", ""
}

func parseCompVersion(s string) (*Version, error) {
	if len(s) == 0 {
		return nil, errors.New("empty version string")
	}

	var err error
	ver := new(Version)

	split := strings.Split(s, ".")

	switch len(split) {
	case 3:
		ver.Patch, err = strconv.Atoi(split[2])
		if err != nil {
			return raiseError(err, "patch value is not a number")
		}
		ver.Patch += 2
		fallthrough

	case 2:
		ver.Patch--
		ver.Minor, err = strconv.Atoi(split[1])
		if err != nil {
			return raiseError(err, "minor value is not a number")
		}
		ver.Minor++
		fallthrough

	case 1:
		ver.Major, err = strconv.Atoi(split[0])
		if err != nil {
			return raiseError(err, "major value is not a number")
		}
		ver.Patch--
		ver.Minor--

	default:
		return raiseError(nil, "too many full stops / period or not a valid version string")
	}

	return ver, err
}

const (
	lessThan    = -1
	equalTo     = 0
	greaterThan = 1
)

func compare(version *Version, comparison *Version) int {
	switch {
	case version.Major > comparison.Major:
		return greaterThan
	case version.Major < comparison.Major:
		return lessThan

	case comparison.Minor < 0:
		return equalTo

	case version.Minor > comparison.Minor:
		return greaterThan
	case version.Minor < comparison.Minor:
		return lessThan

	case comparison.Patch < 0:
		return equalTo

	case version.Patch > comparison.Patch:
		return greaterThan
	case version.Patch < comparison.Patch:
		return lessThan

	default:
		return equalTo
	}
}
