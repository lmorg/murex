package httpclient

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"io/ioutil"
	"strconv"
)

func cmdPost(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)

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

	var body io.Reader
	if p.IsMethod {
		body = p.Stdin
	} else {
		body = nil
	}

	resp, err := request("POST", url, body, enableTimeout)
	if err != nil {
		return err
	}

	jhttp.Status.Code, _ = strconv.Atoi(resp.Status[:3])
	jhttp.Status.Message = resp.Status[4:]

	jhttp.Headers = resp.Header
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	jhttp.Body = string(b)
	if err != nil {
		return err
	}

	b, err = utils.JsonMarshal(jhttp, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	p.Stdout.Write(b)

	return nil
}
