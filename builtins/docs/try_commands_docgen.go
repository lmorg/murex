package docs

func init() {

	Definition["try"] = "# _murex_ Shell Docs\n\n## Command Reference: `try`\n\n> Handles errors inside a block of code\n\n## Description\n\n`try` forces a different execution behavior where a failed process at the end\nof a pipeline will cause the block to terminate regardless of any functions that\nmight follow.\n\nIt's usage is similar to try blocks in other languages (eg Java) but a closer\nfunctional example would be `set -e` in Bash.\n\nTo maintain concurrency within the pipeline, `try` will only check the last\nfunction in any given pipeline (ie series of functions joined via `|`, `->`, or\nsimilar operators). If you need the entire pipeline checked then use `trypipe`.\n\n## Usage\n\n    try { code-block } -> <stdout>\n    \n    <stdin> -> try { -> code-block } -> <stdout>\n\n## Examples\n\n    try {\n        out: \"Hello, World!\" -> grep: \"non-existent string\"\n        out: \"This command will be ignored\"\n    }\n\n## Detail\n\nA failure is determined by:\n\n* Any process that returns a non-zero exit number\n* Any process that returns more output via STDERR than it does via STDOUT\n\nYou can see which run mode your functions are executing under via the `fid-list`\ncommand.\n\n## See Also\n\n* [commands/`catch`](../commands/catch.md):\n  Handles the exception code raised by `try` or `trypipe` \n* [commands/`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [commands/`switch`](../commands/switch.md):\n  Blocks of cascading conditionals\n* [commands/`trypipe`](../commands/trypipe.md):\n  Checks state of each function in a pipeline and exits block on error\n* [commands/evil](../commands/evil.md):\n  \n* [commands/fid-list](../commands/fid-list.md):\n  "

}
