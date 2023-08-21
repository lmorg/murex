package docs

// This file was generated from [builtins/core/escape/escape_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/escape/escape_doc.yaml).

func init() {

	Definition["escape"] = "# `escape`\n\n> Escape or unescape input\n\n## Description\n\n`escape` takes input from either STDIN or the parameters and returns the same\ndata, escaped.\n\n`!escape` does the same process in reverse, where it takes escaped data and\nreturns its unescaped counterpart.\n\n## Usage\n\nEscape\n\n```\n<stdin> -> escape -> <stdout>\n\nescape string to escape -> <stdout>\n```\n\nUnescape\n\n```\n<stdin> -> !escape -> <stdout>\n\n!escape string to unescape -> <stdout>\n```\n\n## Examples\n\nEscape\n\n```\n» out (multi\n» line\n» string) -> escape\n\"multi\\nline\\nstring\\n\" \n```\n\n## Synonyms\n\n* `escape`\n* `!escape`\n\n\n## See Also\n\n* [`esccli`](../commands/esccli.md):\n  Escapes an array so output is valid shell code\n* [`eschtml`](../commands/eschtml.md):\n  Encode or decodes text for HTML\n* [`escurl`](../commands/escurl.md):\n  Encode or decodes text for the URL\n\n<hr/>\n\nThis document was generated from [builtins/core/escape/escape_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/escape/escape_doc.yaml)."

}
