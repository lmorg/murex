package httpclient

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func cmdGet(p *proc.Process) (err error) {
	if p.Parameters.Len() == 0 {
		return errors.New("URL required.")
	}

	var jhttp jsonHttp

	url, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	if !rxHttpProto.MatchString(url) {
		url = "http://" + url
	}

	debug.Log("resp, err := http.Get(url)....")
	//resp, err := http.Get(url)
	resp, err := client.Get(url)
	debug.Log("resp, err := http.Get(url)!!!!")
	if err != nil {
		return err
	}
	debug.Log("jhttp.Status.Code, _ = strconv.Atoi(resp.Status[:3])....")
	jhttp.Status.Code, _ = strconv.Atoi(resp.Status[:3])
	jhttp.Status.Message = resp.Status[4:]

	jhttp.Headers = resp.Header
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	jhttp.Body = string(b)
	if err != nil {
		return err
	}

	debug.Log(`b, err = json.MarshalIndent(jhttp, "", "\t")....`)
	b, err = json.MarshalIndent(jhttp, "", "\t")
	if err != nil {
		return err
	}

	p.Stdout.Write(b)
	p.Stdout.SetDataType(types.Json)

	return nil
}

func cmdGetFile(p *proc.Process) (err error) {
	if p.Parameters.Len() == 0 {
		return errors.New("URL required.")
	}

	url, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	if !rxHttpProto.MatchString(url) {
		url = "http://" + url
	}

	//resp, err := http.Get(url)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	quit := false
	cl := resp.Header.Get("Content-Length")

	defer p.Stdout.SetDataType(types.MimeToMurex(resp.Header.Get("Content-Type")))

	if cl == "" {
		cl = "{unknown}"
	} else {
		i, _ := strconv.Atoi(cl)
		cl = utils.HumanBytes(uint64(i))
		//length = float64(i)
	}

	defer func() {
		quit = true
		resp.Body.Close()
		written, _ := p.Stdout.Stats()
		os.Stderr.WriteString("Downloaded " + utils.HumanBytes(written) + ".\n")
	}()

	go func() {
		//gauge := render.NewGaugeBar("Downloading....")
		var last uint64
		for !quit {
			time.Sleep(1 * time.Second)
			if quit {
				return
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
	if err != nil {
		return err
	}

	return nil
}
