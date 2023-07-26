package docs

func init() {

	Definition["!"] = "# `!` (not) - Command Reference\n\n> Reads the STDIN and exit number from previous process and not's it's condition\n\n## Description\n\nReads the STDIN and exit number from previous process and not's it's condition.\n\n## Usage\n\n```\n<stdin> -> ! -> <stdout>\n```\n\n## Examples\n\n```\n» echo \"Hello, world!\" -> !\nfalse\n```\n\n```\n» false -> !\ntrue\n```\n\n## Synonyms\n\n* `!`\n\n\n## See Also\n\n* [`and`](../commands/and.md):\n  Returns `true` or `false` depending on whether multiple conditions are met\n* [`false`](../commands/false.md):\n  Returns a `false` value\n* [`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [`or`](../commands/or.md):\n  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.\n* [`true`](../commands/true.md):\n  Returns a `true` value"

}
