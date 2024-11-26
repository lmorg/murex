package httpclient

import "net/http"

type metaDomains map[string]metaValues
type metaValues map[string]string

type httpStatus struct {
	Code    int
	Message string
}

type jsonHttp struct {
	Status  httpStatus
	Headers http.Header
	Body    any
}
