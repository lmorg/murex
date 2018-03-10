package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func main() {
	terminal.MakeRaw(int(os.Stdin.Fd()))
	for {
		b := make([]byte, 1024)
		i, err := os.Stdin.Read(b)
		if err != nil {
			panic(err)
		}

		fmt.Println(b[:i])
	}
}
