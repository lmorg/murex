package httpclient

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/lmorg/murex/builtins/pipes/file"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi/codes"
	"github.com/lmorg/murex/utils/humannumbers"
	"github.com/lmorg/murex/utils/readline"
)

func cmdGetFile(p *lang.Process) (err error) {
	if p.Parameters.Len() == 0 {
		return errors.New("URL required")
	}

	url, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	validateURL(&url, p.Config)

	var body stdio.Io
	if p.IsMethod {
		body = p.Stdin
	} else {
		body = nil
	}

	resp, err := Request(p.Context, "GET", url, body, p.Config, disableTimeout)
	if err != nil {
		return err
	}

	filename := utils.ExtractFileNameFromURL(url)

	p.Stdout.SetDataType(lang.RequestMetadataToMurex(resp.Header.Get(("Content-Type")), filename))

	quit := make(chan bool)
	cl := resp.Header.Get("Content-Length")
	if p.Stdout.IsTTY() {
		p.Stdout, err = file.NewFile(filename)
		if err != nil {
			return err
		}
		p.Stdout.Open()
	}

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
			"%sDownloaded %s to %s\n",
			"\x1b["+strconv.Itoa(readline.GetTermWidth()+2)+"D"+codes.ClearLine+codes.Reset,
			humannumbers.Bytes(written),
			filename,
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
				time.Sleep(2 * time.Second)
				written, _ = p.Stdout.Stats()
				speed = (written - last) * 2000
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

	s := "\x1b[" + strconv.Itoa(width+2) + "D" + codes.ClearLine + codes.Reset
	if cells > 0 {
		s += codes.Invert
	}

	for i := 0; i < width; i++ {
		if cells+1 == i {
			s += codes.Reset
		}

		if i < len(message) {
			s += string([]byte{message[i]})
		} else {
			s += " "
		}
	}

	os.Stderr.WriteString(s + codes.Reset)
}
