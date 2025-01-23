package httpclient

import (
	"errors"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/modver"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/readall"
	"github.com/lmorg/murex/utils/semver"
)

func cmdGet(p *lang.Process) error  { return request(p, "GET") }
func cmdPost(p *lang.Process) error { return request(p, "POST") }

func request(p *lang.Process, method string) (err error) {
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

	var body stdio.Io
	if p.IsMethod {
		body = p.Stdin
	} else {
		body = nil
	}

	resp, err := Request(p.Context, method, url, body, p.Config, enableTimeout)
	if err != nil {
		return err
	}

	jHttp.Status.Code, _ = strconv.Atoi(resp.Status[:3])
	jHttp.Status.Message = resp.Status[4:]

	jHttp.Headers = resp.Header

	b, err := readall.ReadAll(p.Context, resp.Body)
	closeErr := resp.Body.Close()

	if err != nil {
		return err
	}
	if closeErr != nil {
		return err
	}

	jHttp.Body = convertBodyToObj(p, b, resp.Header.Get("Content-Type"))

	b, err = json.Marshal(jHttp, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func convertBodyToObj(p *lang.Process, b []byte, mime string) any {
	if modver.Get(p.FileRef.Source.Module).Compare(semver.Version7_0).IsLessThan() {
		return string(b)
	}

	dataType := lang.MimeToMurex(mime)

	switch dataType {
	case types.String, types.Generic, types.Binary:
		return string(b)

	default:
		v, err := lang.UnmarshalDataBuffered(p, b, dataType)
		if err != nil {
			return string(b)
		}
		return v
	}
}
