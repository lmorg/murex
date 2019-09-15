package count

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	// Env is the name for the exported envvar we should check to see which count logging method to use
	Env = "MUREX_TEST_COUNT"

	// Host is the host name for the HTTP count listener
	Host = "localhost"

	// Port is the port number which the HTTP count is listening on
	Port = 38000
)

// Tests a function to count all the unit tests that have been run
func Tests(t *testing.T, count int, funcName string) {
	switch strings.ToLower(os.Getenv(Env)) {
	case "log":
		t.Logf("%s tests ran: %d", funcName, count)

	case "http":
		httpReq(t, count)

	default:
	}
}

func httpReq(t *testing.T, count int) {
	s := strconv.Itoa(count)
	buf := new(bytes.Buffer)
	_, err := buf.WriteString(s)
	if err != nil {
		t.Errorf("Unable to log test count via HTTP (export %s=http): %s", Env, err)
		return
	}

	req, err := http.Post(fmt.Sprintf("http://%s:%d/count", Host, Port), "int", buf)
	if err != nil {
		t.Errorf("Unable to log test count via HTTP (export %s=http): %s", Env, err)
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Potential mismatch logging test counts via HTTP (export %s=http): %s", Env, err)
	}

	if string(b) != "OK" {
		t.Errorf("Potential mismatch logging test counts via HTTP (export %s=http): %s", Env, `Body != "OK"`)
	}
}
