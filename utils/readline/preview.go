package readline

import (
	"fmt"
	"os"
	"strings"
)

func getPreviewXY() (int, int, error) {
	width, height, err := GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0, err
	}

	if height == 0 {
		height = 25
	}

	if width == 0 {
		width = 80
	}

	maxWidth := width / 3
	maxHeight := height / 3

	return maxWidth, maxHeight, nil
}

func previewCompile(filename string, width, height int) ([]string, error) {
	p := make([]byte, 8*1024)
	var i int

	if strings.HasSuffix(filename, " ") {
		filename = strings.TrimSpace(filename)
	}
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	i, err = f.Read(p)
	if err != nil {
		i = copy(p, []byte(err.Error()))
	}

	var (
		b     byte
		line  []byte
		lines []string
	)

	for j := 0; j <= i; j++ {
		if j < i {
			b = p[j]
		} else {
			b = ' '
		}

		if b < ' ' && b != '\t' && b != '\r' && b != '\n' {
			return []string{fmt.Sprintf("file contains binary data: %d, %d", b, i)}, nil
		}

		switch b {
		case '\r':
			continue
		case '\n':
			/*if len(lines) == 0 {
				continue
			}*/
			lines = append(lines, string(line))
			line = []byte{}
		case '\t':
			line = append(line, ' ', ' ', ' ', ' ')
		default:
			line = append(line, b)
		}

		if len(line) >= width {
			lines = append(lines, string(line))
			line = []byte{}
		}

		if len(lines) >= height {
			break
		}
	}

	if len(line) > 0 {
		lines = append(lines, string(line))
	}

	return lines, nil
}

const (
	//screenSave    = "\x1b[?47h"
	//screenRestore = "\x1b[?47l"
	curHome       = "\x1b[H"
	curPosSave    = "\x1b[s"
	curPosRestore = "\x1b[u"
)

func previewDraw(preview []string, width, height int) error {
	pf := fmt.Sprintf("│%%-%ds│\r\n", width)

	_, _ = os.Stdout.WriteString(curPosSave + curHome)
	defer func() {
		_, _ = os.Stdout.WriteString(curPosRestore)
	}()

	fwd := width*2 - 2
	moveCursorForwards(fwd)
	hr := strings.Repeat("─", width)
	_, err := os.Stdout.WriteString("╭" + hr + "╮\r\n")
	if err != nil {
		return err
	}

	for i := 0; i <= height; i++ {
		moveCursorForwards(fwd)

		if i >= len(preview) {
			blank := strings.Repeat(" ", width)
			_, err = os.Stdout.WriteString("│" + blank + "│\r\n")
			if err != nil {
				return err
			}
			continue
		}

		out := fmt.Sprintf(pf, preview[i])
		_, err = os.Stdout.WriteString(out)
		if err != nil {
			return err
		}
	}

	moveCursorForwards(fwd)
	_, err = os.Stdout.WriteString("╰" + hr + "╯\r\n")
	return err
}
