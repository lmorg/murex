package string

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/lmorg/murex/lang"
	"image/png"
)

func init() {
	lang.GoFunctions["qr"] = cmdQr
}

func cmdQr(p *lang.Process) error {
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

	//// create the output file
	//file, _ := os.Create("qrcode.png")
	//defer file.Close()

	// encode the barcode as png
	return png.Encode(p.Stdout, qrCode)
}
