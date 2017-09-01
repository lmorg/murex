package preview

import (
	"github.com/lmorg/murex/builtins/preview/ansimage"
	"github.com/lmorg/murex/lang/types/data"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
)

type color struct{}

func init() {
	data.SetMime("image", "image/jpeg", "image/gif", "image/png", "image/bmp", "image/tiff", "image/webp")
	data.SetFileExtensions("image", "jpeg", "jpg", "gif", "png", "bmp", "tiff", "webp")
}

// Implement the Go color.Color interface.
func (col color) RGBA() (uint32, uint32, uint32, uint32) {
	return 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF
}

func pvImage(writer io.Writer, reader io.Reader) error {
	tx, ty, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	img, err := ansimage.NewScaledFromReader(2*(ty-1), tx, ansimage.ScaleModeFit, color{}, reader)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(img.Render()))
	return err
}
