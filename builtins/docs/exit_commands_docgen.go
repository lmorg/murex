package docs

// This file was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).

func init() {

	Definition["exit"] = "# `exit`\n\n> Exit murex\n\n## Description\n\nExit's Murex with either a exit number of 0 (by default if no parameters\nsupplied) or a custom value specified by the first parameter.\n\n`exit` is not scope aware; if it is included in a function then the whole\nshell will still exist and not just that function.\n\n## Usage\n\n```\nexit\nexit number\n```\n\n## Examples\n\n```\n» exit\n```\n\n```\n» exit 42\n```\n\n## See Also\n\n* [`break`](../commands/break.md):\n  Terminate execution of a block within your processes scope\n* [`die`](../commands/die.md):\n  Terminate murex with an exit number of 1\n* [`null`](../commands/devnull.md):\n  null function. Similar to /dev/null\n\n<hr/>\n\nThis document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml)."

}
