package docs

func init() {

	Definition["test"] = "# _murex_ Shell Docs\n\n## Command Reference: `test`\n\n> _murex_'s test framework - define tests, run tests and debug shell scripts\n\n## Description\n\n`test` is used to define tests, run tests and debug _murex_ shell scripts.\n\n## Usage\n\nDefine an inlined test\n\n    test: define test-name { json-properties }\n    \nDefine a state report\n\n    test: state name { code block }\n    \nDefine a unit test\n\n    test: unit function|private|open|event test-name { json-properties }\n    \nEnable or disable boolean test states (more options available in `config`)\n\n    test: config [ enable|!enable ] [ verbose|!verbose ] [ auto-report|!auto-report ]\n    \nDisable test mode\n\n    !test\n    \nExecute a function with testing enabled\n\n    test: run { code-block }\n    \nExecute unit test(s)\n\n    test: run-unit package[/module[/test-name]|*\n\n## Synonyms\n\n* `test`\n* `!test`\n\n\n## See Also\n\n* [commands/`<>` (murex named pipe)](../commands/namedpipe.md):\n  Reads from a _murex_ named pipe\n* [commands/`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [parser/namedpipe](../parser/namedpipe.md):\n  "

}
