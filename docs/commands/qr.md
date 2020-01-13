# _murex_ Shell Docs

## Command Reference: `qr` (optional)

> Creates a QR code from STDIN

## Description

`qr` is an optional builtin which generates a PNG format image based on the
input from STDIN. `qr` must be run as a method.

## Usage

    <stdin> -> qr -> <stdout>

## Examples

Write the PNG to disk

    » out "Hello, World!" -> qr -> > qr.png
    
Display PNG in the terminal

    » out "Hello, World!" -> qr -> open-image

## Detail

`qr` sets stdout's data-type to be "image", which is defined in with the
`open-image` optional builtin. So if you have that disabled then you may
have to `cast` the output in some circumstances.

## See Also

* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`open-image` (optional)](../commands/open-image.md):
  Renders bitmap image data on your terminal