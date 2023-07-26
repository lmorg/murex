package docs

func init() {

	Definition["source"] = "# `source`  - Command Reference\n\n> Import Murex code from another file of code block\n\n## Description\n\n`source` imports code from another file or code block. It can be used as either\nan \"import\" / \"include\" directive (eg Python, Go, C, etc) or an \"eval\" (eg\nPython, Perl, etc).\n\n## Usage\n\nExecute source from STDIN\n\n```\n<stdin> -> source\n```\n\nExecute source from a file\n\n```\nsource: filename.mx\n```\n\nExecute a code block from parameter\n\n```\nsource: { code-block }\n```\n\n## Examples\n\nExecute source from stdin:\n\n```\n» tout: block { out: \"Hello, world!\" } -> source\nHello, world!\n```\n\nExecute source from file:\n\n```\n» tout: block { out: \"Hello, world!\" } |> example.mx\n» source: example.mx\nHello, world!\n```\n\nExecute a code block from parameter\n\n```\n» source { out: \"Hello, world!\" }\nHello, world!\n```\n\n## Synonyms\n\n* `source`\n* `.`\n\n\n## See Also\n\n* [`args` ](../commands/args.md):\n  Command line flag parser for Murex shell scripting\n* [`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`function`](../commands/function.md):\n  Define a function block\n* [`murex-parser` ](../commands/murex-parser.md):\n  Runs the Murex parser against a block of code \n* [`private`](../commands/private.md):\n  Define a private function block\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex\n* [`version` ](../commands/version.md):\n  Get Murex version"

}
