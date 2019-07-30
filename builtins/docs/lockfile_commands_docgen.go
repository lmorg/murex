package docs

func init() {

	Definition["lockfile"] = "# _murex_ Language Guide\n\n## Command Reference: `lockfile`\n\n> Create and manage lock files\n\n### Description\n\n`lockfile` is used to create and manage lock files\n\n### Usage\n\nCreate a lock file with the name `identifier`\n\n    lockfile: lock identifier\n    \nDelete a lock file with the name `identifier`\n\n    lockfile: unlock identifier\n    \nWait until lock file with the name `identifier` has been deleted\n\n    lockfile: wait identifier\n    \nOutput the the file name and path of a lock file with the name `identifier`\n\n    lockfile: path identifier -> <stdout>\n\n### Examples\n\n    lockfile lock example\n    out \"lock file created: ${lockfile path example}\"\n    \n    bg {\n        sleep 10\n        lockfile unlock example\n    }\n    \n    out \"waiting for lock file to be deleted (sleep 10 seconds)....\"\n    lockfile wait example\n    out \"lock file gone!\"\n\n### See Also\n\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [`out`](../commands/out.md):\n  `echo` a string to the STDOUT with a trailing new line character\n* [`sleep` (optional)](../commands/sleep.md):\n  Suspends the shell for a number of seconds"

}
