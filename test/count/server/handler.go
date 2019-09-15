package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
)

var (
	counter  int32
	response = []byte{'O', 'K'}
)

type testHTTPHandler struct{}

func (h testHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/count":
		count(w, r)
		return

	case "/t":
		total(w, false)

	case "/total":
		total(w, true)

	default:
		fmt.Fprintf(os.Stderr, "Invalid path: %s\n", r.RequestURI)
		os.Exit(1)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading count: %s\n", err)
	}

	i, err := strconv.Atoi(string(b))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading count: %s\n", err)
	}

	atomic.AddInt32(&counter, int32(i))

	_, err = w.Write(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing response: %s\n", err)
	}
}

func total(w http.ResponseWriter, human bool) {
	i := atomic.LoadInt32(&counter)
	s := strconv.Itoa(int(i))
	if human {
		s += " tests ran\n"
	}

	_, err := w.Write([]byte(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing response: %s\n", err)
	}

	exit <- true
}
