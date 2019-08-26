package docs

func init() {

	Definition["rx"] = "# _murex_ Shell Guide\n\n## Command Reference: `rx`\n\n> Regexp pattern matching for file system objects (eg '.*\\.txt')\n\n### Description\n\nReturns a list of files and directories that match a regexp pattern.\n\nOutput is a JSON list.\n\n### Usage\n\n    rx: pattern -> <stdout>\n\n### Examples\n\n    # inline regex file matching\n    cat: @{ rx: '.*\\.txt' }\n    \n    # writing a list of files to disk\n    rx: '.*\\.go' -> > filelist.txt\n    \n    # checking if any files exist\n    if { rx: somefiles.* } then {\n        # files exist\n    }\n    \n    # checking if no files exist\n    !if { rx: somefiles.* } then {\n        # files do not exist\n    }\n\n### Detail\n\nUnlike globbing (`g`) which can traverse directories (eg `g: /path/*`), `rx` is\nonly designed to match file system objects in the current working directory.\n\n`rx` uses Go (lang)'s standard regexp engine.\n\n### See Also\n\n* [`f`](../commands/f.md):\n  Lists objects (eg files) in the current working directory\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [`match`](../commands/match.md):\n  Match an exact value in an array\n* [`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings"

}
