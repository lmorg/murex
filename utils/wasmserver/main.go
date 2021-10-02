package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	projectPath = "github.com/lmorg/murex/"
	wasmPath    = "gen/website/wasm/"
)

func main() {
	path := pathBuilder()

	fmt.Println("Listening on :8080....")
	fmt.Printf("Serving: %s....\n", path)
	fmt.Printf("(press ^c to exit)")

	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(path)))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}

func goPath() string {
	return os.Getenv("GOPATH")
}

func pathBuilder() string {
	return goPath() + "/src/" + projectPath + wasmPath
}
