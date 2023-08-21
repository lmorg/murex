package docs

// This file was generated from [builtins/core/processes/bgfg_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/bgfg_doc.yaml).

func init() {

	Definition["bg"] = "# `bg`\n\n> Run processes in the background\n\n## Description\n\n`bg` supports two modes: it can either be run as a function block which will\nexecute in the background, or it can take stopped processes and daemonize\nthem.\n\n## Usage\n\nPOSIX only:\n ```\n bg { code block }\n\n bg fid\n ```\n\n## Examples\n\n\n```\nbg { sleep 5; out \"Morning\" }\n```\n\n## Detail\n\nThe examples above will work on any system (Windows included). However the\n`ctrl+z` usage of backgrounding a stopped process (like Bash) is only\nsupported on POSIX systems due to the limitation of required signals on\nnon-platforms. This means the usage described in the examples is cross\ncross platform while `bg int` currently does not work on Windows nor Plan 9.\n\n## See Also\n\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`fg`](../commands/fg.md):\n  Sends a background process into the foreground\n* [`fid-kill`](../commands/fid-kill.md):\n  Terminate a running Murex function\n* [`fid-killall`](../commands/fid-killall.md):\n  Terminate _all_ running Murex functions\n* [`fid-list`](../commands/fid-list.md):\n  Lists all running functions within the current Murex session\n* [`jobs`](../commands/fid-list.md):\n  Lists all running functions within the current Murex session\n\n<hr/>\n\nThis document was generated from [builtins/core/processes/bgfg_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/bgfg_doc.yaml)."

}
