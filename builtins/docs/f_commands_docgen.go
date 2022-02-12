package docs

func init() {

	Definition["f"] = "# _murex_ Shell Docs\n\n## Command Reference: `f`\n\n> Lists objects (eg files) in the current working directory\n\n## Description\n\nLists objects (eg files) in the current working directory.\n\n## Usage\n\n    f: options -> <stdout>\n    \n    <stdin> -> f: options -> <stdout>\n\n## Examples\n\nReturn only directories:\n\n    f: +d\n    \nReturn file and directories but exclude symlinks:\n\n    f: +d +f -s\n    \nCompare list against files (eg created by `g`) against conditions set by `f`:\n\n    g /* -> f +f\n\n## Flags\n\n* `d`\n    directories\n* `f`\n    files\n* `s`\n    symlinks (automatically included with files and directories)\n\n## Detail\n\nBy default `f` will return no results. It is then your responsibility to select\nwhich types of objects to be included or excluded from the results.\n\n## See Also\n\n* [commands/`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [commands/`rx`](../commands/rx.md):\n  Regexp pattern matching for file system objects (eg '.*\\.txt')"

}
