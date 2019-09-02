package docs

func init() {

	Definition["f"] = "# _murex_ Shell Guide\n\n## Command Reference: `f`\n\n> Lists objects (eg files) in the current working directory\n\n### Description\n\nLists objects (eg files) in the current working directory.\n\n### Usage\n\n    f: options -> <stdout>\n\n### Examples\n\n    # return only directories:\n    f: +d\n    \n    # return file and directories but exclude symlinks:\n    f: +d +f -s\n\n### Flags\n\n* `d`\n    directories\n* `f`\n    files\n* `s`\n    symlinks (automatically included with files and directories)\n\n### Detail\n\nBy default `f` will return no results. It is then your responsibility to select\nwhich types of objects to be included or excluded from the results.\n\n### See Also\n\n* [commands/`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [commands/`rx`](../commands/rx.md):\n  Regexp pattern matching for file system objects (eg '.*\\.txt')"

}
