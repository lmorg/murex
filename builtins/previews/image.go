package previews

import (
	"github.com/eliukblau/pixterm/ansimage"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

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

	mc, err := colorful.Hex("#000000")

	img, err := ansimage.NewScaledFromFile(2*(ty-1), tx, ansimage.ScaleModeFit, mc, filename)
	if err != nil {
		return err
	}

	s := img.Render()

	_, err = p.Stdout.Write([]byte(s))
	return nil
}
