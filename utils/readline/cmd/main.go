package main

import (
	"fmt"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/readline"
	"strings"
)

func main() {
	readline.SyntaxHighlight = shell.Highlight
	readline.TabCompleter = Tab

	for {
		s, err := readline.Readline()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Readline: '" + s + "'")
	}
}

func Tab(line []rune, pos int) (suggestions []string) {
	items := []string{
		"qwerty",
		"12345",
		"qwertyuiop",
		"qwa1234",
		"aa",
		"abaya",
		"abomasum",
		"absquatulate",
		"adscititious",
		"afreet",
		"Albertopolis",
		"alcazar",
		"amphibology",
		"amphisbaena",
		"anfractuous",
		"anguilliform",
		"apoptosis",
		"apple-knocker",
		"argle-bargle",
		"Argus-eyed",
		"argute",
		"ariel",
		"aristotle",
		"aspergillum",
		"astrobleme",
		"Attic",
		"autotomy",
		"badmash",
		"bandoline",
		"bardolatry",
		"Barmecide",
		"barn",
		"bashment",
		"bawbee",
		"benthos",
		"bergschrund",
		"bezoar",
		"bibliopole",
		"bichon",
		"bilboes",
		"bindlestiff",
		"bingle",
		"blatherskite",
		"bleeding",
		"blind",
		"bobsy-die",
		"boffola",
		"boilover",
		"borborygmus",
		"breatharian",
		"Brobdingnagian",
		"bruxism",
		"bumbo",
	}

	for i := range items {
		if strings.HasPrefix(items[i], string(line)) {
			suggestions = append(suggestions, items[i][pos:])
		}
	}

	return
}
