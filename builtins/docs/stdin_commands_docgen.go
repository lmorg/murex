package docs

func init() {

	Definition["<stdin>"] = "# _murex_ Shell Docs\n\n## Command Reference: `<stdin>` \n\n> Read the STDIN belonging to the parent code block\n\n### Description\n\nThis is used inside functions and other code blocks to pass that block's\nSTDIN down a pipeline\n\n### Usage\n\n    <stdin> -> <stdout>\n\n### Examples\n\nWhen writing more complex scripts, you cannot always invoke your read as the\nfirst command in a code block. For example a simple pipeline might be:\n\n    » function: example { -> match: 2 }\n    \nBut this only works if `->` is the very first command. The following would\nfail:\n\n    # Incorrect code\n    function: example {\n        out: \"only match 2\"\n        -> match 2\n    }\n    \nThis is where `<stdin>` comes to our rescue:\n\n    function: example {\n        out: \"only match 2\"\n        <stdin> -> match 2\n    }\n    \nThis could also be written as:\n\n    function: example { out: \"only match 2\"; <stdin> -> match 2 }\n\n### Synonyms\n\n* `<stdin>`\n\n\n### See Also\n\n* [commands/`<>` (read pipe)](../commands/readpipe.md):\n  Reads from a _murex_ named pipe\n* [commands/pipe](../commands/pipe.md):\n  "

}
