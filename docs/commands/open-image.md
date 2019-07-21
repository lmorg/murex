# _murex_ Language Guide

## Command Reference: `open-image` (optional)

> Renders bitmap image data on your terminal

### Description

`open-image` is an optional builtin which will render images (JPEG, GIF,
PNG, BMP, TIFF and WebP) to the terminal using block characters and ANSI
colour sequences.

### Usage

    <stdin> -> open-image -> <stdout>
    
    open-image file-path -> <stdout>

### Examples

As a method

    » cat example.png -> open-image
    
As a function

    » open-image example.png

### Detail

`open-image` will fail if STDOUT is not a TTY.

### See Also

* [`qr` (optional)](../commands/qr.md):
  Creates a QR code from STDIN
* [open](../commands/open.md):
  