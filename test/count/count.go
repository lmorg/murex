package count

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

const env = "MUREX_TEST_COUNT"

// Tests a function to count all the unit tests that have been run
func Tests(t *testing.T, count int, funcName string) {
	switch strings.ToLower(os.Getenv(env)) {
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
		t.Errorf("Unable to log test count via HTTP (export %s=http): %s", env, err)
		return
	}

	req, err := http.Post("http://localhost:38000/count", "int", buf)
	if err != nil {
		t.Errorf("Unable to log test count via HTTP (export %s=http): %s", env, err)
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Potential mismatch logging test counts via HTTP (export %s=http): %s", env, err)
	}

	if string(b) != "OK" {
		t.Errorf("Potential mismatch logging test counts via HTTP (export %s=http): %s", env, `Body != "OK"`)
	}
}
