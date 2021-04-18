// +build windows

package textmanip

/*func init() {
	lang.GoFunctions["printf"] = cmdPretty
}

func cmdSprintf(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.Len() == 0 {
		return errors.New("Parameters missing")
	}

	s := p.Parameters.StringAll()
	var a []interface{}

	err := p.Stdin.ReadArray(func(b []byte) {
		a = append(a, string(b))
	})

	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(fmt.Sprintf(s, a...)))
	return err
}
*/
