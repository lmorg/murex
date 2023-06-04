package ranges

func newNumber(r *rangeParameters) (err error) {
	rf, err := createRfIndex(r)
	if err != nil {
		return err
	}

	if rf.start < 0 {
		rf.start -= 2
		rf.end -= 2
	}

	r.Match = rf
	return nil
}
