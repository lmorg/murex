# _murex_ Shell Docs

## Command Reference: `murex-update-exe-list`

> Forces _murex_ to rescan $PATH looking for exectables

### Description

On application lauch, _murex_ scans and caches all the executables found in
$PATH on your host. _murex_ then does regular scans there after. However if
you want to force a new scan (for example you've just installed a new
program and you want it to appear in tab completion) then you can run `murex-update-exe-list`.

### Usage

    murex-update-exe-list

### Examples

    » murex-update-exe-list

### See Also

* [commands/`cpuarch`](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [commands/`cpucount`](../commands/cpucount.md):
  Output the number of CPU cores available on your host
* [commands/`os`](../commands/os.md):
  Output the auto-detected OS name