package preview

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/eliukblau/pixterm/pkg/ansimage"
	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/murex/utils/readline"
)

type color struct{}

// RGBA is required for the Go color.Color interface.
func (col color) RGBA() (uint32, uint32, uint32, uint32) {
	return 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF
}

var rxImage = regexp.MustCompile(`\.(bmp|jpg|jpeg|png|gif|tiff|webp)$`)

func File(_ []rune, filename string, incImages bool, size *readline.PreviewSizeT) ([]string, int, error) {
	p := make([]byte, 1*1024*1024)
	var i int

	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	var lines []string

	if incImages && rxImage.MatchString(filename) {
		img, err := ansimage.NewScaledFromReader(f, (2*size.Height)+5, size.Width, color{}, ansimage.ScaleModeFit, ansimage.NoDithering)
		if err != nil {
			return nil, 0, err
		}
		lines = strings.Split(img.Render(), "\n")
		for i := range lines {
			count := strings.Count(lines[i], "\x1b") / 2
			if count < size.Width {
				lines[i] += strings.Repeat(" ", size.Width-count)
			}
		}
		return lines, 0, nil
	}

	i, err = f.Read(p)
	if err != nil {
		i = copy(p, []byte(err.Error()))
	}

	lines, _, err = parse(p[:i], size)
	if err != nil && err.Error() == errBinaryFile.Error() {
		file := previewFile(filename)
		if len(file) == 0 {
			return []string{err.Error()}, 0, nil
		}
		return parse(file, size)
	}

	return lines, 0, err
}

var errBinaryFile = errors.New("file contains binary data")

func parse(p []byte, size *readline.PreviewSizeT) ([]string, int, error) {
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

		if b < ' ' && b != '\t' && b != '\r' && b != '\n' {
			return nil, 0, errBinaryFile
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

func Command(_ []rune, command string, _ bool, size *readline.PreviewSizeT) ([]string, int, error) {
	if lang.GlobalAliases.Exists(command) {
		alias := lang.GlobalAliases.Get(command)
		if len(alias) == 0 {
			return nil, 0, nil
		}
		return Command(nil, alias[0], false, size)
	}

	if lang.MxFunctions.Exists(command) {
		r, err := lang.MxFunctions.Block(command)
		if err != nil {
			return nil, 0, err
		}
		return parse([]byte(string(r)), size)
	}

	syn := docs.Synonym[command]
	if syn != "" {
		return parse([]byte(docs.Definition[syn]), size)
	}

	if (*autocomplete.GlobalExes.Get())[command] {
		return parse(manPage(command, size), size)
	}

	return nil, 0, nil
}

func Parameter(block []rune, parameter string, incImages bool, size *readline.PreviewSizeT) ([]string, int, error) {
	if utils.Exists(parameter) {
		return File(nil, parameter, incImages, size)
	}

	pt, _ := parser.Parse(block, 0)
	lines, _, err := Command(nil, pt.FuncName, false, size)
	if err != nil {
		return lines, 0, err
	}
	for i := range lines {
		if strings.Contains(lines[i], "  "+parameter) {
			return lines, i, nil
		}
	}

	return lines, 0, nil
}
