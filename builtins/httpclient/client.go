package httpclient

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

/*
func createClient(client *http.Client, request *http.Request) {
		client = new(http.Client)

		u, err := url.Parse(job.URL)
		//isErr(err)

		tr := http.Transport{
			Dial: dialTimeout(job),
		}

		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: job.Insecure}

		client.Transport = &tr
		client.Timeout = job.Timeout

		//request, err = http.NewRequest("GET", u.Scheme+"://"+ip+u.RequestURI(), nil)
		request, err = http.NewRequest(job.Method, job.URL, nil) // TODO: this will eventually support IPs with hostnames using the above code
		//isErr(err)

		request.Header.Set("User-Agent", job.UserAgent)
		request.Header.Set("Referer", job.Referrer)
		for header, _ := range job.Headers {
			request.Header.Set(header, job.Headers[header])
		}
		// for some reason 'request.Host' isn't setting the request header, so doing so manually with request.Header
		//request.Host = u.Host
		request.Header.Set("Host", u.Host)
		//job.AddCookies(request)

		return client, request
	}
}


func dialTimeout(job *Job) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		// http://stackoverflow.com/questions/16895294/how-to-set-timeout-for-http-get-requests-in-golang#16930649
		conn, err := net.DialTimeout(netw, addr, job.Timeout) //connect
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(job.Timeout)) //reply
		return conn, nil
	}
}*/
