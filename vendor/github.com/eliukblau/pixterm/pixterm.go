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

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/eliukblau/pixterm/ansimage"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	pxtVersion = "1.2.4"
	pxtLogo    = `

   ___  _____  ____
  / _ \/  _/ |/_/ /____ ______ _    Made with love by Eliuk Blau
 / ___// /_>  </ __/ -_) __/  ' \   github.com/eliukblau/pixterm
/_/  /___/_/|_|\__/\__/_/ /_/_/_/   v{{VERSION}}

`
)

var (
	flagCredits bool
	flagDither  uint
	flagMatte   string
	flagScale   uint
	flagRows    uint
	flagCols    uint
)

func main() {
	validateFlags()
	runPixterm()
}

func printLogo() {
	fmt.Print(strings.Trim(strings.Replace(pxtLogo, "{{VERSION}}", pxtVersion, 1), "\n"), "\n\n")
}

func printCredits() {
	printLogo()
	fmt.Print("CONTRIBUTORS:\n\n")

	fmt.Print("  > @disq - http://github.com/disq\n")
	fmt.Print("      Original code for image transparency support.\n")
	fmt.Println()

	fmt.Print("  > @timob - http://github.com/timob\n")
	fmt.Print("      Fix for ANSIpixel type: use 8bit color component for output.\n")
	fmt.Println()

	fmt.Print("  > @HongjiangHuang - http://github.com/HongjiangHuang\n")
	fmt.Print("      Original code for image download support.\n")
	fmt.Println()

	fmt.Print("  > @danirod - http://github.com/danirod\n")
	fmt.Print("  > @Xpktro - http://github.com/Xpktro\n")
	fmt.Print("      Moral support.\n")
	fmt.Println()
}

func throwError(code int, v ...interface{}) {
	printLogo()
	log.New(os.Stderr, "[PIXTERM ERROR] ", log.LstdFlags).Println(v...)
	os.Exit(code)
}

func configureFlags() {
	flag.CommandLine.Usage = func() {
		printLogo()

		_, file := filepath.Split(os.Args[0])
		fmt.Print("USAGE:\n\n")
		fmt.Printf("  %s [options] image/url\n\n", file)

		fmt.Print("  Supported image formats: JPEG, PNG, GIF, BMP, TIFF, WebP.\n")
		fmt.Print("  Supported URL protocols: HTTP, HTTPS.\n\n")

		fmt.Print("OPTIONS:\n\n")
		flag.CommandLine.SetOutput(os.Stdout)
		flag.CommandLine.PrintDefaults()
		flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
		fmt.Print("  -help\n\tprints this message :D LOL\n")
		fmt.Println()
	}

	flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
	flag.CommandLine.Init(os.Args[0], flag.ExitOnError)

	flag.CommandLine.BoolVar(&flagCredits, "credits", false, "shows some love to contributors <3")
	flag.CommandLine.UintVar(&flagDither, "d", 0, "dithering `mode`:\n   0 - no dithering (default)\n   1 - with blocks\n   2 - with chars")
	flag.CommandLine.StringVar(&flagMatte, "m", "", "matte `color` for transparency or background\n(optional, hex format, default: 000000)")
	flag.CommandLine.UintVar(&flagScale, "s", 0, "scale `method`:\n   0 - resize (default)\n   1 - fill\n   2 - fit")
	flag.CommandLine.UintVar(&flagRows, "tr", 0, "terminal `rows` (optional, >=2; when piping, default: 24)")
	flag.CommandLine.UintVar(&flagCols, "tc", 0, "terminal `columns` (optional, >=2; when piping, default: 80)")

	flag.CommandLine.Parse(os.Args[1:])
}

func validateFlags() {
	if flagCredits {
		printCredits()
		os.Exit(0)
	}

	if flagDither != 0 && flagDither != 1 && flagDither != 2 {
		flag.CommandLine.Usage()
		os.Exit(2)
	}

	if flagScale != 0 && flagScale != 1 && flagScale != 2 {
		flag.CommandLine.Usage()
		os.Exit(2)
	}

	if (flagRows > 0 && flagRows < 2) || (flagCols > 0 && flagCols < 2) {
		flag.CommandLine.Usage()
		os.Exit(2)
	}

	// this is image filename
	if flag.CommandLine.Arg(0) == "" {
		flag.CommandLine.Usage()
		os.Exit(2)
	}
}

func isTerminal() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}

func getTerminalSize() (width, height int, err error) {
	if isTerminal() {
		return terminal.GetSize(int(os.Stdout.Fd()))
	}
	// fallback when piping to a file!
	return 80, 24, nil // VT100 terminal size
}

func runPixterm() {
	var (
		pix *ansimage.ANSImage
		err error
	)

	// get terminal size
	tx, ty, err := getTerminalSize()
	if err != nil {
		throwError(1, err)
	}

	// use custom terminal size (if applies)
	if ty--; flagRows != 0 { // no custom rows? subtract 1 for prompt spacing
		ty = int(flagRows) + 1 // weird, but in this case is necessary to add 1 :O
	}
	if flagCols != 0 {
		tx = int(flagCols)
	}

	// get scale mode from flag
	sm := ansimage.ScaleMode(flagScale)

	// get dithering mode from flag
	dm := ansimage.DitheringMode(flagDither)

	// set image scale factor for ANSIPixel grid
	sfy, sfx := ansimage.BlockSizeY, ansimage.BlockSizeX // 8x4 --> with dithering
	if ansimage.DitheringMode(flagDither) == ansimage.NoDithering {
		sfy, sfx = 2, 1 // 2x1 --> without dithering
	}

	// get matte color
	if flagMatte == "" {
		flagMatte = "000000" // black background
	}
	mc, err := colorful.Hex("#" + flagMatte) // RGB color from Hex format
	if err != nil {
		throwError(2, fmt.Sprintf("matte color : %s is not a hex-color", flagMatte))
	}

	// create new ANSImage from file
	file := flag.CommandLine.Arg(0)
	if matched, _ := regexp.MatchString(`^https?://`, file); matched {
		pix, err = ansimage.NewScaledFromURL(file, sfy*ty, sfx*tx, mc, sm, dm)
	} else {
		pix, err = ansimage.NewScaledFromFile(file, sfy*ty, sfx*tx, mc, sm, dm)
	}
	if err != nil {
		throwError(1, err)
	}

	// draw ANSImage to terminal
	if isTerminal() {
		ansimage.ClearTerminal()
	}
	pix.SetMaxProcs(runtime.NumCPU()) // maximum number of parallel goroutines!
	pix.Draw()
	if isTerminal() {
		fmt.Println()
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use paralelism for goroutines!
	configureFlags()
}
