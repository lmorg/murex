package docs

func init() {

	Definition["fg"] = "# _murex_ Shell Guide\n\n## Command Reference: `fg`\n\n> Sends a background process into the foreground\n\n### Description\n\n`fg` resumes a stopped process and sends it into the foreground.\n\n### Usage\n\nPOSIX only:\n\n    fg fid\n\n### Detail\n\nThis builtin is only supported on POSIX systems. There is no support planned\nfor Windows (due to the kernel not supporting the right signals) nor Plan 9.\n\n### See Also\n\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [fid-kill](../commands/fid-kill.md):\n  \n* [fid-killall](../commands/fid-killall.md):\n  \n* [fid-list](../commands/fid-list.md):\n  "

}
