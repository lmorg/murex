package csv

import (
	"bytes"
	enc "encoding/csv"
	"io"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func readIndex(p *lang.Process, params []string) error {
	cRecords := make(chan []string, 1)
	status := make(chan error)

	r := enc.NewReader(p.Stdin)

	v, err := p.Config.Get("csv", "separator", types.String)
	if err != nil {
		return err
	}
	if len(v.(string)) != 0 {
		r.Comma = rune(v.(string)[0])
	}

	v, err = p.Config.Get("csv", "comment", types.String)
	if err != nil {
		return err
	}
	if len(v.(string)) != 0 {
		r.Comment = rune(v.(string)[0])
	}

	go func() {
		for {
			recs, err := r.Read()
			switch {
			case err == io.EOF:
				close(cRecords)
				status <- nil
				return

			case err != nil && strings.HasSuffix(err.Error(), enc.ErrFieldCount.Error()):
				fallthrough

			case err == nil:
				cRecords <- recs
				continue

			default:
				close(cRecords)
				status <- err
				return
			}

		}
	}()

	marshaller := func(s []string) []byte {
		var b []byte
		buf := bytes.NewBuffer(b)
		w := enc.NewWriter(buf)
		w.Comma = r.Comma
		err := w.Write(s)
		if err != nil {
			// this shouldn't ever happen anyway
			panic(err)
		}
		w.Flush()
		if w.Error() != nil {
			// this shouldn't ever happen anyway
			panic(err)
		}

		return utils.CrLfTrim(buf.Bytes())
	}

	go func() {
		err := lang.IndexTemplateTable(p, params, cRecords, marshaller)
		status <- err
		return
	}()

	return <-status
}
