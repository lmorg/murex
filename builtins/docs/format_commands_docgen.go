package docs

func init() {

    Definition["format"] = "# `format`\n\n> Reformat one data-type into another data-type\n\n## Description\n\n`format` takes a data from STDIN and returns that data reformated in another\nspecified data-type\n\n## Usage\n\n```\n<stdin> -> format data-type -> <stdout>\n```\n\n## Examples\n\n```\n» tout json { \"One\": 1, \"Two\": 2, \"Three\": 3 } -> format yaml\nOne: 1\nThree: 3\nTwo: 2\n```\n\n## See Also\n\n* [`Marshal()` (type)](../apis/Marshal.md):\n  Converts structured memory into a structured file format (eg for stdio)\n* [`Unmarshal()` (type)](../apis/Unmarshal.md):\n  Converts a structured file format into structured memory\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type"

}