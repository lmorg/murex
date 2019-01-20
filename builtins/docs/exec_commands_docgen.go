package docs

func init() {

	Definition["exec"] = "# _murex_ Language Guide\n\n## Command Reference: `exec`\n\n> Runs an executable\n\n### Description\n\nWith _murex_, like most other shells, you launch a process by calling the\nname of that executable directly. While this is suitable 99% of the time,\noccasionally you might run into an edge case where that wouldn't work. The\nprimary reason being if you needed to launch a process from a variable, eg\n\n    » set exe=uname\n    » $exe\n    uname\n    \nAs you can see here, _murex_'s behavior here is to output the contents of\nthe variable rather then executing the contents of the variable. This is\ndone for safety reasons, however if you wanted to override that behavior\nthen you could prefix the variable with exec:\n\n    » set exe=uname\n    » exec $exe\n    Linux\n\n### Usage\n\n    <stdin> -> exec\n    <stdin> -> exec -> <stdout>\n               exec -> <stdout>\n\n### Examples\n\n    » exec printf \"Hello, world!\"\n    Hello, world!\n\n### See Also\n\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [eval](../commands/eval.md):\n  \n* [let](../commands/let.md):\n  "

}
