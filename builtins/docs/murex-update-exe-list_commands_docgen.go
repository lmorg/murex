package docs

func init() {

	Definition["murex-update-exe-list"] = "# _murex_ Shell Docs\n\n## Command Reference: `murex-update-exe-list`\n\n> Forces _murex_ to rescan $PATH looking for exectables\n\n## Description\n\nOn application lauch, _murex_ scans and caches all the executables found in\n$PATH on your host. _murex_ then does regular scans there after. However if\nyou want to force a new scan (for example you've just installed a new\nprogram and you want it to appear in tab completion) then you can run `murex-update-exe-list`.\n\n## Usage\n\n    murex-update-exe-list\n\n## Examples\n\n    » murex-update-exe-list\n\n## See Also\n\n* [`cpuarch`](../commands/cpuarch.md):\n  Output the hosts CPU architecture\n* [`cpucount`](../commands/cpucount.md):\n  Output the number of CPU cores available on your host\n* [`os`](../commands/os.md):\n  Output the auto-detected OS name"

}
