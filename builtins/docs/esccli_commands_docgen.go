package docs

func init() {

    Definition["esccli"] = "# `esccli`\n\n> Escapes an array so output is valid shell code\n\n## Description\n\n`esccli` takes an array and escapes any characters that might cause problems\nwhen pasted back into the terminal. Typically you'd want to use this against\ncommand parameters.\n\n## Usage\n\n```\n<stdin> -> esccli -> <stdout>\n\nesccli @array -> <stdout>\n```\n\n## Examples\n\nAs a method\n\n```\n» alias foobar=out 'foo$b@r'\n» alias -> [foobar]\n[\n    \"out\",\n    \"foo$b@r\"\n]\n» alias -> [foobar] -> esccli\nout foo\\$b\\@r\n```\n\nAs a function\n\n```\n» alias -> [foobar] -> set fb\n» $fb\n[\"out\",\"foo$b@r\"]\n» esccli @fb\nout foo\\$b\\@r\n```\n\n## See Also\n\n* [`alias`](../commands/alias.md):\n  Create an alias for a command\n* [`escape`](../commands/escape.md):\n  Escape or unescape input\n* [`eschtml`](../commands/eschtml.md):\n  Encode or decodes text for HTML\n* [`escurl`](../commands/escurl.md):\n  Encode or decodes text for the URL\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table"

}