package docs

func init() {

	Definition["murex-update-exe-list"] = "# `murex-update-exe-list` - Command Reference\n\n> Forces Murex to rescan $PATH looking for executables\n\n## Description\n\nOn application launch, Murex scans and caches all the executables found in\n$PATH on your host. Murex then does regular scans there after. However if\nyou want to force a new scan (for example you've just installed a new\nprogram and you want it to appear in tab completion) then you can run `murex-update-exe-list`.\n\n## Usage\n\n    murex-update-exe-list\n\n## Examples\n\n    » murex-update-exe-list\n\n## Detail\n\nMurex will automatically update the exe list each time tab completion is\ninvoked for command name completion via the REPL shell.\n\n## See Also\n\n* [`cpuarch`](../commands/cpuarch.md):\n  Output the hosts CPU architecture\n* [`cpucount`](../commands/cpucount.md):\n  Output the number of CPU cores available on your host\n* [`os`](../commands/os.md):\n  Output the auto-detected OS name"

}
