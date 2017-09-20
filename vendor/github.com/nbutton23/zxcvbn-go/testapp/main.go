package main

import (
	"fmt"
	"github.com/nbutton23/zxcvbn-go"
	"time"
)

func main() {
	//password :="Testaaatyhg890l33t"
	//fmt.Println(zxcvbn.PasswordStrength(password, nil))

	length := 5
	//pass := "68f9698fe2540c525fe35b15c6ae1a1788e079962b2ada3d1872c7665c95e148"
	pass := "NathanButtonTheAmazingAndAwesom12340987tyghjuikolpblkjhgfdsalabcdef"

	for length <= len(pass) {
		fmt.Printf("\nTested Password: %s\n", pass[0:length])
		startTime := time.Now().UTC()

		quality := zxcvbn.PasswordStrength(pass[0:length], []string{})

		fmt.Printf(
			`Password score    (0-4): %d
Estimated entropy (bit): %f
Estimated time to crack: %s%s`,
			quality.Score,
			quality.Entropy,
			quality.CrackTimeDisplay, "\n",
		)

		length += 1
		runtime := time.Now().UTC().Sub(startTime)
		fmt.Printf("Evaluation took: %s\n", runtime.String())
	}
}
