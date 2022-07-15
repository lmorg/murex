package docs

func init() {

	Definition["cd"] = "# _murex_ Shell Docs\n\n## Command Reference: `cd`\n\n> Change (working) directory\n\n## Description\n\nChanges current working directory.\n\n## Usage\n\n    cd path\n\n## Examples\n\n    # Home directory\n    » cd: ~ \n    \n    # Absolute path\n    » cd: /etc/\n    \n    # Relative path\n    » cd: Documents\n    » cd: ./Documents\n\n## Detail\n\n`cd` updates an environmental variable, `$PWDHIST` with an array of paths.\nYou can then use that to change to a previous directory\n\n    # View the working directory history\n    » $PWDHIST\n    \n    # Change to a previous directory\n    » cd $PWDHIST[0]\n\n## See Also\n\n* [user-guide/Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by _murex_\n* [commands/`source` ](../commands/source.md):\n  Import _murex_ code from another file of code block"

}
