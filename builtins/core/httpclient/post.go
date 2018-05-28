package httpclient

import (
	"errors"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
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
	validateURL(&url, p.Config)

	var body io.Reader
	var contentType string
	if p.IsMethod {
		body = p.Stdin

		contentType, err = p.Parameters.String(1)
		if err != nil {
			return err
		}
	} else {
		body = nil
	}

	resp, err := Request("POST", url, body, p.Config, enableTimeout, contentType)
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

	b, err = json.Marshal(jhttp, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)

	return err
}
