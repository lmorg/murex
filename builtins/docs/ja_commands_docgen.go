package docs

func init() {

	Definition["ja"] = "# `ja` (mkarray) - Command Reference\n\n> A sophisticated yet simply way to build a JSON array\n\n## Description\n\nMurex has a pretty sophisticated builtin for generating JSON arrays.\nIt works a little bit like Bash's `{1..9}` syntax but includes a few\nadditional nifty features.\n\n**Please note that while this builtin is not marked for deprecation, it has\nbeen superseded by the `%[]` tokens.** ([read more](../parser/create-array.md))\n\n## Usage\n\n```\nja: [start..end] -> <stdout>\nja: [start..end.base] -> <stdout>\nja: [start..end,start..end] -> <stdout>\nja: [start..end][start..end] -> <stdout>\n```\n\n## Examples\n\n```\n» ja: [1..5]\n[\n    \"1\",\n    \"2\",\n    \"3\",\n    \"4\",\n    \"5\"\n]\n```\n\n```\n» ja: [Monday..Sunday]\n[\n    \"Monday\",\n    \"Tuesday\",\n    \"Wednesday\",\n    \"Thursday\",\n    \"Friday\",\n    \"Saturday\",\n    \"Sunday\"\n]\n```\n\nPlease note that as per the first example, all arrays generated by `ja` are\narrays of strings - even if you're command is ranging over integers.\n\n## Detail\n\nPlease read the documentation on `a` for a more detailed breakdown on of\n`ja`'s supported features.\n\n## See Also\n\n* [Create array (`%[]`) constructor](../parser/create-array.md):\n  Quickly generate arrays\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`json` ](../types/json.md):\n  JavaScript Object Notation (JSON)\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [`ta` (mkarray)](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type"

}
