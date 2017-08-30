package previews

import (
	"github.com/eliukblau/pixterm/ansimage"
	"github.com/lmorg/murex/lang/proc"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type color struct{}

// Implement the Go color.Color interface.
func (col color) RGBA() (uint32, uint32, uint32, uint32) {
	return 0, 0, 0, 0xFFFF
}

func init() {
	proc.GoFunctions["img"] = cmdImg
}

func cmdImg(p *proc.Process) error {
	filename, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	tx, ty, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	img, err := ansimage.NewScaledFromFile(2*(ty-1), tx, ansimage.ScaleModeFit, color{}, filename)

	if err != nil {
		return err
	}

	s := img.Render()

	_, err = p.Stdout.Write([]byte(s))
	return nil
}
