package httpclient

import (
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/readall"
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

	if cl == "" {
		cl = "{unknown}"
	} else {
		i, _ := strconv.Atoi(cl)
		cl = utils.HumanBytes(uint64(i))
		//length = float64(i)
	}

	defer func() {
		quit <- true
		resp.Body.Close()
		written, _ := p.Stdout.Stats()
		os.Stderr.WriteString("Downloaded " + utils.HumanBytes(written) + ".\n")
	}()

	go func() {
		//gauge := render.NewGaugeBar("Downloading....")
		var last uint64
		select {
		case <-quit:
			return
		default:
		}

		for {
			time.Sleep(1 * time.Second)

			select {
			case <-quit:
				return
			default:
			}

			written, _ := p.Stdout.Stats()
			speed := written - last
			last = written
			//percent := int((float64(written) / length) * 100)
			os.Stderr.WriteString("Downloaded " + utils.HumanBytes(written) + " of " + cl + " @ " + utils.HumanBytes(speed) + "/s....\n")
			//render.UpdateGaugeBar(gauge, percent, "Downloaded "+utils.HumanBytes(written)+" of "+cl+" @ "+utils.HumanBytes(speed)+"/s....")
		}
	}()

	_, err = io.Copy(p.Stdout, resp.Body)
	return err
}
