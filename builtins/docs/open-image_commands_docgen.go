package docs

func init() {

	Definition["open-image"] = "# _murex_ Shell Docs\n\n## Command Reference: `open-image` \n\n> Renders bitmap image data on your terminal\n\n## Description\n\n`open-image` is an optional builtin which will render images (JPEG, GIF,\nPNG, BMP, TIFF and WebP) to the terminal using block characters and ANSI\ncolour sequences.\n\n## Usage\n\n    <stdin> -> open-image -> <stdout>\n    \n    open-image file-path -> <stdout>\n\n## Examples\n\nAs a method\n\n    » cat example.png -> open-image\n    \nAs a function\n\n    » open-image example.png\n\n## Detail\n\n`open-image` will fail if STDOUT is not a TTY.\n\n## See Also\n\n* [optional/`qr` ](../optional/qr.md):\n  Creates a QR code from STDIN\n* [commands/open](../commands/open.md):\n  "

}
