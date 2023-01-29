package docs

func init() {

	Definition["man-get-flags"] = "# _murex_ Shell Docs\n\n## Command Reference: `man-get-flags` \n\n> Parses man page files for command line flags \n\n## Description\n\nSometimes you might want to programmatically search `man` pages for any\nsupported flag. Particularly if you're writing a dynamic autocompletion.\n`man-get-flags` does this and returns a JSON document.\n\nYou can either pipe a man page to `man-get-flags`, or pass the name of\nthe command as a parameter.\n\n## Usage\n\n    <stdin> -> man-get-flags -> <stdout>\n    \n    man-get-flags command -> <stdout>\n\n## See Also\n\n* [`murex-docs`](../commands/murex-docs.md):\n  Displays the man pages for _murex_ builtins"

}
