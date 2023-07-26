package docs

func init() {

	Definition["alias"] = "# `alias` - Command Reference\n\n> Create an alias for a command\n\n## Description\n\n`alias` defines an alias for global usage\n\n## Usage\n\n```\nalias: alias=command parameter parameter\n\n!alias: command\n```\n\n## Examples\n\nBecause aliases are parsed into an array of parameters, you cannot put the\nentire alias within quotes. For example:\n\n```\n# bad :(\n» alias hw=\"out Hello, World!\"\n» hw\nexec: \"out\\\\ Hello,\\\\ World!\": executable file not found in $PATH\n\n# good :)\n» alias hw=out \"Hello, World!\"\n» hw\nHello, World!\n```\n\nNotice how only the command `out \"Hello, World!\"` is quoted in `alias` the\nsame way you would have done if you'd run that command \"naked\" in the command\nline? This is how `alias` expects it's parameters and where `alias` on Murex\ndiffers from `alias` in POSIX shells.\n\nIn some ways this makes `alias` a little less flexible than it might\notherwise be. However the design of this is to keep `alias` focused on it's\ncore objective. For any more advanced requirements you can use a `function`\ninstead.\n\n## Detail\n\n### Allowed characters\n\nAlias names can only include alpha-numeric characters, hyphen and underscore.\nThe following regex is used to validate the `alias`'s parameters:\n`^([-_a-zA-Z0-9]+)=(.*?)$`\n\n### Undefining an alias\n\nLike all other definable states in Murex, you can delete an alias with the\nbang prefix:\n\n```\n» alias hw=out \"Hello, World!\"\n» hw\nHello, World!\n\n» !alias hw\n» hw\nexec: \"hw\": executable file not found in $PATH\n```\n\n### Order of preference\n\nThere is an order of precedence for which commands are looked up:\n\n1. `runmode`: this is executed before the rest of the script. It is invoked by\n   the pre-compiler forking process and is required to sit at the top of any\n   scripts.\n\n1. `test` and `pipe` functions also alter the behavior of the compiler and thus\n   are executed ahead of any scripts.\n\n4. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function directly from the interactive command line\n   (however you can force an indirect call via `fexec`).\n\n2. Aliases - defined via `alias`. All aliases are global.\n\n3. Murex functions - defined via `function`. All functions are global.\n\n5. Variables (dollar prefixed) which are declared via `global`, `set` or `let`.\n   Also environmental variables too, declared via `export`.\n\n6. globbing: however this only applies for commands executed in the interactive\n   shell.\n\n7. Murex builtins.\n\n8. External executable files\n\nYou can override this order of precedence via the `fexec` and `exec` builtins.\n\n## Synonyms\n\n* `alias`\n* `!alias`\n\n\n## See Also\n\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`export`](../commands/export.md):\n  Define an environmental variable and set it's value\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`function`](../commands/function.md):\n  Define a function block\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg `*.txt`)\n* [`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [`method`](../commands/method.md):\n  Define a methods supported data-types\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [`source` ](../commands/source.md):\n  Import Murex code from another file of code block"

}
