package docs

func init() {

	Definition["function"] = "# _murex_ Shell Docs\n\n## Command Reference: `function`\n\n> Define a function block\n\n## Description\n\n`function` defines a block of code as a function\n\n## Usage\n\nDefine a function:\n\n    function: name { code-block }\n    \nDefine a function with variable names defined (**default value** and\n**description** are optional parameters):\n\n    function: name (\n        variable1: data-type [default-value] \"description\",\n        variable2: data-type [default-value] \"description\"\n    ) {\n        code-block\n    }\n    \nUndefine a function:\n\n    !function: command\n\n## Examples\n\n    » function hw { out \"Hello, World!\" }\n    » hw\n    Hello, World!\n    \n    » !function hw\n    » hw\n    exec: \"hw\": executable file not found in $PATH\n\n## Detail\n\n### Allowed characters\n\nFunction names can only include any characters apart from dollar (`$`).\nThis is to prevent functions from overwriting variables (see the order of\npreference below).\n\n### Undefining a function\n\nLike all other definable states in _murex_, you can delete a function with\nthe bang prefix (see the example above).\n\n### Using parameterized variable names\n\nBy default, if you wanted to query the parameters passed to a _murex_ function\nyou would have to use either:\n\n* the Bash syntax where of `$2` style numbered reserved variables,\n\n* and/or the _murex_ convention of `$PARAM` / `$ARGS` arrays (see **reserved-vars**\n  document below),\n  \n* and/or the older _murex_ convention of the `args` builtin for any flags.\n\nStarting from _murex_ `2.7.x` it's been possible to declare parameters from\nwithin the function declaration:\n\n    function: name (\n        variable1: data-type [default-value] \"description\",\n        variable2: data-type [default-value] \"description\"\n    ) {\n        code-block\n    }\n    \n#### Syntax\n\nFirst off, the syntax doesn't have to follow exactly as above:\n\n* Variables shouldn't be prefixed with a dollar (`$`). This is a little like\n  declaring variables via `set`, etc. However it should be followed by a colon\n  (`:`) or comma (`,`). Normal rules apply with regards to allowed characters\n  in variable names: limited to ASCII letters (upper and lower case), numbers,\n  underscore (`_`), and hyphen (`-`). Unicode characters as variable names are\n  not currently supported.\n\n* **data-type** is the _murex_ data type. This is an optional field in version\n  `2.8.x` (defaults to `str`) but is required in `2.7.x`.\n\n* The default value must be inside square brackets (`[...]`). Any value is\n  allowed (including Unicode) _except_ for carriage returns / new lines (`\\r`,\n  `\\n`) and a closing square bracket (`]`) as the latter would indicate the end\n  of this field. You cannot escape these characters either.\n\n  This field is optional.\n\n* The description must sit inside double quotes (`\"...\"`). Any value is allowed\n  (including Unicode) _except_ for carriage returns / new lines (`\\r`, `\\n`)\n  and double quotes (`\"`) as the latter would indicate the end of this field.\n  You cannot escape these characters either.\n\n  This field is optional.\n\n* You do not need a new line between each parameter, however you do need to\n  separate them with a comma (like with JSON, there should not be a trailing\n  comma at the end of the parameters). Thus the following is valid:\n  `variable1: data-type, variable2: data-type`.\n\n#### Variables\n\nAny variable name you declare in your function declaration will be exposed in\nyour function body as a local variable. For example:\n\n    function: hello (name: str) {\n        out: \"Hello $name, pleased to meet you.\"\n    }\n    \nIf the function isn't called with the complete list of parameters and it is\nrunning in the foreground (ie not part of `autocomplete`, `event`, `bg`, etc)\nthen you will be prompted for it's value. That could look something like this:\n\n    » function: hello (name: str) {\n    »     out: \"Hello $name, pleased to meet you.\"\n    » }\n    \n    » hello\n    Please enter a value for 'name': Bob\n    Hello Bob, pleased to meet you.\n    \n(in this example you typed `Bob` when prompted)\n\n#### Data-Types\n\nThis is the _murex_ data type of the variable. From version `2.8.x` this field\nis optional and will default to `str` when omitted.\n\nThe advantage of setting this field is that values are type checked and the\nfunction will fail early if an incorrect value is presented. For example:\n\n    » function: age (age: int) { out: \"$age is a great age.\" }\n    \n    » age\n    Please enter a value for 'age': ten\n    Error in `age` ( 2,1): cannot convert parameter 1 'ten' to data type 'int'\n    \n    » age ten\n    Error in `age` ( 2,1): cannot convert parameter 1 'ten' to data type 'int'\n    \nHowever it will try to automatically convert values if it can:\n\n    » age 1.2\n    1 is a great age.\n    \n#### Default values\n\nDefault values are only relevant when functions are run interactively. It\nallows the user to press enter without inputting a value:\n\n    » function: hello (name: str [John]) { out: \"Hello $name, pleased to meet you.\" }\n    \n    » hello\n    Please enter a value for 'name' [John]: \n    Hello John, pleased to meet you.\n    \nHere no value was entered so `$name` defaulted to `John`.\n\nDefault values will not auto-populate when the function is run in the\nbackground. For example:\n\n    » bg {hello}\n    Error in `hello` ( 2,2): cannot prompt for parameters when a function is run in the background: too few parameters\n    \n#### Description\n\nDescriptions are only relevant when functions are run interactively. It allows\nyou to define a more useful prompt should that function be called without\nsufficient parameters. For example:\n\n    » function hello (name: str \"What is your name?\") { out \"Hello $name\" }\n    \n    » hello\n    What is your name?: Sally\n    Hello Sally\n    \n### Order of precedence\n\nThere is an order of precedence for which commands are looked up:\n\n1. `test` and `pipe` functions because they alter the behavior of the compiler\n\n2. Aliases - defined via `alias`. All aliases are global\n\n3. _murex_ functions - defined via `function`. All functions are global\n\n4. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function from the interactive command line\n\n5. variables (dollar prefixed) - declared via `set` or `let`\n\n6. auto-globbing prefix: `@g`\n\n7. murex builtins\n\n8. external executable files\n\n## Synonyms\n\n* `function`\n* `!function`\n\n\n## See Also\n\n* [user-guide/Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by _murex_\n* [commands/`alias`](../commands/alias.md):\n  Create an alias for a command\n* [commands/`args` ](../commands/args.md):\n  Command line flag parser for _murex_ shell scripting\n* [commands/`exec`](../commands/exec.md):\n  Runs an executable\n* [commands/`export`](../commands/export.md):\n  Define an environmental variable and set it's value\n* [commands/`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [commands/`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [commands/`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [commands/`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable\n* [commands/`method`](../commands/method.md):\n  Define a methods supported data-types\n* [commands/`private`](../commands/private.md):\n  Define a private function block\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [commands/`source` ](../commands/source.md):\n  Import _murex_ code from another file of code block\n* [commands/`version` ](../commands/version.md):\n  Get _murex_ version"

}
