# `EVENT_RETURN` (json)

> Return values for events

## Description

Some events support return parameters outside of your typical stdout and stderr
streams. `$EVENT_RETURN` allows you to modify those parameters.

## Examples

```
event onPreview example=exec {
    -> set event
    out "Preview event for $(event.Interrupt.PreviewItem)"
    
    $EVENT_RETURN.CacheTTL = 0 # don't cache this response.
}
```

## Detail

`$EVENT_RETURN` will support different values for different events. Please read
the respective event document for details on using this variable.

## Other Reserved Variables

* [Numeric (str)](../variables/numeric.md):
  Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above
* [`$.`, Meta Values (json)](../variables/meta-values.md):
  State information for iteration blocks
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
* [`COLUMNS` (int)](../variables/lines.md):
  Character height of terminal
* [`EVENT_RETURN` (json)](../variables/event_return.md):
  Return values for events
* [`HOME` (path)](../variables/home.md):
  Return the home directory for the current session user
* [`HOSTNAME` (str)](../variables/hostname.md):
  Hostname of the current machine
* [`LOGNAME` (str)](../variables/logname.md):
  Username for the current session (historic)
* [`MUREX_ARGV` (json)](../variables/murex_argv.md):
  Array of the command name and parameters passed to the current shell
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`OLDPWD` (path)](../variables/oldpwd.md):
  Return the home directory for the current session user
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`PWDHIST` (json)](../variables/pwdhist.md):
  History of each change to the sessions working directory
* [`PWD` (path)](../variables/pwd.md):
  Current working directory
* [`RANDOM` (int)](../variables/random.md):
  Return a random 32-bit integer (historical)
* [`SELF` (json)](../variables/self.md):
  Meta information about the running scope.
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell
* [`TMPDIR` (path)](../variables/tmpdir.md):
  Return the temporary directory
* [`USER` (str)](../variables/user.md):
  Username for the current session

## See Also

* [Murex Event Subsystem (`event`)](../commands/event.md):
  Event driven programming for shell scripts
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`onKeyPress`](../events/onkeypress.md):
  Custom definable key bindings and macros
* [`onPreview`](../events/onpreview.md):
  Full screen previews for files and command documentation

<hr/>

This document was generated from [gen/variables/EVENT_RETURN_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/EVENT_RETURN_doc.yaml).