package docs

func init() {

	Definition["round"] = "# `round`  - Command Reference\n\n> Round a number by a user defined precision\n\n## Description\n\n`round` supports a few different levels of precision:\n\n### Nearest decimal place\n\nSyntax: `0.12345` (any numbers can be used)\n\nIf a decimal place is supplied then `round` will round your number to however\nmany decimal places you specify. It doesn't matter what digits you include in\nyour precision value as the only thing which is used to drive the result is the\nposition of the decimal point. Thus a precision value of `0.000` would perform\nthe same rounding as `9.999`.\n\nDecimal places are always rounded to the nearest. `--down` and `--up` flags are\nnot supported.\n\n### Nearest integer\n\nSyntax: either `0` or `1` (either behaves the same)\n\nThis will round your value to the nearest whole number. For example `3.33`\nwould be rounded to `3`.\n\nIf `--down` flag is supplied then the remainder is dropped. For example `9.99`\nwould then be rounded to `9` instead of `10`.\n\nIf `--up` flag is is supplied then the the input value would always be rounded\nup to the nearest whole number. For example `3.33` would be rounded to `4`\ninstead of `3`.\n\n### Nearest Multiple\n\nSyntax: `50` (any integer greater than `1)\n\nThis will round your input value to the nearest multiple of your precision.\n\nLike with **nearest integer** (see above), `--down` and `--up` will specify to\nround whether to always round down or up rather than returning the nearest\nmatch in either direction.\n\n## Usage\n\n```\nround value precision -> <stdout>\n```\n\n## Examples\n\n**Rounding to the nearest multiple of `20`:**\n\n```\n» round 15 20\n20\n```\n\n## Flags\n\n* `--down`\n    Rounds down to the nearest multiple (not supported when precision is to decimal places)\n* `--up`\n    Rounds up to the nearest multiple (not supported when precision is to decimal places)\n* `-d`\n    shorthand for `--down`\n* `-u`\n    shorthand for `--up`\n\n## See Also\n\n* [`expr`](../commands/expr.md):\n  Expressions: mathematical, string comparisons, logical operators"

}
