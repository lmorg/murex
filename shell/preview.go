package shell

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/utils/readline"
)

const binaryFile = "file contains binary data"

func errBinaryFile(b byte) error {
	return fmt.Errorf("%s: %d", binaryFile, b)
}

func previewParse(p []byte, size *readline.PreviewSizeT) ([]string, int, error) {
	var (
		lines []string
		line  []byte
		b     byte
		i     = len(p)
		last  byte
	)

	for j := 0; j <= i; j++ {
		if j < i {
			b = p[j]
		} else {
			b = ' '
		}

		if b == 8 {
			// handle backspace gracefully
			if len(line) > 0 {
				line = line[:len(line)-1]
			}
			continue
		}

		if b < ' ' && b != '\t' && b != '\r' && b != '\n' {
			return nil, 0, errBinaryFile(p[j])
		}

		switch b {
		case '\r':
			last = b
			continue
		case '\n':
			if (len(line) == 0 && len(lines) > 0 && len(lines[len(lines)-1]) == size.Width) ||
				last == '\r' {
				last = b
				continue
			}
			lines = append(lines, string(line))
			line = []byte{}
		case '\t':
			line = append(line, ' ', ' ', ' ', ' ')
		default:
			line = append(line, b)
		}

		if len(line) >= size.Width {
			lines = append(lines, string(line))
			line = []byte{}
		}
		last = b
	}

	if len(line) > 0 {
		lines = append(lines, string(line))
	}

	return lines, 0, nil
}

func previewPos(lines []string, item string) int {
	for i := range lines {
		switch {
		case strings.HasPrefix(item, "--"):
			switch {
			case strings.Contains(lines[i], ", "+item):
				// comma separated
				return i
			case strings.Contains(lines[i], "  "+item):
				// whitespace separated
				return i
			default:
				continue
			}
		default:
			if strings.Contains(lines[i], "  "+item) {
				return i
			}
		}
	}

	return 0
}
