package readline

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/eliukblau/pixterm/pkg/ansimage"
)

func getPreviewWidth(width int) (preview, forward int) {
	preview = width / 3
	forward = preview * 2
	preview += width - (preview + forward)
	forward -= 2
	return
}

type previewSizeT struct {
	Height  int
	Width   int
	Forward int
}

type previewCacheT struct {
	pos   int
	len   int
	lines []string
	size  *previewSizeT
}

func getPreviewXY() (*previewSizeT, error) {
	width, height, err := GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}

	if height == 0 {
		height = 25
	}

	if width == 0 {
		width = 80
	}

	preview, forward := getPreviewWidth(width)
	size := &previewSizeT{
		Height:  height / 3,
		Width:   preview,
		Forward: forward,
	}

	return size, nil
}

type color struct{}

// RGBA is required for the Go color.Color interface.
func (col color) RGBA() (uint32, uint32, uint32, uint32) {
	return 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF
}

var rxImage = regexp.MustCompile(`\.(bmp|jpg|jpeg|png|gif|tiff|webp)$`)

func previewCompile(filename string, incImages bool, size *previewSizeT) ([]string, error) {
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

	var (
		b     byte
		line  []byte
		lines []string
	)

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

parsePreview:
	for j := 0; j <= i; j++ {
		if j < i {
			b = p[j]
		} else {
			b = ' '
		}

		if b < ' ' && b != '\t' && b != '\r' && b != '\n' {
			file := previewFile(filename)
			if len(file) == 0 {
				return []string{"file contains binary data"}, nil
			}
			i = copy(p, file)
			line = []byte{}
			lines = []string{}
			goto parsePreview
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

const (
	//screenSave    = "\x1b[?47h"
	//screenRestore = "\x1b[?47l"
	curHome       = "\x1b[H"
	curPosSave    = "\x1b[s"
	curPosRestore = "\x1b[u"
)

func previewDraw(preview []string, size *previewSizeT) error {
	pf := fmt.Sprintf("│%%-%ds│\r\n", size.Width)

	_, _ = os.Stdout.WriteString(curPosSave + curHome)
	defer func() {
		_, _ = os.Stdout.WriteString(curPosRestore)
	}()

	moveCursorForwards(size.Forward)
	hr := strings.Repeat("─", size.Width)
	_, err := os.Stdout.WriteString("╭" + hr + "╮\r\n")
	if err != nil {
		return err
	}

	for i := 0; i <= size.Height; i++ {
		moveCursorForwards(size.Forward)

		if i >= len(preview) {
			blank := strings.Repeat(" ", size.Width)
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

	moveCursorForwards(size.Forward)
	_, err = os.Stdout.WriteString("╰" + hr + "╯\r\n")
	return err
}

func (rl *Instance) writePreview(item string) {
	if rl.ShowPreviews {
		size, err := getPreviewXY()
		if err != nil {
			rl.previewCache = nil
			return
		}

		item = strings.ReplaceAll(item, "\\", "")

		lines, err := previewCompile(item, rl.PreviewImages, size)
		if err != nil {
			rl.previewCache = nil
			return
		}
		err = previewDraw(lines, size)
		if err != nil {
			rl.previewCache = nil
			return
		}

		rl.previewCache = &previewCacheT{
			pos:   0,
			len:   size.Height,
			lines: lines,
			size:  size,
		}
	}
}

func (rl *Instance) previewPageUp() {
	if rl.previewCache == nil {
		return
	}

	rl.previewCache.pos -= rl.previewCache.len
	if rl.previewCache.pos < 0 {
		rl.previewCache.pos = 0
	}

	_ = previewDraw(rl.previewCache.lines[rl.previewCache.pos:], rl.previewCache.size)
}

func (rl *Instance) previewPageDown() {
	if rl.previewCache == nil {
		return
	}

	rl.previewCache.pos += rl.previewCache.len
	if rl.previewCache.pos > len(rl.previewCache.lines)-rl.previewCache.len-2 {
		rl.previewCache.pos = len(rl.previewCache.lines) - rl.previewCache.len - 2
		if rl.previewCache.pos < 0 {
			rl.previewCache.pos = 0
		}
	}

	_ = previewDraw(rl.previewCache.lines[rl.previewCache.pos:], rl.previewCache.size)
}
