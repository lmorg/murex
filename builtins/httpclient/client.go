package httpclient

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"net"
	"net/http"
	net_url "net/url"
	"time"
)

func get(url string) (response *http.Response, err error) {
	toStr, err := proc.GlobalConf.Get("http", "Timeout", types.String)
	if err != nil {
		return
	}

	toDur, err := time.ParseDuration(toStr.(string) + "s")
	if err != nil {
		return
	}

	ua, err := proc.GlobalConf.Get("http", "User-Agent", types.String)
	if err != nil {
		return
	}

	tr := http.Transport{
		Dial: dialTimeout(toDur, toDur),
	}
	client := &http.Client{
		Timeout:   toDur,
		Transport: &tr,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	request.Header.Set("User-Agent", ua.(string))

	urlParsed, err := net_url.Parse(url)
	if err != nil {
		return
	}

	request.Header.Set("Host", urlParsed.Host)

	return client.Do(request)
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
