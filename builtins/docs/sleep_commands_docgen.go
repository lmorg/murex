package docs

func init() {

	Definition["sleep"] = "# `sleep`  - Optional Command Reference\n\n> Suspends the shell for a number of seconds\n\n## Description\n\n`sleep` is an optional builtin which suspends the shell for a defined number\nof seconds.\n\n## Usage\n\n    sleep: integer\n\n## Examples\n\n    » sleep 5\n    # murex sleeps for 5 seconds\n\n## Detail\n\n`sleep` is very simplistic - particularly when compared to its GNU coreutil\n(for example) counterpart. If you want to use the `sleep` binary on Linux\nor similar platforms then you will need to launch with the `exec` builtin:\n\n    » exec: sleep 5\n\n## See Also\n\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`source` ](../commands/source.md):\n  Import Murex code from another file of code block\n* [`time` ](../commands/time.md):\n  Returns the execution run time of a command or block"

}
