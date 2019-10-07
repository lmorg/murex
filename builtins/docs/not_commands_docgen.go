package docs

func init() {

	Definition["!"] = "# _murex_ Shell Docs\n\n## Command Reference: `!` (not)\n\n> Reads the STDIN and exit number from previous process and not's it's condition\n\n### Description\n\nReads the STDIN and exit number from previous process and not's it's condition.\n\n### Usage\n\n    <stdin> -> ! -> <stdout>\n\n### Examples\n\n    » echo \"Hello, world!\" -> !\n    false\n    \n    » false -> !\n    true\n\n### Synonyms\n\n* `!`\n\n\n### See Also\n\n* [commands/`and`](../commands/and.md):\n  Returns `true` or `false` depending on whether multiple conditions are met\n* [commands/`false`](../commands/false.md):\n  Returns a `false` value\n* [commands/`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [commands/`or`](../commands/or.md):\n  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.\n* [commands/`true`](../commands/true.md):\n  Returns a `true` value"

}
