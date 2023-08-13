package docs

func init() {

	Definition["cast"] = "# `cast`\n\n> Alters the data type of the previous function without altering it's output\n\n## Description\n\n`cast` works a little like when you case variables in lower level languages\nwhere the value of the variable is unchanged. In Murex the contents in\nthe pipeline are preserved however the reported data type is altered.\n\n## Usage\n\n```\n<stdin> -> cast data-type -> <stdout>\n```\n\n## Examples\n\n```\n» out: {\"Array\":[1,2,3],\"Map\":{\"String\": \"Foobar\",\"Number\":123.456}} -> cast json\n{\"Array\":[1,2,3],\"Map\":{\"String\": \"Foobar\",\"Number\":123.456}}\n```\n\n## Detail\n\nIf you want to reformat the STDIN into the new data type then use `format`\ninstead.\n\n## See Also\n\n* [`format`](../commands/format.md):\n  Reformat one data-type into another data-type\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type"

}
