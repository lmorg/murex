package docs

func init() {

	Definition["event"] = "# _murex_ Shell Guide\n\n## Command Reference: `event`\n\n> Event driven programming for shell scripts\n\n### Description\n\nCreate or destroy an event interrupt\n\n### Usage\n\n    event: event-type name=interrupt { code block }\n    \n    !event: event-type name\n\n### Examples\n\nCreate an event:\n\n    event: onSecondsElapsed autoquit=60 {\n        out \"You're 60 second timeout has elapsed. Quitting murex\"\n        exit 1\n    }\n    \nDestroy an event:\n\n    !event onSecondsElapsed autoquit\n\n### Detail\n\nThe `interrupt` field in the CLI supports ANSI constants. eg\n\n    event: onKeyPress f1={F1-VT100} {\n        tout: qs HintText=\"Key F1 Pressed\"\n    }\n    \nTo list compiled event types:\n\n    » runtime: --events -> formap k v { out $k }\n    onFileSystemChange\n    onKeyPress\n    onSecondsElapsed\n\n### Synonyms\n\n* `event`\n* `!event`\n\n\n### See Also\n\n* [`function`](../commands/function.md):\n  Define a function block\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [formap](../commands/formap.md):\n  \n* [open](../commands/open.md):\n  "

}
