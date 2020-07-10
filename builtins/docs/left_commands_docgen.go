package docs

func init() {

	Definition["left"] = "# _murex_ Shell Docs\n\n## Command Reference: `left`\n\n> Left substring a list\n\n## Description\n\nTakes a list from STDIN and returns a left substring of that same list.\n\nOne parameter is required and that is the number of characters to return. If\nthe parameter is a negative then `left` counts from the right.\n\n## Usage\n\n    <stdin> -> left int -> <stdout>\n\n## Examples\n\nCount from the left\n\n    » ja: [Monday..Wednesday] -> left 2\n    [\n        \"Mo\",\n        \"Tu\",\n        \"We\"\n    ]\n    \nCount from the right\n\n    » ja: [Monday..Wednesday] -> left -3\n    [\n        \"Mon\",\n        \"Tues\",\n        \"Wednes\"\n    ]\n\n## Detail\n\nSupported data types can queried via `runtime`\n\n    runtime: --marshallers\n    runtime: --unmarshallers\n\n## See Also\n\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`right`](../commands/right.md):\n  Right substring a list\n* [api/marshaldata](../api/marshaldata.md):\n  \n* [commands/prefix](../commands/prefix.md):\n  \n* [commands/suffix](../commands/suffix.md):\n  \n* [api/unmarshaldata](../api/unmarshaldata.md):\n  "

}
