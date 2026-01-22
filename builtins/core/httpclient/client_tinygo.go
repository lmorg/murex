//go:build tinygo
// +build tinygo

package httpclient

import (
	"net/http"

	"github.com/lmorg/murex/config"
)

func createClient(conf *config.Config, setTimeout bool) (*http.Client, error) {
	return &http.Client{}, nil
}
