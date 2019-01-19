package docs

func init() {

	Definition["jsplit"] = "# _murex_ Language Guide\n\n## Command Reference: `jsplit` \n\n> Splits STDIN into a JSON array based on a regex parameter\n\n### Description\n\n`jsplit` will read from STDIN and split it based on a regex parameter. It outputs a JSON array.\n\n### Usage\n\n    <STDIN> -> jsplit: regex -> <stdout>\n\n### Examples\n\n    » (hello, world) -> jsplit: l+ \n    [\n        \"he\",\n        \"o, wor\",\n        \"d\"\n    ]\n\n### See Also\n\n* [`2darray` ](../commands/2darray.md):\n  Create a 2D JSON array from multiple input sources\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`len` ](../commands/len.md):\n  Outputs the length of an array\n* [`map` ](../commands/map.md):\n  Creates a map from two data sources\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [a](../commands/a.md):\n  \n* [ja](../commands/ja.md):\n  "

}
