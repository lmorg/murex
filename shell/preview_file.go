package shell

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/eliukblau/pixterm/pkg/ansimage"
	"github.com/lmorg/readline/v4"
)

type color struct{}

// RGBA is required for the Go color.Color interface.
func (col color) RGBA() (uint32, uint32, uint32, uint32) {
	return 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF
}

var rxImage = regexp.MustCompile(`\.(bmp|jpg|jpeg|png|gif|tiff|webp)$`)

func PreviewFile(ctx context.Context, _ []rune, filename string, incImages bool, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT) {
	p := make([]byte, 1*1024*1024)
	var i int

	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return //nil, 0, err
	}
	defer f.Close()

	var lines []string

	if incImages && rxImage.MatchString(filename) {
		img, err := ansimage.NewScaledFromReader(f, (2*size.Height)+5, size.Width, color{}, ansimage.ScaleModeFit, ansimage.NoDithering)
		if err != nil {
			return //nil, 0, err
		}
		lines = strings.Split(img.Render(), "\n")
		for i := range lines {
			count := strings.Count(lines[i], "\x1b") / 2
			if count < size.Width {
				lines[i] += strings.Repeat(" ", size.Width-count)
			}
		}
		callback(lines, 0, nil)
		return
	}

	i, err = f.Read(p)
	if err != nil {
		i = copy(p, []byte(err.Error()))
	}

	lines, _, err = previewParse(p[:i], size)
	if err != nil && strings.HasPrefix(err.Error(), binaryFile) {
		file := previewFile(filename)
		if len(file) == 0 {
			callback([]string{err.Error()}, 0, nil)
			return
		}
		callback(previewParse(file, size))
		return
	}

	callback(lines, 0, nil)
}
