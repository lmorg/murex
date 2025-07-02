package openimage

import (
	"errors"
	"io"
	"os"

	"github.com/eliukblau/pixterm/pkg/ansimage"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/readline/v4"
)

func init() {
	lang.SetMime("image", "image/jpeg", "image/gif", "image/png", "image/bmp", "image/tiff", "image/webp")
	lang.SetFileExtensions("image", "jpeg", "jpg", "gif", "png", "bmp", "tiff", "webp")
	lang.DefineMethod("open-image", cmdOpenImage, "image", types.Null)
}

// color implements the Go color.Color interface.
type color struct{}

// RGBA is required for the Go color.Color interface.
func (col color) RGBA() (uint32, uint32, uint32, uint32) {
	return 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF
}

func cmdOpenImage(p *lang.Process) error {
	var reader io.Reader

	switch {
	case !p.Stdout.IsTTY():
		return errors.New("this function is expecting to output to the terminal")

	case p.IsMethod:
		reader = p.Stdin

	default:
		name, err := p.Parameters.String(0)
		if err != nil {
			return err
		}
		reader, err = os.Open(name)
		if err != nil {
			return err
		}
	}

	tx, ty, err := readline.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	img, err := ansimage.NewScaledFromReader(reader, 2*(ty-1), tx, color{}, ansimage.ScaleModeFit, ansimage.NoDithering)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(img.Render()))
	return err
}
