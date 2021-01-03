package docs

func init() {

	Definition["esccli"] = "# _murex_ Shell Docs\n\n## Command Reference: `esccli`\n\n> Escapes an array so output is valid shell code\n\n## Description\n\n`esccli` takes an array and escapes any characters that might cause problems\nwhen pasted back into the terminal. Typically you'd want to use this against\ncommand parameters.\n\n## Usage\n\n    <stdin> -> esccli -> <stdout>\n    \n    esccli @array -> <stdout>\n\n## Examples\n\nAs a method\n\n    » alias foobar=out 'foo$b@r'\n    » alias -> [foobar]\n    [\n        \"out\",\n        \"foo$b@r\"\n    ]\n    » alias -> [foobar] -> esccli\n    out foo\\$b\\@r\n    \nAs a function\n\n    » alias -> [foobar] -> set: fb\n    » $fb\n    [\"out\",\"foo$b@r\"]\n    » esccli: @fb\n    out foo\\$b\\@r\n\n## See Also\n\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`alias`](../commands/alias.md):\n  Create an alias for a command\n* [commands/`escape`](../commands/escape.md):\n  Escape or unescapes input \n* [commands/`eschtml`](../commands/eschtml.md):\n  Encode or decodes text for HTML\n* [commands/`escurl`](../commands/escurl.md):\n  Encode or decodes text for the URL\n* [commands/`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character"

}
