//go:build !tinygo
// +build !tinygo

package httpclient

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
)

func createClient(conf *config.Config, setTimeout bool) (*http.Client, error) {
	toStr, err := conf.Get("http", "timeout", types.String)
	if err != nil {
		return nil, err
	}
	toDur, err := time.ParseDuration(toStr.(string) + "s")
	if err != nil {
		return nil, err
	}

	insecure, err := conf.Get("http", "insecure", types.Boolean)
	if err != nil {
		return nil, err
	}

	if setTimeout {
		tr := http.Transport{
			Dial:            dialTimeout(toDur, toDur),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure.(bool)},
		}

		return &http.Client{Timeout: toDur, Transport: &tr}, nil
	}

	tr := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure.(bool)},
	}

	return &http.Client{Transport: &tr}, nil
}
