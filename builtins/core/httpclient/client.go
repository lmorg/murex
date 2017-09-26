package httpclient

import (
	"crypto/tls"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"io"
	"net"
	"net/http"
	net_url "net/url"
	"regexp"
	"time"
)

const (
	enableTimeout  = true
	disableTimeout = false
)

var rxHttpProto = regexp.MustCompile(`(?i)^http(s)?://`)

func Request(method, url string, body io.Reader, setTimeout bool) (response *http.Response, err error) {
	toStr, err := proc.GlobalConf.Get("http", "timeout", types.String)
	if err != nil {
		return
	}
	toDur, err := time.ParseDuration(toStr.(string) + "s")
	if err != nil {
		return
	}

	insecure, err := proc.GlobalConf.Get("http", "insecure", types.Boolean)
	if err != nil {
		return
	}

	client := &http.Client{}
	if setTimeout {
		tr := http.Transport{
			Dial:            dialTimeout(toDur, toDur),
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

	userAgent, err := proc.GlobalConf.Get("http", "user-agent", types.String)
	if err != nil {
		return
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	request.Header.Set("User-Agent", userAgent.(string))

	urlParsed, err := net_url.Parse(url)
	if err != nil {
		return
	}

	request.Header.Set("Host", urlParsed.Host)

	redirects, err := proc.GlobalConf.Get("http", "redirect", types.Boolean)
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
func dialTimeout(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func validateURL(url *string) {
	if !rxHttpProto.MatchString(*url) {
		v, err := proc.GlobalConf.Get("http", "default-https", types.Boolean)
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
