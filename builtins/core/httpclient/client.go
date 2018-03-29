package httpclient

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	net_url "net/url"
	"regexp"
	"time"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
)

const (
	enableTimeout  = true
	disableTimeout = false
)

var rxHttpProto = regexp.MustCompile(`(?i)^http(s)?://`)

// Request generates a HTTP request
func Request(method, url string, body io.Reader, conf *config.Config, setTimeout bool, contentType string) (response *http.Response, err error) {
	toStr, err := conf.Get("http", "timeout", types.String)
	if err != nil {
		return
	}
	toDur, err := time.ParseDuration(toStr.(string) + "s")
	if err != nil {
		return
	}

	insecure, err := conf.Get("http", "insecure", types.Boolean)
	if err != nil {
		return
	}

	client := &http.Client{}
	if setTimeout {
		tr := http.Transport{
			Dial:            dialTimeout(toDur, toDur, conf),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure.(bool)},
		}

		client = &http.Client{
			Timeout:   toDur,
			Transport: &tr,
		}

	} else {
		tr := http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure.(bool)},
		}

		client = &http.Client{
			Transport: &tr,
		}
	}

	userAgent, err := conf.Get("http", "user-agent", types.String)
	if err != nil {
		return
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	request.Header.Set("User-Agent", userAgent.(string))

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	urlParsed, err := net_url.Parse(url)
	if err != nil {
		return
	}

	request.Header.Set("Host", urlParsed.Host)

	redirects, err := conf.Get("http", "redirect", types.Boolean)
	if err != nil {
		return
	}

	if redirects.(bool) {
		response, err = client.Do(request)
	} else {
		response, err = client.Transport.RoundTrip(request)
	}

	return
}

// Code unashamedly copy and pasted from:
// https://stackoverflow.com/questions/16895294/how-to-set-timeout-for-http-get-requests-in-golang#16930649
func dialTimeout(cTimeout time.Duration, rwTimeout time.Duration, conf *config.Config) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func validateURL(url *string, conf *config.Config) {
	if !rxHttpProto.MatchString(*url) {
		v, err := conf.Get("http", "default-https", types.Boolean)
		if err != nil {
			v = false
		}

		if v.(bool) {
			*url = "https://" + *url
			return
		}

		*url = "http://" + *url
	}
}
