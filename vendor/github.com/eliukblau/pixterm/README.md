```
   ___  _____  ____
  / _ \/  _/ |/_/ /____ ______ _    Made with love by Eliuk Blau
 / ___// /_>  </ __/ -_) __/  ' \   github.com/eliukblau/pixterm
/_/  /___/_/|_|\__/\__/_/ /_/_/_/   v1.2.4

```

# `PIXterm` - *draw images in your ANSI terminal with true color*

**`PIXterm`** ***shows images directly in your terminal***, recreating the pixels through a combination of [ANSI character background color](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors) and the [unicode lower half block element](http://en.wikipedia.org/wiki/Block_Elements). If image has transparency, an optional matte color can be used for background. Also, you can specify a dithering mode; in which case, the image is rendered using block elements with different shades, or using standard ASCII characters in the same way. In dithering mode, the matte color is used to fill the background of the blocks or characters.

The conversion process runs fast because it is parallelized in all CPUs.

Supported image formats: JPEG, PNG, GIF, BMP, TIFF, WebP.

Fetching images from HTTP/HTTPS is supported too.

#### Cool Screenshots

##### No Dithering (Classic Mode)

![Screenshot 1](screenshot01.png)

![Screenshot 2](screenshot02.png)

![Screenshot 3](screenshot03.png)

![Screenshot 4](screenshot04.png)

![Screenshot 5](screenshot05.png)

![Screenshot 6](screenshot06.png)

##### Dithering with Blocks

![Screenshot 7](screenshot07.png)

![Screenshot 8](screenshot08.png)

![Screenshot 9](screenshot09.png)

##### Dithering with Characters

![Screenshot 10](screenshot10.png)

![Screenshot 10](screenshot11.png)

![Screenshot 10](screenshot12.png)

#### Requirements
Your terminal emulator must be support *true color* feature in order to display image colors in a right way. In addition, you must use a monospaced font that includes the lower half block unicode character: `▄ (U+2584)`. I personally recommend [Envy Code R](http://damieng.com/blog/2008/05/26/envy-code-r-preview-7-coding-font-released). It's the nice font that shows in the screenshots. If you want to use the dithering mode with blocks, the font must also includes the following unicode characters: `█ (U+2588)`, `▓ (U+2593)`, `▒ (U+2592)`, `░ (U+2591)`. The dithering mode with characters works with standard ASCII chars.

#### Dependencies

All dependencies are directly included in the project via [Go's Vendor Directories](http://golang.org/cmd/go/#hdr-Vendor_Directories). You should not do anything else. Anyway, if you want to get the dependencies manually, project uses the [Glide Vendor Package Management](http://glide.sh). Follow its instructions.

###### Dependencies for `PIXterm` CLI tool

- Package [colorful](github.com/lucasb-eyer/go-colorful): `github.com/lucasb-eyer/go-colorful`
- Package [terminal](http://godoc.org/golang.org/x/crypto/ssh/terminal): `golang.org/x/crypto/ssh/terminal`

###### Dependencies for `ANSImage` Package

- Package [colorful](github.com/lucasb-eyer/go-colorful): `github.com/lucasb-eyer/go-colorful`
- Package [imaging](http://github.com/disintegration/imaging): `github.com/disintegration/imaging`
- Package [webp](http://godoc.org/golang.org/x/image/webp): `golang.org/x/image/webp`
- Package [bmp](http://godoc.org/golang.org/x/image/bmp): `golang.org/x/image/bmp`
- Package [tiff](http://godoc.org/golang.org/x/image/tiff): `golang.org/x/image/tiff`

#### Installation

*You need the [Go compiler](http://golang.org) version 1.7 or superior installed in your system.*

Run this command to automatically download sources and install **`PIXterm`** binary in your `$GOPATH/bin` directory:

`go get -u github.com/eliukblau/pixterm`

If you use Arch Linux, `eigengrau` has kindly created an AUR package for **`PIXterm`** (thanks man!). Run this command to install it:

`yaourt -S pixterm-git`

#### About

**`PIXterm`** is a terminal toy application that I made to exercise my skills on Go programming language. If you have not tried this language yet, please give it a try! It's easy, fast and very well organized. You'll not regret :D

*This application is originaly inspired by the clever [termpix](http://github.com/hopey-dishwasher/termpix), implemented in [Rust](http://www.rust-lang.org).*

*The dithering mode is my own port of the [Processing Textmode Engine](http://github.com/no-carrier/ProcessingTextmodeEngine)'s render.*

#### License

[Mozilla Public License Version 2.0](http://mozilla.org/MPL/2.0)

#### Contributors

- [@disq](http://github.com/disq) - Original code for image transparency support.
- [@timob](http://github.com/timob) - Fix for `ANSIpixel` type: use 8bit color component for output.
- [@HongjiangHuang](http://github.com/HongjiangHuang) - Original code for image download support.
