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

* **Usability improvements** such as in-line spell checking, context sensitive
  hint text that details a commands behavior before you hit return, and
  auto-parsing man pages for auto-completions on commands that don't have auto-completions already defined.
  
* **Smarter handling of errors** and **debugging tools**. For example try/catch
  blocks, line numbers included in error messages, STDOUT highlighted in red
  and script testing and debugging frameworks baked into the language itself.

## Examples

**JSON wrangling:**

<img src="images/murex-open-foreach.png" class="readme">

**Inline spellchecking**

<img src="images/murex-spellchecker.png" class="readme">

**Autocomplete descriptions, process IDs accompanied by process names:**

<img src="images/murex-kill-autocomplete.png" class="readme">

More examples: [/examples](https://github.com/lmorg/murex/tree/master/examples)

## Install instructions

See [INSTALL](INSTALL.md) for details.

## Language Tour

Read the [language tour](docs/GUIDE.quick-start.md) to get started.

## Known bugs / TODO

_murex_ is considered stable, however if you do run into problems then please
raise them on the project's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)
