package httpclient

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"testing"
)

/*
	This file creates a test HTTP server which just returns "$METHOD SUCCESSFUL"
	(where $METHOD will be GET / POST / etc).

	To instigate the server call `StartHTTPServer(t)` inside your testing func.
	StartHTTPServer returns the address it is listening on - which will typically
	be localhost and on a port number starting from 38001. Each instance of the
	server will increment the port number. So you can run multiple tests (eg with
	`-count n` flag) without worrying about shared server conflicts.
*/

var (
	testPort int32 = 38000
	testHost       = "localhost"
)

type testHTTPHandler struct{}

func (h testHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(r.Method + " SUCCESSFUL"))
	if err != nil {
		panic(err)
	}
}

func StartHTTPServer(t *testing.T) (addr string) {
	port := atomic.AddInt32(&testPort, 1)
	addr = fmt.Sprintf("%s:%d", testHost, port)
	go testHTTPServer(t, addr)
	return
}

func testHTTPServer(t *testing.T, addr string) {
	err := http.ListenAndServe(addr, testHTTPHandler{})
	if err != nil {
		t.Error(err)
	}
}
