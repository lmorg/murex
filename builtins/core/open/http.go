package open

import (
	"io"

	"github.com/lmorg/murex/builtins/core/httpclient"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
)

func http(p *proc.Process, url string) (io.ReadCloser, string, error) {
	resp, err := httpclient.Request(p.Context, "GET", url, nil, p.Config, true)

	if err != nil {
		return nil, "", err
	}

	dt := define.MimeToMurex(resp.Header.Get("Content-Type"))

	// TODO: insert something about content-length detection

	return resp.Body, dt, nil
}
