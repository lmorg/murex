# _murex_ Shell Docs

## Command Reference: `open-image` 

> Renders bitmap image data on your terminal

## Description

`open-image` is an optional builtin which will render images (JPEG, GIF,
PNG, BMP, TIFF and WebP) to the terminal using block characters and ANSI
colour sequences.

## Usage

    <stdin> -> open-image -> <stdout>
    
    open-image file-path -> <stdout>

## Examples

As a method

    » cat example.png -> open-image
    
As a function

    » open-image example.png

## Detail

`open-image` will fail if STDOUT is not a TTY.

## See Also

* [commands/`open`](../commands/open.md):
  Open a file with a preferred handler
* [optional/`qr` ](../optional/qr.md):
  Creates a QR code from STDIN