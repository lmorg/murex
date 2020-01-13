package docs

func init() {

	Definition["<>"] = "# _murex_ Shell Docs\n\n## Command Reference: `<>` (read pipe)\n\n> Reads from a _murex_ named pipe\n\n## Description\n\nSometimes you will need to start a commandline with a _murex_ named pipe:\n\n    » <readpipe> -> match: foobar\n    \n> See the documentation on `pipe` for more details about _murex_ named pipes.\n\n## Usage\n\n    <example> -> <stdout>\n\n## Examples\n\nThe follow two examples function the same\n\n    » pipe: example\n    » bg { <example> -> match: 2 }\n    » a: <example> [1..3]\n    2\n    » !pipe: example\n\n## Synonyms\n\n* `<>`\n\n\n## See Also\n\n* [commands/`<stdin>` ](../commands/stdin.md):\n  Read the STDIN belonging to the parent code block\n* [commands/pipe](../commands/pipe.md):\n  "

}
