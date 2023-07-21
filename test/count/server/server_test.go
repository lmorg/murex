package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"testing"
	"time"

	client "github.com/lmorg/murex/test/count"
)

// There are two man purposes to this test:
//
// 1. Test the code actually compiles (since it's a separate project within the
//    murex project hierarchy)
// 2. Test the machine readable total API doesn't suffer a regression bug
func TestServer(t *testing.T) {
	client.Tests(t, 1)

	port--

	var err error
	go func() {
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), testHTTPHandler{})
	}()

	time.Sleep(500 * time.Millisecond)

	if err != nil {
		// lets not get caught up with testing if there is already a listener
		t.SkipNow()
	}

	client.Tests(t, 2)

	testCount(t)
	//testT(t) //this test doesn't yet work
	//os.Exit(0)
}

func testCount(t *testing.T) {
	buf := new(bytes.Buffer)
	_, err := buf.WriteString("1")
	if err != nil {
		t.Errorf("Unable to log test count via HTTP (export %s=http): %s", client.Env, err)
		return
	}

	req, err := http.Post(fmt.Sprintf("http://%s:%d/count", host, port), "int", buf)
	if err != nil {
		t.Errorf("Unable to log test count via HTTP (export %s=http): %s", client.Env, err)
		return
	}

	b, err := io.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Potential mismatch logging test counts via HTTP (export %s=http): %s", client.Env, err)
	}

	if string(b) != "OK" {
		t.Errorf("Potential mismatch logging test counts via HTTP (export %s=http): %s", client.Env, `Body != "OK"`)
	}
}

func testT(t *testing.T) {
	req, err := http.Get(fmt.Sprintf("http://%s:%d/t", host, port))
	if err != nil {
		t.Errorf("Unable to log test count via HTTP (export %s=http): %s", client.Env, err)
		return
	}

	b, err := io.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Potential mismatch logging test counts via HTTP (export %s=http): %s", client.Env, err)
	}

	regex := `^[0-9]+$`
	rx := regexp.MustCompile(regex)
	if !rx.Match(b) {
		t.Errorf("Potential mismatch logging test counts via HTTP (export %s=http): Body !~ %s", client.Env, regex)
	}
}
