package docs

func init() {

	Definition["qr"] = "# _murex_ Shell Guide\n\n## Command Reference: `qr` (optional)\n\n> Creates a QR code from STDIN\n\n### Description\n\n`qr` is an optional builtin which generates a PNG format image based on the\ninput from STDIN. `qr` must be run as a method.\n\n### Usage\n\n    <stdin> -> qr -> <stdout>\n\n### Examples\n\nWrite the PNG to disk\n\n    » out \"Hello, World!\" -> qr -> > qr.png\n    \nDisplay PNG in the terminal\n\n    » out \"Hello, World!\" -> qr -> open-image\n\n### Detail\n\n`qr` sets stdout's data-type to be \"image\", which is defined in with the\n`open-image` optional builtin. So if you have that disabled then you may\nhave to `cast` the output in some circumstances.\n\n### See Also\n\n* commands/[`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* commands/[`open-image` (optional)](../commands/open-image.md):\n  Renders bitmap image data on your terminal"

}
