package docs

func init() {

	Definition["map"] = "# _murex_ Language Guide\n\n## Command Reference: `map` \n\n> Creates a map from two data sources\n\n### Description\n\nThis takes two parameters - which are code blocks - and combines them to output a key/value map in JSON.\n\nThe first block is the key and the second is the value.\n\n### Usage\n\n    map { code-block } { code-block } -> <stdout>\n\n### Examples\n\n    » { tout: json ([\"key 1\", \"key 2\", \"key 3\"]) } { tout: json ([\"value 1\", \"value 2\", \"value 3\"]) } \n    {\n        \"key 1\": \"value 1\",\n        \"key 2\": \"value 2\",\n        \"key 3\": \"value 3\"\n    }\n\n### See Also\n\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`len` ](../commands/len.md):\n  Outputs the length of an array\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [a](../commands/a.md):\n  \n* [ja](../commands/ja.md):\n  "

}
