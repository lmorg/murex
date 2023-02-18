package preview

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/eliukblau/pixterm/pkg/ansimage"
	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/readline"
)

type color struct{}

// RGBA is required for the Go color.Color interface.
func (col color) RGBA() (uint32, uint32, uint32, uint32) {
	return 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF
}

var rxImage = regexp.MustCompile(`\.(bmp|jpg|jpeg|png|gif|tiff|webp)$`)

func File(filename string, incImages bool, size *readline.PreviewSizeT) ([]string, error) {
	p := make([]byte, 1*1024*1024)
	var i int

	if strings.HasSuffix(filename, " ") {
		filename = strings.TrimSpace(filename)
	}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string

	if incImages && rxImage.MatchString(filename) {
		img, err := ansimage.NewScaledFromReader(f, 2*size.Height-1, size.Width, color{}, ansimage.ScaleModeFit, ansimage.NoDithering)
		if err != nil {
			return nil, err
		}
		lines = strings.Split(img.Render(), "\n")
		for i := range lines {
			count := strings.Count(lines[i], "\x1b") / 2
			if count < size.Width {
				lines[i] += strings.Repeat(" ", size.Width-count)
			}
		}
		return lines, nil
	}

	i, err = f.Read(p)
	if err != nil {
		i = copy(p, []byte(err.Error()))
	}

	lines, err = parse(p[:i], size)
	if err != nil && err.Error() == errBinaryFile.Error() {
		file := previewFile(filename)
		if len(file) == 0 {
			return []string{err.Error()}, nil
		}
		return parse(file, size)
	}

	return lines, err
}

var errBinaryFile = errors.New("file contains binary data")

func parse(p []byte, size *readline.PreviewSizeT) ([]string, error) {
	var (
		lines []string
		line  []byte
		b     byte
		i     = len(p)
	)

	for j := 0; j <= i; j++ {
		if j+1 < i && p[j+1] == 8 { // backspace
			j += 2
			continue
		}

		if j < i {
			b = p[j]
		} else {
			b = ' '
		}

		if b == 8 { // backspace
			continue
		}

		if b < ' ' && b != '\t' && b != '\r' && b != '\n' {
			panic(fmt.Sprint(line, json.LazyLoggingPretty(lines)))
			return nil, errBinaryFile
		}

		switch b {
		case '\r':
			continue
		case '\n':
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
	}

	if len(line) > 0 {
		lines = append(lines, string(line))
	}

	return lines, nil
}

func Command(command string, _ bool, size *readline.PreviewSizeT) ([]string, error) {
	if strings.HasSuffix(command, " ") {
		command = strings.TrimSpace(command)
	}

	if lang.GlobalAliases.Exists(command) {
		return nil, nil
	}

	if lang.MxFunctions.Exists(command) {
		r, err := lang.MxFunctions.Block(command)
		if err != nil {
			return nil, err
		}
		return parse([]byte(string(r)), size)
	}

	if lang.GoFunctions[command] != nil {
		return parse([]byte(docs.Definition[command]), size)
	}

	if (*autocomplete.GlobalExes.Get())[command] {
		return parse(manPage(command, size), size)
	}

	return nil, nil
}
