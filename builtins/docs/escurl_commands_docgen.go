package docs

func init() {

	Definition["escurl"] = "# `escurl`\n\n> Encode or decodes text for the URL\n\n## Description\n\n`escurl` takes input from either STDIN or the parameters and returns the same\ndata, escaped for the URL.\n\n`!eschtml` does the same process in reverse, where it takes URL escaped data\nand returns its unescaped counterpart.\n\n## Usage\n\nEscape\n\n```\n<stdin> -> escurl -> <stdout>\n\nescurl string to escape -> <stdout>\n```\n\nUnescape\n\n```\n<stdin> -> !escurl -> <stdout>\n\n!escurl string to unescape -> <stdout>\n```\n\n## Examples\n\nEscape\n\n```\n» out \"!? <>\" -> escurl\n%21%3F%20%3C%3E%0A \n```\n\nUnescape\n\n```\nout '%21%3F%20%3C%3E%0A' -> !escurl\n!? <>\n```\n\n## Synonyms\n\n* `escurl`\n* `!escurl`\n\n\n## See Also\n\n* [`escape`](../commands/escape.md):\n  Escape or unescape input\n* [`esccli`](../commands/esccli.md):\n  Escapes an array so output is valid shell code\n* [`eschtml`](../commands/eschtml.md):\n  Encode or decodes text for HTML\n* [`get`](../commands/get.md):\n  Makes a standard HTTP request and returns the result as a JSON object\n* [`getfile`](../commands/getfile.md):\n  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.\n* [`post`](../commands/post.md):\n  HTTP POST request with a JSON-parsable return"

}
