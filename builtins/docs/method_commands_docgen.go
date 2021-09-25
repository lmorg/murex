package docs

func init() {

	Definition["method"] = "# _murex_ Shell Docs\n\n## Command Reference: `method`\n\n> Define a methods supported data-types\n\n## Description\n\n`method` defines what the typical data type would be for a function's STDIN\nand STDOUT.\n\n## Usage\n\n    method: define name { code-block }\n\n## Examples\n\n    method: define name {\n        \"Stdin\": \"@Any\",\n        \"Stdout\": \"json\"\n    }\n\n## See Also\n\n* [commands/`alias`](../commands/alias.md):\n  Create an alias for a command\n* [commands/`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [commands/`function`](../commands/function.md):\n  Define a function block\n* [commands/`private`](../commands/private.md):\n  Define a private function block\n* [commands/`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_"

}
