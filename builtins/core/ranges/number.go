package ranges

func newNumber(r *rangeParameters) (err error) {
	rf, err := createRfIndex(r)
	if err != nil {
		return err
	}

	r.Match = rf
	return nil
}
