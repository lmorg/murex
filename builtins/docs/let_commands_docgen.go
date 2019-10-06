package docs

func init() {

	Definition["let"] = "# _murex_ Shell Docs\n\n## Command Reference: `let`\n\n> Evaluate a mathematical function and assign to variable\n\n### Description\n\n`let` evaluates a mathematical function and then assigns it to a locally\nscoped variable (like `set`)\n\n### Usage\n\n    let var_name=evaluation\n    \n    let var_name++\n    \n    let var_name--\n\n### Examples\n\n    » let: age=18\n    » $age\n    18\n    \n    » let: age++\n    » $age\n    19\n    \n    » let: under18=age<18\n    » $under18\n    false\n    \n    » let: under21 = age < 21\n    » $under21\n    true\n\n### Detail\n\n#### Variables\n\nThere are two ways you can use variables with the math functions. Either by\nstring interpolation like you would normally with any other function, or\ndirectly by name.\n\nString interpolation:\n\n    » set abc=123\n    » = $abc==123\n    true\n    \nDirectly by name:\n\n    » set abc=123\n    » = abc==123\n    false\n    \nTo understand the difference between the two, you must first understand how\nstring interpolation works; which is where the parser tokenised the parameters\nlike so\n\n    command line: = $abc==123\n    token 1: command (name: \"=\")\n    token 2: parameter 1, string (content: \"\")\n    token 3: parameter 1, variable (name: \"abc\")\n    token 4: parameter 1, string (content: \"==123\")\n    \nThen when the command line gets executed, the parameters are compiled on demand\nsimilarly to this crude pseudo-code\n\n    command: \"=\"\n    parameters 1: concatenate(\"\", GetValue(abc), \"==123\")\n    output: \"=\" \"123==123\"\n    \nThus the actual command getting run is literally `123==123` due to the variable\nbeing replace **before** the command executes.\n\nWhereas when you call the variable by name it's up to `=` or `let` to do the\nvariable substitution.\n\n    command line: = abc==123\n    token 1: command (name: \"=\")\n    token 2: parameter 1, string (content: \"abc==123\")\n    \n    command: \"=\"\n    parameters 1: concatenate(\"abc==123\")\n    output: \"=\" \"abc==123\"\n    \nThe main advantage (or disadvantage, depending on your perspective) of using\nvariables this way is that their data-type is preserved.\n\n    » set str abc=123\n    » = abc==123\n    false\n    \n    » set int abc=123\n    » = abc==123\n    true\n    \nUnfortunately is one of the biggest areas in _murex_ where you'd need to be\ncareful. The simple addition or omission of the dollar prefix, `$`, can change\nthe behavior of `=` and `let`.\n\n#### Strings\n\nBecause the usual _murex_ tools for encapsulating a string (`\"`, `'` and `()`)\nare interpreted by the shell language parser, it means we need a new token for\nhandling strings inside `=` and `let`. This is where backtick comes to our\nrescue.\n\n    » set str abc=123\n    » = abc==`123`\n    true\n    \nPlease be mindful that if you use string interpolation then you will need to\ninstruct `=` and `let` that your field is a string\n\n    » set str abc=123\n    » = `$abc`==`123`\n    true\n    \n#### Best practice recommendation\n\nAs you can see from the sections above, string interpolation offers us some\nconveniences when comparing variables of differing data-types, such as a `str`\ntype with a number (eg `num` or `int`). However it makes for less readable code\nwhen just comparing strings. Thus the recommendation is to avoid using string\ninterpolation except only where it really makes sense (ie use it sparingly).\n\n#### Non-boolean logic\n\nThus far the examples given have been focused on comparisons however `=` and\n`let` supports all the usual arithmetic operators:\n\n    » = 10+10\n    20\n    \n    » = 10/10\n    1\n    \n    » = (4 * (3 + 2))\n    20\n    \n    » = `foo`+`bar`\n    foobar\n    \n#### Read more\n\n_murex_ uses the [govaluate package](https://github.com/Knetic/govaluate). More information can be found in it's [manual](https://github.com/Knetic/govaluate/blob/master/MANUAL.md).\n\n#### Scoping\n\nVariables are only scoped inside the code block they're defined in (or any\nchildren of that code block). For example `$foo` will return an empty string in\nthe following code because it's defined within a `try` block then being queried\noutside of the `try` block:\n\n    » try {\n    »     set foo=bar\n    » }\n    » out \"foo: $foo\"\n    foo:\n    \nHowever if we define `$foo` above the `try` block then it's value will be changed\neven though it is being set inside the `try` block:\n\n    » set foo\n    » try {\n    »     set foo=bar\n    » }\n    » out \"foo: $foo\"\n    foo: bar\n    \nSo unlike the previous example, this will return `bar`.\n\nWhere `global` differs from `set` is that the variables defined with `global`\nwill scoped at the global shell level (please note this is not the same as\nenvironmental variables!) so will cascade down through all scoped code-blocks\nincluding those running in other threads.\n\nIt's also worth remembering that any variable defined using `set` in the shell's\nFID (ie in the interactive shell) is literally the same as using `global`\n\nExported variables (defined via `export`) are system environmental variables.\nInside _murex_ environmental variables behave much like `global` variables\nhowever their real purpose is passing data to external processes. For example\n`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):\n\n    » export foo=bar\n    » env -> grep foo\n    foo=bar\n    \n#### Function Names\n\nAs a security feature function names cannot include variables. This is done to\nreduce the risk of code executing by mistake due to executables being hidden\nbehind variable names.\n\nInstead _murex_ will assume you want the output of the variable printed:\n\n    » out \"Hello, world!\" -> set hw\n    » $hw\n    Hello, world!\n    \nOn the rare occasions you want to force variables to be expanded inside a\nfunction name, then call that function via `exec`:\n\n    » set cmd=grep\n    » ls -> exec: $cmd main.go\n    main.go\n    \nThis only works for external executables. There is currently no way to call\naliases, functions nor builtins from a variable and even the above `exec` trick\nis considered bad form because it reduces the readability of your shell scripts.\n\n#### Usage Inside Quotation Marks\n\nLike with Bash, Perl and PHP: _murex_ will expand the variable when it is used\ninside a double quotes but will escape the variable name when used inside single\nquotes:\n\n    » out \"$foo\"\n    bar\n    \n    » out '$foo'\n    $foo\n    \n    » out ($foo)\n    bar\n\n### See Also\n\n* [commands/`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [commands/`=` (arithmetic evaluation)](../commands/equ.md):\n  Evaluate a mathematical function\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`export`](../commands/export.md):\n  Define a local variable and set it's value\n* [commands/`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [commands/`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value"

}
