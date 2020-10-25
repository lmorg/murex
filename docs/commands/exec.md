# _murex_ Shell Docs

## Command Reference: `exec`

> Runs an executable

## Description

With _murex_, like most other shells, you launch a process by calling the
name of that executable directly. While this is suitable 99% of the time,
occasionally you might run into an edge case where that wouldn't work. The
primary reason being if you needed to launch a process from a variable, eg

    » set exe=uname
    » $exe
    uname
    
As you can see here, _murex_'s behavior here is to output the contents of
the variable rather then executing the contents of the variable. This is
done for safety reasons, however if you wanted to override that behavior
then you could prefix the variable with exec:

    » set exe=uname
    » exec $exe
    Linux

## Usage

    <stdin> -> exec
    <stdin> -> exec -> <stdout>
               exec -> <stdout>

## Examples

    » exec printf "Hello, world!"
    Hello, world!

## Detail

If any command doesn't exist as a builtin, function nor alias, then _murex_
will default to forking out to any command with this name (subject to an
absolute path or the order of precedence in `$PATH`). Any forked process will
show up in both the operating systems process viewer (eg `ps`) but also
_murex_'s own process viewer, `fid-list`. However inside `fid-list` you will
notice that all external processes are listed as `exec` with the process name
as part of `exec`'s parameters. That is because that is literally how _murex_
handles any programs that aren't native to _murex_.

## See Also

* [commands/`=` (arithmetic evaluation)](../commands/equ.md):
  Evaluate a mathematical function
* [commands/`bg`](../commands/bg.md):
  Run processes in the background
* [commands/`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [commands/`fid-kill`](../commands/fid-kill.md):
  Terminate a running _murex_ function
* [commands/`fid-killall`](../commands/fid-killall.md):
  Terminate _ALL_ running _murex_ functions
* [commands/`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable
* [commands/`murex-update-exe-list`](../commands/murex-update-exe-list.md):
  Forces _murex_ to rescan $PATH looking for exectables
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value
* [commands/bexists](../commands/bexists.md):
  
* [commands/builtins](../commands/builtins.md):
  
* [commands/jobs](../commands/jobs.md):
  