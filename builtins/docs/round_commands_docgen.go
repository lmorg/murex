package docs

func init() {

	Definition["round"] = "# `round`  - Command Reference\n\n> Round a number by a user defined precision\n\n## Description\n\n`round` supports a few different levels of precision\n\n## Usage\n\n    round input precision -> <stdout>\n\n## Flags\n\n* `--down`\n    Rounds down to the nearest multiple (not supported when precision is to decimal places)\n* `--up`\n    Rounds up to the nearest multiple (not supported when precision is to decimal places)\n* `-d`\n    alias of `--down\n* `-u`\n    alias of `--up\n\n## See Also\n\n* [`expr`](../commands/expr.md):\n  Expressions: mathematical, string comparisons, logical operators"

}
