package docs

func init() {

	Definition["cd"] = "# `cd` - Command Reference\n\n> Change (working) directory\n\n## Description\n\nChanges current working directory.\n\n## Usage\n\n```\ncd path\n```\n\n## Examples\n\n```\n# Home directory\n» cd: ~ \n\n# Absolute path\n» cd: /etc/\n\n# Relative path\n» cd: Documents\n» cd: ./Documents\n```\n\n## Detail\n\n`cd` updates an environmental variable, `$PWDHIST` with an array of paths.\nYou can then use that to change to a previous directory\n\n```\n# View the working directory history\n» $PWDHIST\n\n# Change to a previous directory\n» cd $PWDHIST[0]\n```\n\n### auto-cd\n\nSome people prefer to omit `cd` and just write the path, with their shell\nautomatically changing to that directory if the \"command\" is just a directory.\nIn Murex you can enable this behaviour by turning on \"auto-cd\":\n\n```\nconfig: set shell auto-cd true\n```\n\n## See Also\n\n* [Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by Murex\n* [`source` ](../commands/source.md):\n  Import Murex code from another file of code block"

}
