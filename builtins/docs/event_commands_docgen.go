package docs

func init() {

	Definition["event"] = "# `event` - Command Reference\n\n> Event driven programming for shell scripts\n\n## Description\n\nCreate or destroy an event interrupt,\n\nEach event will have subtilty different behaviour depending on the event itself\ndue to the differing roles of each event system. Therefore it is recommended\nthat you read the docs on each event to understand its behaviour.\n\nHowever while they might differ, the `event` API does try to retain a level of\nexternal consistency. For example each event in defined via `name=interrupt`\nwhere **name** is a user defined handle (like a variable or function would have\na name) and **interrupt** is a system state you wish the event to be fired on.\n\nEach event function will have a payload sent via STDIN which would look a\nlittle like the following:\n\n```\n{\n    \"Name\": \"\",\n    \"Interrupt\": {}\n}\n```\n\n**Name** will always refer to the name you passed when defining the event. And\n**Interrupt** will carry any event specific metadata that might be useful to\nthe event function. Thus the value of **Interrupt** will vary from one event to\nanother.\n\n## Usage\n\n```\nevent: event-type name=interrupt { code block }\n\n!event: event-type name\n```\n\n## Examples\n\nCreate an event:\n\n```\nevent: onSecondsElapsed autoquit=60 {\n    out \"You're 60 second timeout has elapsed. Quitting murex\"\n    exit 1\n}\n```\n\nDestroy an event:\n\n```\n!event onSecondsElapsed autoquit\n```\n\n## Detail\n\n### Supported events\n\n* [`onCommandCompletion`](../events/oncommandcompletion.md):\n  Trigger an event upon a command's completion\n* [`onFileSystemChange`](../events/onfilesystemchange.md):\n  Add a filesystem watch\n* [`onPrompt`](../events/onprompt.md):\n  Events triggered by changes in state of the interactive shell\n* [`onSecondsElapsed`](../events/onsecondselapsed.md):\n  Events triggered by time intervals\n\n### ANSI constants\n\nThe `interrupt` field in the CLI supports ANSI constants. eg\n\n```\nevent: onKeyPress f1={F1-VT100} {\n    tout: qs HintText=\"Key F1 Pressed\"\n}\n```\n\n### Compiled events\n\nTo list compiled event types:\n\n```\n» runtime --events -> formap event ! { out $event }\nonCommandCompletion\nonFileSystemChange\nonKeyPress\nonPrompt\nonSecondsElapsed\n```\n\n## Synonyms\n\n* `event`\n* `!event`\n\n\n## See Also\n\n* [`formap`](../commands/formap.md):\n  Iterate through a map or other collection of data\n* [`function`](../commands/function.md):\n  Define a function block\n* [`open`](../commands/open.md):\n  Open a file with a preferred handler\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex"

}
