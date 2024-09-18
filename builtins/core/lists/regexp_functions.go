package lists

import (
	"fmt"
	"regexp"

	"github.com/lmorg/murex/lang"
)

func regexMatch(p *lang.Process, rx *regexp.Regexp, dt string, withHeading bool) error {
	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	var (
		count int
		fn    func([]byte)
	)

	callback := func(b []byte) { fn(b) }

	readArray := func(b []byte) {
		matched := rx.Match(b)
		if (matched && !p.IsNot) || (!matched && p.IsNot) {

			count++
			err = aw.Write(b)
			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}

		}
	}

	readHeading := func(b []byte) {
		err = aw.Write(b)
		if err != nil {
			p.Stdin.ForceClose()
			p.Done()
		}
		fn = readArray
	}

	// this is a bit of a kludge to allow us to read the heading and run one
	// type of function against it, but then process the rest of the array
	// differently. The aim of this is to avoid having `if` blocks inside the
	// `ReadArray()` loop.
	if withHeading {
		fn = readHeading
	} else {
		callback = readArray
	}
	p.Stdin.ReadArray(p.Context, callback)

	if p.HasCancelled() {
		return err
	}

	if count == 0 {
		return fmt.Errorf("nothing matched: %s", rx.String())
	}

	return aw.Close()
}

func regexSubstitute(p *lang.Process, rx *regexp.Regexp, sRegex []string, dt string) error {
	if len(sRegex) < 3 {
		return fmt.Errorf("invalid regex: too few parameters\nexpecting s/find/substitute/ in: `%s`", p.Parameters.StringAll())
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	sub := []byte(sRegex[2])

	p.Stdin.ReadArray(p.Context, func(b []byte) {
		err = aw.Write(rx.ReplaceAll(b, sub))
		if err != nil {
			p.Stdin.ForceClose()
			p.Done()
		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}

func regexFind(p *lang.Process, rx *regexp.Regexp, dt string) error {
	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(p.Context, func(b []byte) {
		match := rx.FindAllStringSubmatch(string(b), -1)
		for _, found := range match {
			if len(found) > 1 {

				for i := 1; i < len(found); i++ {
					err = aw.WriteString(found[i])
					if err != nil {
						p.Stdin.ForceClose()
						p.Done()
					}

				}

			}
		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}
