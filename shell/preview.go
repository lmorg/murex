package shell

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/utils/readline"
	"github.com/mattn/go-runewidth"
)

const binaryFile = "file contains binary data"

func errBinaryFile(b byte) error {
	return fmt.Errorf("%s: %d", binaryFile, b)
}

func PreviewParseAppendEvent(previous []string, p []byte, size *readline.PreviewSizeT, title string) ([]string, error) {
	heading := append(
		previous,
		strings.Repeat("─", size.Width),
		fmt.Sprintf("Event `%s`:", title),
		strings.Repeat("╶", size.Width),
	)
	if len(previous) == 0 {
		heading = heading[1:]
	}

	lines, _, err := previewParse(p, size)

	return append(heading, lines...), err
}

func previewParse(p []byte, size *readline.PreviewSizeT) ([]string, int, error) {
	var (
		lines []string
		line  []rune
		width int
	)

	s := string(p)
	for _, r := range s {
		switch r {
		case 8:
			// handle backspace gracefully
			if len(line) > 0 {
				line = line[:len(line)-1]
			}
			continue

		case '\r':
			continue
		case '\n':
			if width == size.Width {
				continue
			}
			lines = append(lines, string(line))
			line = []rune{}
			width = 0

		case '\t':
			line = append(line, ' ', ' ', ' ', ' ')
			width += 4

		default:
			if r < ' ' && r != '\t' && r != '\r' && r != '\n' {
				return nil, 0, errBinaryFile(byte(r))
			}

			line = append(line, r)
		}

		width += runewidth.RuneWidth(r)
		if width >= size.Width {
			lines = append(lines, string(line))
			line = []rune{}
			width = 0
		}
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
