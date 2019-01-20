package docs

func init() {

	Definition["append"] = "# _murex_ Language Guide\n\n## Command Reference: `append`\n\n> Add data to the end of an array\n\n### Description\n\n`append` data to the end of an array.\n\n### Usage\n\n    <stdin> -> append: value -> <stdout>\n\n### Examples\n\n    » a: [Monday..Sunday] -> append: Funday\n    Monday\n    Tuesday\n    Wednesday\n    Thursday\n    Friday\n    Saturday\n    Sunday\n    Funday\n\n### Detail\n\nIt's worth noting that `prepend` and `append` are not data type aware. So \nany integers in data type aware structures will be converted into strings:\n\n    » tout: json [1,2,3] -> append: new \n    [\n        \"1\",\n        \"2\",\n        \"3\",\n        \"new\"\n    ]\n\n### See Also\n\n* [`len` ](../commands/len.md):\n  Outputs the length of an array\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [a](../commands/a.md):\n  \n* [cast](../commands/cast.md):\n  \n* [index](../commands/index.md):\n  \n* [ja](../commands/ja.md):\n  \n* [range](../commands/range.md):\n  \n* [square-bracket-open](../commands/square-bracket-open.md):\n  \n* [update](../commands/update.md):\n  "

}
