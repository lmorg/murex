package docs

func init() {

	Definition["!bz2"] = "# `!bz2`  - Optional Command Reference\n\n> Decompress a bz2 file\n\n## Description\n\n`!bz2` is an optional builtin for decompressing a bz2 stream from STDIN.\n\n## Usage\n\n```\n<stdin> -> !bz2 -> <stdout>\n```\n\n## Detail\n\nCurrently there is no support for compressing a stream using bz2.\n\n## Synonyms\n\n* `!bz2`\n\n\n## See Also\n\n* [`base64` ](../optional/base64.md):\n  Encode or decode a base64 string\n* [`escape`](../commands/escape.md):\n  Escape or unescape input \n* [`esccli`](../commands/esccli.md):\n  Escapes an array so output is valid shell code\n* [`eschtml`](../commands/eschtml.md):\n  Encode or decodes text for HTML\n* [`escurl`](../commands/escurl.md):\n  Encode or decodes text for the URL\n* [`gz` ](../optional/gz.md):\n  Compress or decompress a gzip file"

}
