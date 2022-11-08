package ranges

import (
	"errors"
	"strconv"
)

func newNumber(r *rangeParameters) (err error) {
	rf := new(rfIndex)

	sStart := r.Start
	sEnd := r.End

	if sStart == "" {
		sStart = "0"
	}

	if sEnd == "" {
		sEnd = "-1"
	}

	rf.start, err = strconv.Atoi(sStart)
	if err != nil {
		return errors.New("cannot convert start value to integer: " + err.Error())
	}

	rf.end, err = strconv.Atoi(sEnd)
	if err != nil {
		return errors.New("cannot convert end value to integer: " + err.Error())
	}

	if rf.start > 0 && !r.Exclude {
		rf.end++
	}

	r.Match = rf

	return
}
