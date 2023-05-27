package ranges

import (
	"fmt"
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

func (rf *rfIndex) SetLength(i int) {
	rf.start += i
	rf.end += i
}

func createRfIndex(r *rangeParameters) (*rfIndex, error) {
	rf := new(rfIndex)

	sStart := r.Start
	sEnd := r.End

	if sStart == "" {
		sStart = "0"
	}

	if sEnd == "" {
		sEnd = "-1"
	}

	var err error

	rf.start, err = strconv.Atoi(sStart)
	if err != nil {
		return nil, fmt.Errorf("cannot convert start value to integer: %s", err.Error())
	}

	rf.end, err = strconv.Atoi(sEnd)
	if err != nil {
		return nil, fmt.Errorf("cannot convert end value to integer: %s", err.Error())
	}

	if rf.start < 0 {
		r.Buffer = true
		if r.End == "" {
			rf.end = 1
		}
	}

	if rf.start > 0 && !r.Exclude {
		rf.end++
	}

	return rf, nil
}

func newIndex(r *rangeParameters) error {
	rf, err := createRfIndex(r)
	if err != nil {
		return err
	}

	rf.start -= 1
	rf.end -= 1

	r.Match = rf

	return nil
}
