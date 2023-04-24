package docs

func init() {

	Definition["false"] = "# `false` - Command Reference\n\n> Returns a `false` value\n\n## Description\n\nReturns a `false` value.\n\n## Usage\n\n    false -> <stdout>\n\n## Examples\n\nBy default, `false` also outputs the term \"false\":\n\n    » false\n    false\n    \nHowever you can suppress that with the silent flag:\n\n    » false -s\n\n## Flags\n\n* `-s`\n    silent - don't output the term \"false\"\n\n## See Also\n\n* [`!` (not)](../commands/not.md):\n  Reads the STDIN and exit number from previous process and not's it's condition\n* [`and`](../commands/and.md):\n  Returns `true` or `false` depending on whether multiple conditions are met\n* [`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [`or`](../commands/or.md):\n  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.\n* [`true`](../commands/true.md):\n  Returns a `true` value"

}
