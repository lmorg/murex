package docs

// This file was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).

func init() {

	Definition["true"] = "# `true`\n\n> Returns a `true` value\n\n## Description\n\nReturns a `true` value.\n\n## Usage\n\n```\ntrue -> <stdout>\n```\n\n## Examples\n\nBy default, `true` also outputs the term \"true\":\n\n```\n» true\ntrue\n```\n\nHowever you can suppress that with the silent flag:\n\n```\n» true -s\n```\n\n## Flags\n\n* `-s`\n    silent - don't output the term \"true\"\n\n## See Also\n\n* [`!` (not)](../commands/not.md):\n  Reads the STDIN and exit number from previous process and not's it's condition\n* [`and`](../commands/and.md):\n  Returns `true` or `false` depending on whether multiple conditions are met\n* [`false`](../commands/false.md):\n  Returns a `false` value\n* [`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [`or`](../commands/or.md):\n  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.\n\n<hr/>\n\nThis document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml)."

}
