package docs

func init() {

	Definition["fg"] = "# _murex_ Shell Docs\n\n## Command Reference: `fg`\n\n> Sends a background process into the foreground\n\n## Description\n\n`fg` resumes a stopped process and sends it into the foreground.\n\n## Usage\n\nPOSIX only:\n\n    fg fid\n\n## Detail\n\nThis builtin is only supported on POSIX systems. There is no support planned\nfor Windows (due to the kernel not supporting the right signals) nor Plan 9.\n\n## See Also\n\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`fid-kill`](../commands/fid-kill.md):\n  Terminate a running _murex_ function\n* [`fid-killall`](../commands/fid-killall.md):\n  Terminate _all_ running _murex_ functions\n* [`fid-list`](../commands/fid-list.md):\n  Lists all running functions within the current _murex_ session\n* [`jobs`](../commands/fid-list.md):\n  Lists all running functions within the current _murex_ session"

}
