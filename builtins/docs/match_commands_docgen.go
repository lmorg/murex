package docs

func init() {

	Definition["match"] = "# _murex_ Shell Docs\n\n## Command Reference: `match`\n\n> Match an exact value in an array\n\n### Description\n\n`match` takes input from STDIN and returns any array items / lines which\ncontain an exact match of the parameters supplied.\n\nWhen multiple parameters are supplied they are concatenated into the search\nstring and white space delimited. eg all three of the below are the same:\n    match \"a b c\"\n    match a\\sb\\sc\n    match a b c\n    match a    b    c\n\n### Usage\n\n    <stdin> -> match search string -> <stdout>\n\n### Examples\n\n    » ja: [Monday..Friday] -> match Wed\n    [\n        \"Wednesday\"\n    ]\n\n### Detail\n\n`match` is data-type aware so will work against lists or arrays of whichever\n_murex_ data-type is passed to it via STDIN and return the output in the\nsame data-type.\n\n### See Also\n\n* [commands/`2darray` ](../commands/2darray.md):\n  Create a 2D JSON array from multiple input sources\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`append`](../commands/append.md):\n  Add data to the end of an array\n* [commands/`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [commands/`len` ](../commands/len.md):\n  Outputs the length of an array\n* [commands/`map` ](../commands/map.md):\n  Creates a map from two data sources\n* [commands/`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [commands/`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [commands/`pretty`](../commands/pretty.md):\n  Prettifies JSON to make it human readable\n* [commands/`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings\n* [commands/`ta`](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type\n* [commands/prefix](../commands/prefix.md):\n  \n* [commands/suffix](../commands/suffix.md):\n  "

}
