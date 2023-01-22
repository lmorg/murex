[![Version](version.svg)](DOWNLOAD.md)
[![CodeBuild](https://codebuild.eu-west-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoib3cxVnoyZUtBZU5wN1VUYUtKQTJUVmtmMHBJcUJXSUFWMXEyc2d3WWJldUdPTHh4QWQ1eFNRendpOUJHVnZ5UXBpMXpFVkVSb3k2UUhKL2xCY2JhVnhJPSIsIml2UGFyYW1ldGVyU3BlYyI6Im9QZ2dPS3ozdWFyWHIvbm8iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](DOWNLOAD.md)
[![CircleCI](https://circleci.com/gh/lmorg/murex/tree/master.svg?style=svg)](https://circleci.com/gh/lmorg/murex/tree/master)
[![codecov](https://codecov.io/gh/lmorg/murex/branch/master/graph/badge.svg)](https://codecov.io/gh/lmorg/murex)

## About _murex_

_murex_ is a shell, like bash / zsh / fish / etc. It follows a similar syntax
to POSIX shells like Bash however supports more advanced features than you'd
typically expect from a $SHELL.

A non-exhaustive list features would include:

* Support for **additional type information in pipelines**, which can be used
  for complex data formats like JSON or tables. Meaning all of your existing
  UNIX tools to work more intelligently and without any additional configuration.

  JSON wrangling:
  </br>![json-example](images/murex-open-foreach.png)
  
  Running SQL queries on the output of standard UNIX tools:
  </br>![tabulated-data-example](images/murex-ps-select.png)

* **Usability improvements** such as in-line spell checking, context sensitive
  hint text that details a commands behavior before you hit return, and
  auto-parsing man pages for auto-completions on commands that don't have auto-completions already defined.

  Inline spellchecking:
  </br>![spellchecking](images/murex-spellchecker.png)

  Autcomplete descriptions, process IDs accompanied by process names:
  </br>![smarter-autocomplete](images/murex-kill-autocomplete.png)
  
* **Smarter handling of errors** and **debugging tools**. For example try/catch
  blocks, line numbers included in error messages, STDOUT highlighted in red
  and script testing and debugging frameworks baked into the language itself.

## More Examples!

### Getting indexes from tabulated data:

```
ps aux | [PID %CPU COMMAND] | head -n5
```

Outputs:
```
PID     %CPU    COMMAND
77045   127.5   /usr/sbin/netbiosd
85046   14.9    /Applications/iTerm.app/Contents/MacOS/iTerm2
371     3.7     /System/Library/PrivateFrameworks/SkyLight.framework/Resources/WindowServer
4302    3.3     /Applications/Firefox.app/Contents/MacOS/firefox
```

### Arrays used as parameters:

```
fruit = %[apples oranges bananas]
out: "I have the following fruit in my fruit bowl:" @fruit ","
out: "But I mostly love $fruit[1]."
```

Outputs:
```
I have the following fruit in my fruit bowl: apples oranges bananas,
But I mostly love oranges.
```

### Iteration:

```
%[ A[1..3],Letter ] | foreach page_size {
    if { $page_size == 'Letter' } then {
        out: "$page_size is loaded"
    } else {
        out: "$page_size is unsupported"
    }
}
```

Outputs:
```
A3 is unsupported
A4 is unsupported
A5 is unsupported
Letter is loaded
```

## Install instructions

See [INSTALL](INSTALL.md) for details.

## Known bugs / TODO

_murex_ is considered stable, however if you do run into problems then please
raise them on the project's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)
