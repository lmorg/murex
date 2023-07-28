package docs

func init() {

	Definition["lockfile"] = "# `lockfile`\n\n> Create and manage lock files\n\n## Description\n\n`lockfile` is used to create and manage lock files\n\n## Usage\n\nCreate a lock file with the name `identifier`\n\n```\nlockfile: lock identifier\n```\n\nDelete a lock file with the name `identifier`\n\n```\nlockfile: unlock identifier\n```\n\nWait until lock file with the name `identifier` has been deleted\n\n```\nlockfile: wait identifier\n```\n\nOutput the the file name and path of a lock file with the name `identifier`\n\n```\nlockfile: path identifier -> <stdout>\n```\n\n## Examples\n\n```\nlockfile: lock example\nout: \"lock file created: ${lockfile path example}\"\n\nbg {\n    sleep: 10\n    lockfile: unlock example\n}\n\nout: \"waiting for lock file to be deleted (sleep 10 seconds)....\"\nlockfile: wait example\nout: \"lock file gone!\"\n```\n\n## See Also\n\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character"

}
