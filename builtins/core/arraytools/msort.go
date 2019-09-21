package arraytools

import (
	"sort"

	"github.com/lmorg/murex/lang"
)

func init() {
	lang.GoFunctions["msort"] = cmdMSort
}

func cmdMSort(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	var a []string

	err := p.Stdin.ReadArray(func(b []byte) {
		if p.HasCancelled() {
			return
		}

		a = append(a, string(b))
	})

	if err != nil {
		return err
	}

	sort.Strings(a)

	b, err := lang.MarshalData(p, dt, a)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
