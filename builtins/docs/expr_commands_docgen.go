package docs

func init() {

	Definition["expr"] = "# `expr`\n\n> Expressions: mathematical, string comparisons, logical operators\n\n## Description\n\n\n\n## Usage\n\n```\nexpr expression -> <stdout>\n```\n\n## Examples\n\nOrder of operations:\n\n```\n» 3 * (3 + 1)                                                                                                                                                                                                                         \n12\n```\n\nJSON array:\n\n```\n» %[apples oranges grapes]\n[\n    \"apples\",\n    \"oranges\",\n    \"grapes\"\n]\n```\n\n## See Also\n\n* [`=` (arithmetic evaluation)](../commands/equ.md):\n  Evaluate a mathematical function (deprecated)\n* [`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value"

}
