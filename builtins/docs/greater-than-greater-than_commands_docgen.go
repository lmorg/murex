package docs

func init() {

	Definition[">>"] = "# _murex_ Shell Docs\n\n## Command Reference: `>>` (append file)\n\n> Writes STDIN to disk - appending contents if file already exists\n\n## Description\n\nRedirects output to file.\n\nIf a file already exists, the contents will be appended to existing contents.\nOtherwise a new file is created.\n\n## Usage\n\n    <stdin> -> >> filename\n\n## Examples\n\n    g * -> >> files.txt\n\n## Synonyms\n\n* `>>`\n* `fappend`\n\n\n## See Also\n\n* [commands/`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* [commands/`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [commands/`pipe`](../commands/pipe.md):\n  Manage _murex_ named pipes\n* [commands/`tmp`](../commands/tmp.md):\n  Create a temporary file and write to it"

}
