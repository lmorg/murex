//       ___  _____  ____
//      / _ \/  _/ |/_/ /____ ______ _
//     / ___// /_>  </ __/ -_) __/  ' \
//    /_/  /___/_/|_|\__/\__/_/ /_/_/_/
//
//    Copyright 2017 Eliuk Blau
//
//    This Source Code Form is subject to the terms of the Mozilla Public
//    License, v. 2.0. If a copy of the MPL was not distributed with this
//    file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ansimage

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"  // initialize decoder
	_ "image/jpeg" // initialize decoder
	_ "image/png"  // initialize decoder
	"io"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	_ "golang.org/x/image/bmp"  // initialize decoder
	_ "golang.org/x/image/tiff" // initialize decoder
	_ "golang.org/x/image/webp" // initialize decoder
)

// Unicode Block Element character used to represent lower pixel in terminal row.
// INFO: http://en.wikipedia.org/wiki/Block_Elements
const lowerHalfBlock = "\u2584"

// ANSImage scale modes:
// resize (full scaled to area),
// fill (resize and crop the image with a center anchor point to fill area),
// fit (resize the image to fit area, preserving the aspect ratio).
const (
	ScaleModeResize = scaleMode(iota)
	ScaleModeFill
	ScaleModeFit
)

var (
	// ErrOddHeight happens when ANSImage height is not even value.
	ErrOddHeight = errors.New("ANSImage: height must be even value")

	// ErrInvalidBounds happens when ANSImage height or width are invalid values.
	ErrInvalidBounds = errors.New("ANSImage: height or width must be >=2")

	// ErrOutOfBounds happens when ANSI-pixel coordinates are out of ANSImage bounds.
	ErrOutOfBounds = errors.New("ANSImage: out of bounds")
)

// scaleMode type is used for image scale mode constants.
type scaleMode uint8

// ANSIpixel represents a pixel of an ANSImage.
type ANSIpixel struct {
	R, G, B uint8
	upper   bool
}

// ANSImage represents an image encoded in ANSI escape codes.
type ANSImage struct {
	h, w     int
	maxprocs int
	pixmap   [][]*ANSIpixel
}

// Render returns the ANSI-compatible string form of ANSI-pixel.
func (ap *ANSIpixel) Render() string {
	if ap.upper {
		return fmt.Sprintf(
			"\033[48;2;%d;%d;%dm",
			ap.R, ap.G, ap.B,
		)
	}
	return fmt.Sprintf(
		"\033[38;2;%d;%d;%dm%s",
		ap.R, ap.G, ap.B,
		lowerHalfBlock,
	)
}

// Height gets total rows of ANSImage.
func (ai *ANSImage) Height() int {
	return ai.h
}

// Width gets total columns of ANSImage.
func (ai *ANSImage) Width() int {
	return ai.w
}

// SetMaxProcs sets the maximum number of parallel goroutines to render the ANSImage
// (user should manually sets `runtime.GOMAXPROCS(max)` before to this change takes effect).
func (ai *ANSImage) SetMaxProcs(max int) {
	ai.maxprocs = max
}

// GetMaxProcs gets the maximum number of parallels goroutines to render the ANSImage.
func (ai *ANSImage) GetMaxProcs() int {
	return ai.maxprocs
}

// SetAt sets ANSI-pixel color (RBG) in coordinates (y,x).
func (ai *ANSImage) SetAt(y, x int, r, g, b uint8) error {
	if y >= 0 && y < ai.h && x >= 0 && x < ai.w {
		ai.pixmap[y][x].R = r
		ai.pixmap[y][x].G = g
		ai.pixmap[y][x].B = b
		ai.pixmap[y][x].upper = (y%2 == 0)
		return nil
	}
	return ErrOutOfBounds
}

// GetAt gets ANSI-pixel in coordinates (y,x).
func (ai *ANSImage) GetAt(y, x int) (*ANSIpixel, error) {
	if y >= 0 && y < ai.h && x >= 0 && x < ai.w {
		return &ANSIpixel{
				R:     ai.pixmap[y][x].R,
				G:     ai.pixmap[y][x].G,
				B:     ai.pixmap[y][x].B,
				upper: ai.pixmap[y][x].upper,
			},
			nil
	}
	return nil, ErrOutOfBounds
}

// Render returns the ANSI-compatible string form of ANSImage.
// (Nice info for ANSI True Colour - https://gist.github.com/XVilka/8346728)
func (ai *ANSImage) Render() string {
	type renderData struct {
		row    int
		render string
	}

	rows := make([]string, ai.h/2)

	for y := 0; y < ai.h; y += ai.maxprocs {
		ch := make(chan renderData, ai.maxprocs)

		for n, r := 1, y+1; (n <= ai.maxprocs) && (2*r+1 < ai.h); n, r = n+1, y+n+1 {
			go func(r, y int) {
				var str string
				for x := 0; x < ai.w; x++ {
					str += ai.pixmap[y][x].Render()   // upper pixel
					str += ai.pixmap[y+1][x].Render() // lower pixel
				}
				str += "\033[0m\n" // reset ansi style
				ch <- renderData{row: r, render: str}
			}(r, 2*r)

			// DEBUG:
			// fmt.Printf("y:%d | n:%d | r:%d | 2*r:%d\n", y, n, r, 2*r)
			// time.Sleep(time.Millisecond * 100)
		}

		for n, r := 1, y+1; (n <= ai.maxprocs) && (2*r+1 < ai.h); n, r = n+1, y+n+1 {
			data := <-ch
			rows[data.row] = data.render

			// DEBUG:
			// fmt.Printf("data.row:%d\n", data.row)
			// time.Sleep(time.Millisecond * 100)
		}
	}

	return strings.Join(rows, "")
}

// RenderOLD returns the ANSI-compatible string form of ANSImage
// func (ai *ANSImage) RenderOLD() string {
// 	var str string
// 	for y := 0; y < ai.h; y += 2 {
// 		for x := 0; x < ai.w; x++ {
// 			str += ai.pixmap[y][x].Render()   // upper pixel
// 			str += ai.pixmap[y+1][x].Render() // lower pixel
// 			//fmt.Printf("%d:%d\n", x, y)
// 		}
// 		str += "\x1b[m\n" // reset ansi style
// 	}
// 	return str
// }

// Draw writes the ANSImage to standard output (terminal).
func (ai *ANSImage) Draw() {
	fmt.Print(ai.Render())
}

// New creates a new empty ANSImage ready to draw on it.
func New(h, w int) (*ANSImage, error) {
	if h%2 != 0 {
		return nil, ErrOddHeight
	}
	if h < 2 || w < 2 {
		return nil, ErrInvalidBounds
	}

	ansimage := &ANSImage{
		h: h, w: w,
		maxprocs: 1,
		pixmap: func() [][]*ANSIpixel {
			v := make([][]*ANSIpixel, h)
			for y := 0; y < h; y++ {
				v[y] = make([]*ANSIpixel, w)
				for x := 0; x < w; x++ {
					v[y][x] = &ANSIpixel{
						R: 0, G: 0, B: 0,
						upper: (y%2 == 0),
					}
				}
			}
			return v
		}(),
	}

	return ansimage, nil
}

// NewFromReader creates a new ANSImage from an io.Reader.
// Background color is used to fill when image has transparency.
func NewFromReader(bg color.Color, reader io.Reader) (*ANSImage, error) {
	image, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return createANSImage(bg, image)
}

// NewScaledFromReader creates a new scaled ANSImage from an io.Reader.
// Background color is used to fill when image has transparency.
func NewScaledFromReader(y, x int, sm scaleMode, bg color.Color, reader io.Reader) (*ANSImage, error) {
	image, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	switch sm {
	case ScaleModeResize:
		image = imaging.Resize(image, x, y, imaging.Lanczos)
	case ScaleModeFill:
		image = imaging.Fill(image, x, y, imaging.Center, imaging.Lanczos)
	case ScaleModeFit:
		image = imaging.Fit(image, x, y, imaging.Lanczos)
	}

	return createANSImage(bg, image)
}

// NewFromFile creates a new ANSImage from a file.
// Background color is used to fill when image has transparency.
func NewFromFile(bg color.Color, name string) (*ANSImage, error) {
	reader, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return NewFromReader(bg, reader)
}

// NewScaledFromFile creates a new scaled ANSImage from a file.
// Background color is used to fill when image has transparency.
func NewScaledFromFile(y, x int, sm scaleMode, bg color.Color, name string) (*ANSImage, error) {
	reader, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return NewScaledFromReader(y, x, sm, bg, reader)
}

// ClearTerminal clears current terminal buffer using ANSI escape code.
// (Nice info for ANSI escape codes - http://unix.stackexchange.com/questions/124762/how-does-clear-command-work)
func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

// createANSImage loads data from an image and returns an ANSImage.
// Background color is used to fill when image has transparency.
func createANSImage(bg color.Color, img image.Image) (*ANSImage, error) {
	var rgbaOut *image.RGBA
	bounds := img.Bounds()

	// do compositing only if background color has no transparency (thank you @disq for the idea!)
	// (info - http://stackoverflow.com/questions/36595687/transparent-pixel-color-go-lang-image)
	if _, _, _, a := bg.RGBA(); a >= 0xffff {
		rgbaOut = image.NewRGBA(bounds)
		draw.Draw(rgbaOut, bounds, image.NewUniform(bg), image.ZP, draw.Src)
		draw.Draw(rgbaOut, bounds, img, image.ZP, draw.Over)
	} else {
		if v, ok := img.(*image.RGBA); ok {
			rgbaOut = v
		} else {
			rgbaOut = image.NewRGBA(bounds)
			draw.Draw(rgbaOut, bounds, img, image.ZP, draw.Src)
		}
	}

	yMin, xMin := bounds.Min.Y, bounds.Min.X
	yMax, xMax := bounds.Max.Y, bounds.Max.X

	if yMax%2 != 0 {
		yMax-- // always even value!
	}

	ansimage, err := New(yMax, xMax)
	if err != nil {
		return nil, err
	}

	for y := yMin; y < yMax; y++ {
		for x := xMin; x < xMax; x++ {
			v := rgbaOut.RGBAAt(x, y)
			if err := ansimage.SetAt(y, x, v.R, v.G, v.B); err != nil {
				return nil, err
			}
		}
	}

	return ansimage, nil
}
