package readline

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func getTermWidth() (termWidth int) {
	var err error
	fd := int(os.Stdout.Fd())
	termWidth, _, err = GetSize(fd)
	if err != nil {
		termWidth = 100
	}

	return
}

func printf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	print(s)
}

func print(s string) {
	os.Stdout.WriteString(s)
}

func rLen(r []rune) (length int) {
	for _, i := range r {
		length += utf8.RuneLen(i)
	}
	return
}
