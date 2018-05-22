package open

import (
	"io"
	"os"

	"github.com/eliukblau/pixterm/ansimage"
	"github.com/lmorg/murex/lang/types/define"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	define.SetMime("image", "image/jpeg", "image/gif", "image/png", "image/bmp", "image/tiff", "image/webp")
	define.SetFileExtensions("image", "jpeg", "jpg", "gif", "png", "bmp", "tiff", "webp")
}

// color implements the Go color.Color interface.
type color struct{}

// RGBA is required for the Go color.Color interface.
func (col color) RGBA() (uint32, uint32, uint32, uint32) {
	return 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF
}

func pvImage(writer io.Writer, reader io.Reader) error {
	/*b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	s := base64.StdEncoding.EncodeToString(b)

	if _, err := writer.Write([]byte{27}); err != nil {
		return err
	}

	if _, err := writer.Write([]byte("_Gf=100,t=d,i=1;")); err != nil {
		return err
	}

	if _, err := writer.Write([]byte(s)); err != nil {
		return err
	}

	if _, err := writer.Write([]byte{27, '\\'}); err != nil {
		return err
	}

	return nil*/

	tx, ty, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	img, err := ansimage.NewScaledFromReader(reader, 2*(ty-1), tx, color{}, ansimage.ScaleModeFit, ansimage.NoDithering)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(img.Render()))
	return err
}
