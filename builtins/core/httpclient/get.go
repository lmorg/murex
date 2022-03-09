package httpclient

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/humannumbers"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/readall"
	"github.com/lmorg/murex/utils/readline"
)

func cmdGet(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)

	if p.Parameters.Len() == 0 {
		return errors.New("URL required")
	}

	var jHttp jsonHttp

	url, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	validateURL(&url, p.Config)

	var body io.Reader
	if p.IsMethod {
		body = p.Stdin
	} else {
		body = nil
	}

	resp, err := Request(p.Context, "GET", url, body, p.Config, enableTimeout)
	if err != nil {
		return err
	}

	jHttp.Status.Code, _ = strconv.Atoi(resp.Status[:3])
	jHttp.Status.Message = resp.Status[4:]

	jHttp.Headers = resp.Header

	b, err := readall.ReadAll(p.Context, resp.Body)
	resp.Body.Close()
	jHttp.Body = string(b)
	if err != nil {
		return err
	}

	b, err = json.Marshal(jHttp, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdGetFile(p *lang.Process) (err error) {
	if p.Parameters.Len() == 0 {
		return errors.New("URL required")
	}

	url, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	validateURL(&url, p.Config)

	var body io.Reader
	if p.IsMethod {
		body = p.Stdin
	} else {
		body = nil
	}

	resp, err := Request(p.Context, "GET", url, body, p.Config, disableTimeout)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(lang.MimeToMurex(resp.Header.Get("Content-Type")))

	quit := make(chan bool)
	cl := resp.Header.Get("Content-Length")

	var i int
	if cl == "" {
		cl = "{unknown}"
	} else {
		i, _ = strconv.Atoi(cl)
		cl = humannumbers.Bytes(uint64(i))
	}

	defer func() {
		quit <- true
		resp.Body.Close()
		written, _ := p.Stdout.Stats()

		os.Stderr.WriteString(fmt.Sprintf(
			"%sDownloaded %s.\n",
			"\x1b["+strconv.Itoa(readline.GetTermWidth()+2)+"D"+ansi.ClearLine+ansi.Reset,
			humannumbers.Bytes(written),
		))
	}()

	go func() {
		var last, written, speed uint64
		select {
		case <-quit:
			return
		default:
		}

		for {

			if p.Stderr.IsTTY() {
				time.Sleep(10 * time.Millisecond)
				written, _ = p.Stdout.Stats()
				speed = (written - last) * 100
			} else {
				time.Sleep(2 * time.Millisecond)
				written, _ = p.Stdout.Stats()
				speed = (written - last) / 2
			}
			last = written

			select {
			case <-quit:
				return
			default:
			}

			msg := fmt.Sprintf(
				"Downloading... %s of %s @ %s/s....",
				humannumbers.Bytes(written),
				cl,
				humannumbers.Bytes(speed),
			)
			printGaugeBar(float64(written), float64(i), msg)
		}
	}()

	_, err = io.Copy(p.Stdout, resp.Body)
	return err
}

func printGaugeBar(value, max float64, message string) {
	width := readline.GetTermWidth()
	cells := int((float64(width) / max) * value)

	s := "\x1b[" + strconv.Itoa(width+2) + "D" + ansi.ClearLine + ansi.Reset
	if cells > 0 {
		s += ansi.Invert
	}

	for i := 0; i < width; i++ {
		if cells+1 == i {
			s += ansi.Reset
		}

		if i < len(message) {
			s += string([]byte{message[i]})
		} else {
			s += " "
		}
	}

	os.Stderr.WriteString(s + ansi.Reset)
}
