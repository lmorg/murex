package httpclient

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	neturl "net/url"
	"regexp"
	"time"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

const (
	enableTimeout  = true
	disableTimeout = false
)

var rxHttpProto = regexp.MustCompile(`(?i)^http(s)?://`)

// Request generates a HTTP request
func Request(method, url string, body io.Reader, conf *config.Config, setTimeout bool) (response *http.Response, err error) {
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

	//if contentType != "" {
	//	request.Header.Set("Content-Type", contentType)
	//}

	urlParsed, err := neturl.Parse(url)
	if err != nil {
		return
	}

	// code in meta functions
	setHeaders(request, conf, &urlParsed.Host)
	setCookies(request, conf, &urlParsed.Host)

	// Set important headers (these will override anything set in the 'headers' config)
	request.Header.Set("Host", urlParsed.Host)
	request.Header.Set("User-Agent", userAgent.(string))

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

// dialTimeout function unashamedly copy and pasted from:
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

func getMeta(conf *config.Config, meta string, domain *string) (metaValues, error) {
	s, err := conf.Get("http", meta, "json")
	if err != nil {
		return nil, err
	}

	var domains metaDomains
	err = json.UnmarshalMurex([]byte(s.(string)), &domains)
	if err != nil {
		return nil, err
	}

	if len(domains) > 0 && len(domains[*domain]) > 0 {
		return domains[*domain], nil
	}

	return nil, err
}

func setHeaders(request *http.Request, conf *config.Config, domain *string) error {
	headers, err := getMeta(conf, "headers", domain)
	if err != nil {
		return err
	}

	for name := range headers {
		request.Header.Add(name, headers[name])
	}

	return nil
}

func setCookies(request *http.Request, conf *config.Config, domain *string) error {
	cookies, err := getMeta(conf, "cookies", domain)
	if err != nil {
		return err
	}

	for name := range cookies {
		request.AddCookie(&http.Cookie{
			Name:  name,
			Value: cookies[name],
		})
	}

	return nil
}
