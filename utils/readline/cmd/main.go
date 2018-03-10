package main

import (
	"fmt"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/readline"
)

func main() {
	readline.Prompt = ">> "
	readline.SyntaxHighlight = shell.Highlight

	for {
		s, err := readline.Readline()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Readline: '" + s + "'")
	}
}
