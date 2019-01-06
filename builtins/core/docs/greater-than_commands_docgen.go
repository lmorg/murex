package docs

func init() {

	Definition[">"] = "# _murex_ Language Guide\n\n## Command Reference: `>` (write to new or truncated file)\n\n> Writes STDIN to disk - overwriting contents if file already exists    \n\n### Description\n\nRedirects output to file.\n\nIf a file already exists, the contents will be truncated (overwritten).\nOtherwise a new file is created.\n\n### Usage\n\n    <stdin> -> > filename\n\n### Examples\n\n    g * -> > files.txt\n\n### Synonyms\n\n* `>`\n\n\n### See Also\n\n* [`>>` (write to new or appended file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [pipe](../commands/pipe.md):\n  \n* [tmp](../commands/tmp.md):\n  "

}
