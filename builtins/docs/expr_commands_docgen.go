package docs

func init() {

	Definition["expr"] = "# _murex_ Shell Docs\n\n## Command Reference: `expr`\n\n> Expressions: mathematical, string comparisons, logical operators\n\n## Description\n\n\n\n## Usage\n\n    expr: expression -> <stdout>\n\n## Examples\n\nOrder of operations:\n\n    » 3 * (3 + 1)                                                                                                                                                                                                                         \n    12\n    \nJSON array:\n\n    » %[apples oranges grapes]\n    [\n        \"apples\",\n        \"oranges\",\n        \"grapes\"\n    ]\n\n## See Also\n\n* [commands/`=` (arithmetic evaluation)](../commands/equ.md):\n  Evaluate a mathematical function (deprecated)\n* [commands/`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value (deprecated)"

}
