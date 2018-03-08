package arange

import (
	"regexp"
)

type rfRegexp struct {
	rxStart *regexp.Regexp
	rxEnd   *regexp.Regexp
}

func (rf *rfRegexp) Start(b []byte) bool { return rf.rxStart.Match(b) }
func (rf *rfRegexp) End(b []byte) bool   { return rf.rxEnd.Match(b) }

func NewRegexp(r *rangeParameters) (err error) {
	rf := new(rfRegexp)

	rf.rxStart, err = regexp.Compile(r.Start)
	if err != nil {
		return err
	}

	rf.rxEnd, err = regexp.Compile(r.End)
	if err != nil {
		return err
	}

	r.Match = rf

	return nil
}
