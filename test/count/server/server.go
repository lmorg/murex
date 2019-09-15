package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	client "github.com/lmorg/murex/test/count"
)

var (
	host = client.Host
	port = client.Port
	exit chan bool
)

func main() {
	exit = make(chan bool)

	fmt.Fprint(os.Stderr, "Starting count server....\n")
	fmt.Fprintf(os.Stderr, "\nSet the following to enable:\n    export %s=http\n", client.Env)
	fmt.Fprintf(os.Stderr, "\nTo get the totals:\n    curl %s:%d/total\n", host, port)

	go server()

	<-exit
	time.Sleep(2 * time.Second)
}

func server() {
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), testHTTPHandler{})
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
