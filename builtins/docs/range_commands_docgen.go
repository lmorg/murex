package docs

func init() {

	Definition["@["] = "# `[` (range)  - Command Reference\n\n> Outputs a ranged subset of data from STDIN\n\n## Description\n\nThis will read from STDIN and output a subset of data in a defined range.\n\nThe range can be defined as a number of different range types - such as the\ncontent of the array or it's index / row number. You can also omit either\nthe start or the end of the search criteria to cover all items before or\nafter the remaining search criteria.\n\n**Please note that `@[` syntax has been deprecated in favour of `[` syntax\ninstead**\n\n## Usage\n\n    <stdin> -> [start..end]flags -> <stdout>\n\n## Examples\n\n**Range over all months after March:**\n\n    » a: [January..December] -> [March..]se\n    April\n    May\n    June\n    July\n    August\n    September\n    October\n    November\n    December\n    \n**Range from the 6th to the 10th month:**\n\nBy default, ranges start from one, `1`\n\n    » a: [January..December] -> [5..9]\n    May\n    June\n    July\n    August\n    September\n    \n**Return the first 3 months:**\n\nThis usage is similar to `head -n3`\n\n    » a: [January..December] -> [..3]\n    October\n    November\n    December\n    \n**Return the last 3 months:**\n\nThis usage is similar to `tail -n3`\n\n    » a: [January..December] -> [-3..]\n    October\n    November\n    December\n\n## Flags\n\n* `8`\n    handles backspace characters (char 8) instead of treating it like a printable character\n* `b`\n    removes blank (empty) lines from source\n* `e`\n    exclude the start and end search criteria from the range\n* `n`\n    numeric offset (indexed from 0)\n* `r`\n    regexp match\n* `s`\n    exact string match\n* `t`\n    trims whitespace from source\n\n## Synonyms\n\n* `@[`\n\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`alter`](../commands/alter.md):\n  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array"

}
