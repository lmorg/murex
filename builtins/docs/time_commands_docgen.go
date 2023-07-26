package docs

func init() {

	Definition["time"] = "# `time`  - Command Reference\n\n> Returns the execution run time of a command or block\n\n## Description\n\n`time` is an optional builtin which runs a command or block of code and\nreturns it's running time.\n\n## Usage\n\n```\ntime: command parameters -> <stderr>\n\ntime: { code-block } -> <stderr>\n```\n\n## Examples\n\n```\n» time: sleep 5\n5.000151513\n\n» time { out \"Going to sleep\"; sleep 5; out \"Waking up\" }\nGoing to sleep\nWaking up\n5.000240977\n```\n\n## Detail\n\n`time`'s output is written to STDERR. However any output and errors written\nby the commands executed by time will also be written to `time`'s STDOUT\nand STDERR as usual.\n\n## See Also\n\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`sleep` ](../optional/sleep.md):\n  Suspends the shell for a number of seconds\n* [`source` ](../commands/source.md):\n  Import Murex code from another file of code block"

}
