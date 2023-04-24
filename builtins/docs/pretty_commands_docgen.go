package docs

func init() {

	Definition["pretty"] = "# `pretty` - Command Reference\n\n> Prettifies JSON to make it human readable\n\n## Description\n\nTakes JSON from the STDIN and reformats it to make it human readable, then\noutputs that to STDOUT.\n\n## Usage\n\n    <stdin> -> pretty -> <stdout>\n\n## Examples\n\n    » tout: json {\"Array\":[1,2,3],\"Map\":{\"String\": \"Foobar\",\"Number\":123.456}} -> pretty \n    {\n        \"Array\": [\n            1,\n            2,\n            3\n        ],\n        \"Map\": {\n            \"String\": \"Foobar\",\n            \"Number\": 123.456\n        }\n    }\n\n## See Also\n\n* [`format`](../commands/format.md):\n  Reformat one data-type into another data-type\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type"

}
