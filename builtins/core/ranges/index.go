package ranges

import (
	"errors"
	"strconv"
)

type rfIndex struct {
	start int
	end   int
	i     int
}

func (rf *rfIndex) Start(_ []byte) bool {
	rf.i++
	return rf.i > rf.start
}

func (rf *rfIndex) End(_ []byte) bool {
	if rf.end > -1 {
		rf.i++
		return rf.i > rf.end
	}
	return false
}

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
		return errors.New("Cannot convert start value to integer: " + err.Error())
	}

	rf.end, err = strconv.Atoi(sEnd)
	if err != nil {
		return errors.New("Cannot convert end value to integer: " + err.Error())
	}

	if rf.start > 0 && !r.Exclude {
		rf.end++
	}

	r.Match = rf

	return
}
