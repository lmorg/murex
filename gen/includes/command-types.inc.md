There are a few different types of commands:

### alias

This will be represented in `which` and `type` by the term **alias** and, when
stdout is a TTY, `which` will follow the alias to print what command the alias
points to.

### function

This is a Murex function (defined via `function`) and will be represented in
`which` and `type` by the term **function**.

### private

This is a private function (defined via `private`) and will be represented in
`which` and `type` by the term **private**.

### builtin

This is a shell builtin, like `out` and `exit`. It will be represented in
`which` and `type` by the term **builtin**.

### external executable

This is any other external command, such as `systemctl` and `python`. This
will be represented in `which` by the path to the executable. When stdout is a
TTY, `which` will also print the destination path of any symlinks too.

In `type`, it is represented by the term **executable**.