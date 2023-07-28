package docs

func init() {

	Definition["open-image"] = "# `open-image` \n\n> Renders bitmap image data on your terminal\n\n## Description\n\n`open-image` is an optional builtin which will render images (JPEG, GIF,\nPNG, BMP, TIFF and WebP) to the terminal using block characters and ANSI\ncolour sequences.\n\n## Usage\n\n```\n<stdin> -> open-image -> <stdout>\n\nopen-image file-path -> <stdout>\n```\n\n## Examples\n\nAs a method\n\n```\n» cat example.png -> open-image\n```\n\nAs a function\n\n```\n» open-image example.png\n```\n\n## Detail\n\n`open-image` will fail if STDOUT is not a TTY.\n\n## See Also\n\n* [`open`](../commands/open.md):\n  Open a file with a preferred handler\n* [`qr` ](../optional/qr.md):\n  Creates a QR code from STDIN"

}
