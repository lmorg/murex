package docs

func init() {

	Definition["len"] = "# _murex_ Language Guide\n\n## Command Reference: `len` \n\n> Outputs the length of an array\n\n### Description\n\nThis will read an array from STDIN and outputs the length for that array\n\n### Usage\n\n    <STDIN> -> len -> <stdout>\n\n### Examples\n\n    » tout: json ([\"a\", \"b\", \"c\"]) -> len \n    3\n\n### See Also\n\n* [`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`a`](../commands/a.md):\n  A sophisticated yet simply way to build an array or list\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`map` ](../commands/map.md):\n  Creates a map from two data sources\n* [`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [mtac](../commands/mtac.md):\n  "

}
