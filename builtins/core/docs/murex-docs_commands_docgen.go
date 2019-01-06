package docs

func init() {

	Definition["murex-docs"] = "# _murex_ Language Guide\n\n## Command Reference: `murex-docs`\n\n> Displays the man pages for _murex_ builtins\n\n### Description\n\nDisplays the man pages for _murex_ builtins.\n\n### Usage\n\n    murex-docs: [ flag ] command -> <stdout>\n\n### Examples\n\n    # Output this man page\n    murex-docs: murex-docs\n\n### Flags\n\n* `--digest`\n    returns an abridged description of the command rather than the entire help page.\n\n### Detail\n\nThese man pages are compiled into the _murex_ executable.\n\n### See Also\n\n* [`(` (brace quote)](../docs/commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`>>` (write to new or appended file)](../docs/commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`>` (write to new or truncated file)](../docs/commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists    \n* [`err`](../docs/commands/err.md):\n  Print a line to the STDERR\n* [`out`](../docs/commands/out.md):\n  `echo` a string to the STDOUT with a trailing new line character\n* [`tout`](../docs/commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* [`tread`](../docs/commands/tread.md):\n  `read` a line of input from the user and store as a user defined *typed* variable    \n* [cast](../docs/commands/commands/cast.md):\n  \n* [sprintf](../docs/commands/commands/sprintf.md):\n  "

}
