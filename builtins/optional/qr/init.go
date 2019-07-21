package qrimage

import (
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["qr"] = cmdQr
}

func cmdQr(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	p.Stdout.SetDataType("image")

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	// Create the barcode
	qrCode, err := qr.Encode(string(b), qr.M, qr.Auto)
	if err != nil {
		return err
	}

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// encode the barcode as png
	return png.Encode(p.Stdout, qrCode)
}
