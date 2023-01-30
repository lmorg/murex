package docs

func init() {

	Definition["exit"] = "# _murex_ Shell Docs\n\n## Command Reference: `exit`\n\n> Exit murex\n\n## Description\n\nExit's _murex_ with either a exit number of 0 (by default if no parameters\nsupplied) or a custom value specified by the first parameter.\n\n`exit` is not scope aware; if it is included in a function then the whole\nshell will still exist and not just that function.\n\n## Usage\n\n    exit\n    exit number\n\n## Examples\n\n    » exit\n    \n    » exit 42\n\n## See Also\n\n* [`break`](../commands/break.md):\n  terminate execution of a block within your processes scope\n* [`die`](../commands/die.md):\n  Terminate murex with an exit number of 1\n* [`null`](../commands/devnull.md):\n  null function. Similar to /dev/null"

}
