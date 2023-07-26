package docs

func init() {

	Definition["tabulate"] = "# `tabulate` - Command Reference\n\n> Table transformation tools\n\n## Description\n\n`tabluate` is a swiss army knife for table transforming human readable tables\ninto machine readable data structure.\n\n> Please note that this builtin is still in active development and the default\n> behavior will continue to change and evolve. Any features marked with a flag\n> (see below) will be stable, have numerous tests written against them, and\n> thus safe to use.\n\n## Usage\n\n```\n<stdin> -> tabulate [ flags ] -> <stdout>\n```\n\n## Flags\n\n* `--column-wraps`\n    Boolean, used with --map or --key-value to merge trailing lines if the text wraps within the same column\n* `--help`\n    Boolean, displays a list of flags\n* `--joiner`\n    String, used with --map to concatenate any trailing records in a given field\n* `--key-inc-hint`\n    Boolean, used with --map to split any space or equal delimited hints/examples (eg parsing flags)\n* `--key-value`\n    Boolean, discard any records that don't appear key value pairs (auto-enabled when --map used)\n* `--map`\n    Boolean, return JSON map instead of table\n* `--separator`\n    'String, custom regex pattern for spliting fields (default: `(\\t|\\s[\\s]+)+`)'\n* `--split-comma`\n    Boolean, split first field and duplicate the line if comma found in first field (eg parsing flags in help pages)\n* `--split-space`\n    Boolean, split first field and duplicate the line if white space found in first field (eg parsing flags in help pages)\n\n## Detail\n\n### Dynamic Autocompletion\n\nBecause `tabulate` is designed to parse human readable tables, it is used a lot\nfor dynamically turning command like program help output into JSON maps for\n`autocomplete`'s **DynamicDesc** blocks:\n\n```\nrsync --help -> @[^Options$..--help]re -> tabulate: --map --split-comma --column-wraps --key-inc-hint\n```\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [`formap`](../commands/formap.md):\n  Iterate through a map or other collection of data\n* [`format`](../commands/format.md):\n  Reformat one data-type into another data-type"

}
