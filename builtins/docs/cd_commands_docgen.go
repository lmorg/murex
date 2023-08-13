package docs

func init() {

	Definition["cd"] = "# `cd`\n\n> Change (working) directory\n\n## Description\n\nChanges current working directory.\n\n## Usage\n\n```\ncd [path]\n```\n\n## Examples\n\n**Home directory:**\n\n```\n» cd: ~\n```\n\nIncluding `cd` without a parameter will also change to the current user's home\ndirectory:\n\n```\n» cd\n```\n\n**Absolute path:**\n\n```\n» cd: /etc/\n```\n\n**Relative path:**\n\n```\n» cd: Documents\n» cd: ./Documents\n```\n\n## Detail\n\n`cd` updates an environmental variable, `$PWDHIST` with an array of paths.\nYou can then use that to change to a previous directory.\n\n**View the working directory history:**\n\n```\n» $PWDHIST\n```\n\n**Change to a previous directory:**\n\n```\n» cd $PWDHIST[-1]\n```\n\n### auto-cd\n\nSome people prefer to omit `cd` and just write the path, with their shell\nautomatically changing to that directory if the \"command\" is just a directory.\nIn Murex you can enable this behaviour by turning on \"auto-cd\":\n\n```\nconfig: set shell auto-cd true\n```\n\n## See Also\n\n* [Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by Murex\n* [`source` ](../commands/source.md):\n  Import Murex code from another file of code block"

}
