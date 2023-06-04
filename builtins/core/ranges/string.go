package ranges

type rfString struct {
	sStart string
	sEnd   string
}

func (rf *rfString) Start(b []byte) bool { return string(b) == rf.sStart }
func (rf *rfString) End(b []byte) bool   { return string(b) == rf.sEnd }
func (rf *rfString) SetLength(_ int)     { /* do nothing */ }

func newString(r *rangeParameters) error {
	rf := new(rfString)

	rf.sStart = r.Start
	rf.sEnd = r.End

	r.Match = rf

	return nil
}
