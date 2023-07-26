package docs

func init() {

	Definition[">"] = "# `>` (truncate file) - Command Reference\n\n> Writes STDIN to disk - overwriting contents if file already exists\n\n## Description\n\nRedirects output to file.\n\nIf a file already exists, the contents will be truncated (overwritten).\nOtherwise a new file is created.\n\n## Usage\n\n```\n<stdin> |> filename\n```\n\n## Examples\n\n```\ng * |> files.txt\n```\n\n## Synonyms\n\n* `>`\n* `fwrite`\n\n\n## See Also\n\n* [Arrow Pipe (`->`) Token](../parser/pipe-arrow.md):\n  Pipes STDOUT from the left hand command to STDIN of the right hand command\n* [POSIX Pipe (`|`) Token](../parser/pipe-posix.md):\n  Pipes STDOUT from the left hand command to STDIN of the right hand command\n* [STDERR Pipe (`?`) Token](../parser/pipe-err.md):\n  Pipes STDERR from the left hand command to STDIN of the right hand command\n* [`<>` / `read-named-pipe`](../commands/namedpipe.md):\n  Reads from a Murex named pipe\n* [`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg `*.txt`)\n* [`pipe`](../commands/pipe.md):\n  Manage Murex named pipes\n* [`tmp`](../commands/tmp.md):\n  Create a temporary file and write to it"

}
