package docs

func init() {

	Definition["rand"] = "# _murex_ Shell Docs\n\n## Command Reference: `rand`\n\n> Random field generator\n\n## Description\n\n`rand` can generate random numbers, strings and other data types.\n\n## Usage\n\n    rand data-type [ max-value ]\n\n## Examples\n\nRandom integer: 64-bit on 64-bit machines\n\n    rand int\n    \nRandom number between \n\n## Detail\n\n### Security\n\nWARNING: is should be noted that while `rand` can produce random numbers and\nstrings which might be useful for password generation, neither the RNG nor the\nthe random string generator (which is ostensibly the same RNG but applied to an\narray of bytes within the range of printable ASCII characters) are considered\ncryptographically secure.\n\n## See Also\n\n* [commands/`format`](../commands/format.md):\n  Reformat one data-type into another data-type\n* [commands/`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value"

}
