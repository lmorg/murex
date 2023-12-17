package jsonconcat

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func iface2Str(a []interface{}) []string {
	s := make([]string, len(a))
	for i := range a {
		s[i] = fmt.Sprint(a[i])
	}
	return s
}

func index(p *lang.Process, params []string) error {
	// if it's purely a numeric param then well index by table (ie each row is
	// an array index). Otherwise we'd do our best to splice things into the
	// table
	for i := range params {
		for _, b := range params[i] {
			if b < '0' || b > '9' {
				return indexTable(p, params)
			}
		}
	}

	return indexObject(p, params)
}

func indexObject(p *lang.Process, params []string) error {
	lines := make(map[int]bool)
	for i := range params {
		num, err := strconv.Atoi(params[i])
		if err != nil {
			return fmt.Errorf("parameter, `%s`, isn't an integer. %s", params[i], err)
		}
		lines[num] = true
	}

	var (
		i   int
		err error
	)

	err = p.Stdin.ReadArray(p.Context, func(b []byte) {
		if lines[i] != p.IsNot {
			_, err = p.Stdout.Writeln(b)
			if err != nil {
				return
			}
		}
		i++
	})

	return err
}

func indexTable(p *lang.Process, params []string) error {
	cRecords := make(chan []string, 10)
	status := make(chan error)

	go func() {
		err1 := p.Stdin.ReadArray(p.Context, func(b []byte) {
			var v []interface{}
			err2 := json.Unmarshal(b, &v)
			if err2 != nil {
				status <- err2
				return
			}
			cRecords <- iface2Str(v)
		})
		if err1 != nil {
			close(cRecords)
			status <- err1
			return
		}
		close(cRecords)
	}()

	marshaller := func(s []string) []byte {
		b, err3 := lang.MarshalData(p, types.Json, s)
		if err3 != nil {
			close(cRecords)
			status <- err3
		}
		return b
	}

	go func() {
		err4 := lang.IndexTemplateTable(p, params, cRecords, marshaller)
		status <- err4
	}()

	return <-status
}
