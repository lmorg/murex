package docs

func init() {

	Definition["null"] = "# _murex_ Shell Docs\n\n## Command Reference: `null`\n\n> null function. Similar to /dev/null\n\n## Description\n\n`null` is a function that acts a little like the `null` data type and the\nUNIX /dev/null device.\n\n## Usage\n\n    <stdin> -> null\n\n## Examples\n\n    » out: \"Hello, world!\" -> null\n\n## Detail\n\nWhile this method does exist, a more idiomatic way to suppress STDOUT is to\nuse the named pipe property rather than piping to null:\n\n    » out: <null> \"Hello, world!\"\n\n## Synonyms\n\n* `null`\n\n\n## See Also\n\n* [commands/`break`](../commands/break.md):\n  terminate execution of a block within your processes scope\n* [commands/`die`](../commands/die.md):\n  Terminate murex with an exit number of 1\n* [commands/`exit`](../commands/exit.md):\n  Exit murex"

}
