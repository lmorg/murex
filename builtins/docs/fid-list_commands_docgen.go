package docs

func init() {

	Definition["fid-list"] = "# `fid-list` - Command Reference\n\n> Lists all running functions within the current Murex session\n\n## Description\n\n`fid-list` is a tool for outputting all the functions currently managed by that\nMurex session. Those functions could be Murex functions, builtins or any\nexternal executables launched from Murex.\n\nConceptually `fid-list` is a little like `ps` (on POSIX systems) however\n`fid-list` was not written to be POSIX compliant.\n\nMultiple flags cannot be used with each other.\n\n## Usage\n\n```\nfid-list [ flag ] -> <stdout>\n```\n\n`jobs` is an alias for `fid-list: --jobs`:\n```\njobs -> <stdout>\n```\n\n## Flags\n\n* `--background`\n    Returns a `json` map of background jobs\n* `--csv`\n    Output table in a `csv` format\n* `--help`\n    Outputs a list of parameters and a descriptions\n* `--jobs`\n    Show background and stopped jobs\n* `--jsonl`\n    Output table in a jsonlines (`jsonl`) format (defaulted to when piped)\n* `--stopped`\n    Returns a `json` map of stopped jobs\n* `--tty`\n    Force default TTY output even when piped\n\n## Detail\n\nBecause Murex is a multi-threaded shell, builtins are not forked processes\nlike in a traditional / POSIX shell. This means that you cannot use the\noperating systems default process viewer (eg `ps`) to list Murex functions.\nThis is where `fid-list` comes into play. It is used to view all the functions\nand processes that are managed by the current Murex session. That would\ninclude:\n\n* any aliases within Murex\n* public and private Murex functions\n* builtins (eg `fid-list` is a builtin command)\n* any external processes that were launched from within this shell session\n* any background functions or processes of any of the above\n\n## Synonyms\n\n* `fid-list`\n* `jobs`\n\n\n## See Also\n\n* [`*` (generic) ](../types/generic.md):\n  generic (primitive)\n* [`bexists`](../commands/bexists.md):\n  Check which builtins exist\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [`builtins`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex\n* [`csv` ](../types/csv.md):\n  CSV files (and other character delimited tables)\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`fg`](../commands/fg.md):\n  Sends a background process into the foreground\n* [`fid-kill`](../commands/fid-kill.md):\n  Terminate a running Murex function\n* [`fid-killall`](../commands/fid-killall.md):\n  Terminate _all_ running Murex functions\n* [`jobs`](../commands/fid-list.md):\n  Lists all running functions within the current Murex session\n* [`jsonl` ](../types/jsonl.md):\n  JSON Lines\n* [`murex-update-exe-list`](../commands/murex-update-exe-list.md):\n  Forces Murex to rescan $PATH looking for executables"

}
