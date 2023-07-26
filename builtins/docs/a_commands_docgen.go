package docs

func init() {

	Definition["a"] = "# `a` (mkarray) - Command Reference\n\n> A sophisticated yet simple way to build an array or list\n\n## Description\n\nPronounced \"make array\", like `mkdir` (etc), Murex has a pretty sophisticated\nbuiltin for generating arrays. Think like bash's `{1..9}` syntax:\n\n```\na: [1..9]\n```\n\nExcept Murex also supports other sets of ranges like dates, days of the week,\nand alternative number bases.\n\n## Usage\n\n```\na: [start..end] -> <stdout>\na: [start..end,start..end] -> <stdout>\na: [start..end][start..end] -> <stdout>\n```\n\nAll usages also work with `ja` and `ta` as well, eg:\n\n```\nja: [start..end] -> <stdout>\nta: data-type [start..end] -> <stdout>\n```\n\nYou can also inline arrays with the `%[]` syntax, eg:\n\n```\n%[start..end]\n```\n\n## Examples\n\n```\n» a: [1..3]\n1\n2\n3\n\n» a: [3..1]\n3\n2\n1\n\n» a: [01..03]\n01\n02\n03\n```\n\n## Detail\n\n### Advanced Array Syntax\n\nThe syntax for `a` is a comma separated list of parameters with expansions\nstored in square brackets. You can have an expansion embedded inside a\nparameter or as it's own parameter. Expansions can also have multiple\nparameters.\n\n```\n» a: 01,02,03,05,06,07\n01\n02\n03\n05\n06\n07\n```\n\n```\n» a: 0[1..3],0[5..7]\n01\n02\n03\n05\n06\n07\n```\n\n```\n» a: 0[1..3,5..7]\n01\n02\n03\n05\n06\n07\n```\n\n```\n» a: b[o,i]b\nbob\nbib\n```\n\nYou can also have multiple expansion blocks in a single parameter:\n\n```\n» a: a[1..3]b[5..7]\na1b5\na1b6\na1b7\na2b5\na2b6\na2b7\na3b5\na3b6\na3b7\n```\n\n`a` will cycle through each iteration of the last expansion, moving itself\nbackwards through the string; behaving like an normal counter.\n\n### Creating JSON arrays with `ja`\n\nAs you can see from the previous examples, `a` returns the array as a\nlist of strings. This is so you can stream excessively long arrays, for\nexample every IPv4 address: `a: [0..254].[0..254].[0..254].[0..254]`\n(this kind of array expansion would hang bash).\n\nHowever if you needed a JSON string then you can use all the same syntax\nas `a` but forgo the streaming capability:\n\n```\n» ja: [Monday..Sunday]\n[\n    \"Monday\",\n    \"Tuesday\",\n    \"Wednesday\",\n    \"Thursday\",\n    \"Friday\",\n    \"Saturday\",\n    \"Sunday\"\n]\n```\n\nThis is particularly useful if you are adding formatting that might break\nunder `a`'s formatting (which uses the `str` data type).\n\n### Smart arrays\n\nMurex supports a number of different formats that can be used to generate\narrays. For more details on these please refer to the documents for each format\n\n* [Calendar Date Ranges](../mkarray/date.md):\n  Create arrays of dates\n* [Character arrays](../mkarray/character.md):\n  Making character arrays (a to z)\n* [Decimal Ranges](../mkarray/decimal.md):\n  Create arrays of decimal integers\n* [Non-Decimal Ranges](../mkarray/non-decimal.md):\n  Create arrays of integers from non-decimal number bases\n* [Special Ranges](../mkarray/special.md):\n  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)\n\n## See Also\n\n* [Create array (`%[]`) constructor](../parser/create-array.md):\n  Quickly generate arrays\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [`str` (string) ](../types/str.md):\n  string (primitive)\n* [`ta` (mkarray)](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type"

}
