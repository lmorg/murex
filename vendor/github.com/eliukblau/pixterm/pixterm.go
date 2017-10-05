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
	"runtime"
	"strings"

	"github.com/eliukblau/pixterm/ansimage"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	pxtVersion = "1.1.1"
	pxtLogo    = `

   ___  _____  ____
  / _ \/  _/ |/_/ /____ ______ _    Made with love by Eliuk Blau
 / ___// /_>  </ __/ -_) __/  ' \   github.com/eliukblau/pixterm
/_/  /___/_/|_|\__/\__/_/ /_/_/_/   v{{VERSION}}

`
)

var (
	flagCredits bool
	flagMatte   string
	flagScale   uint
	flagRows    uint
	flagCols    uint
)

func main() {
	validateFlags()
	checkTerminal()
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
		fmt.Printf("  %s [options] image (JPEG, PNG, GIF, BMP, TIFF, WebP)\n\n", file)

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
	flag.CommandLine.StringVar(&flagMatte, "m", "", "matte `color` for image transparency\n\t(optional, hex format, default: 000000)")
	flag.CommandLine.UintVar(&flagScale, "s", 0, "scale `method`:\n\t  0 - resize (default)\n\t  1 - fill\n\t  2 - fit")
	flag.CommandLine.UintVar(&flagRows, "tr", 0, "terminal `rows` (optional, >=2)")
	flag.CommandLine.UintVar(&flagCols, "tc", 0, "terminal `columns` (optional, >=2)")

	flag.CommandLine.Parse(os.Args[1:])
}

func validateFlags() {
	if flagCredits {
		printCredits()
		os.Exit(0)
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

func checkTerminal() {
	if !terminal.IsTerminal(int(os.Stdout.Fd())) {
		throwError(1, "Not running on terminal :(")
	}
}

func getTerminalSize() (width, height int, err error) {
	return terminal.GetSize(int(os.Stdout.Fd()))
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
	if flagRows != 0 {
		ty = int(flagRows) + 1
	}
	if flagCols != 0 {
		tx = int(flagCols)
	}

	// get matte color
	if flagMatte == "" {
		flagMatte = "000000" // black background
	}
	mc, err := colorful.Hex("#" + flagMatte) // RGB color from Hex format
	if err != nil {
		throwError(2, fmt.Sprintf("matte color : %s is not a hex-color", flagMatte))
	}

	// set scale mode and create new ANSImage from file
	file := flag.CommandLine.Arg(0)
	switch flagScale {
	case 0:
		pix, err = ansimage.NewScaledFromFile(2*(ty-1), tx, ansimage.ScaleModeResize, mc, file)
	case 1:
		pix, err = ansimage.NewScaledFromFile(2*(ty-1), tx, ansimage.ScaleModeFill, mc, file)
	case 2:
		pix, err = ansimage.NewScaledFromFile(2*(ty-1), tx, ansimage.ScaleModeFit, mc, file)
	}
	if err != nil {
		throwError(1, err)
	}

	// draw ANSImage to terminal
	ansimage.ClearTerminal()
	pix.SetMaxProcs(runtime.NumCPU()) // maximum number of parallel goroutines!
	pix.Draw()
	fmt.Println()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use paralelism for goroutines!
	configureFlags()
}
