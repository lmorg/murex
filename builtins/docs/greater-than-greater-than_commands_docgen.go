package docs

func init() {

	Definition[">>"] = "# _murex_ Shell Docs\n\n## Command Reference: `>>` (append file)\n\n> Writes STDIN to disk - appending contents if file already exists\n\n## Description\n\nRedirects output to file.\n\nIf a file already exists, the contents will be appended to existing contents.\nOtherwise a new file is created.\n\n## Usage\n\n    <stdin> >> filename\n\n## Examples\n\n    g * >> files.txt\n\n## Synonyms\n\n* `>>`\n* `fappend`\n\n\n## See Also\n\n* [parser/Arrow Pipe (`->`) Token](../parser/pipe-arrow.md):\n  Pipes STDOUT from the left hand command to STDIN of the right hand command\n* [parser/POSIX Pipe (`|`) Token](../parser/pipe-posix.md):\n  Pipes STDOUT from the left hand command to STDIN of the right hand command\n* [parser/STDERR Pipe (`?`) Token](../parser/pipe-err.md):\n  Pipes STDERR from the left hand command to STDIN of the right hand command\n* [commands/`<>` / `read-named-pipe`](../commands/namedpipe.md):\n  Reads from a _murex_ named pipe\n* [commands/`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* [commands/`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [commands/`pipe`](../commands/pipe.md):\n  Manage _murex_ named pipes\n* [commands/`tmp`](../commands/tmp.md):\n  Create a temporary file and write to it"

}
