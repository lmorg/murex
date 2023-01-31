<h2>Table of Contents</h2>

<div id="toc">

- [Overview](#overview)
- [Spellchecker](#spellchecker)
- [Hint Text](#hint-text)
  - [Disable Hint Text](#disable-hint-text)
  - [Hint Text Colour](#hint-text-colour)
  - [Custom Hint Text Statuses](#custom-hint-text-statuses)
- [Autocompletion](#autocompletion)
  - [Tab Completion: `Grid`](#tab-completion-grid)
  - [Tab Completion: `ListView`](#tab-completion-listview)
- [Syntax Completion](#syntax-completion)
- [Syntax Highlighting](#syntax-highlighting)

</div>

## Overview

Aside from _murex_ being carefully designed with scripting in mind, the
interactive shell itself is also built around productivity. To achieve this
we wrote our own readline library. Below is an example of that library in use:

[![asciicast](https://asciinema.org/a/232714.svg)](https://asciinema.org/a/232714)

The above demo includes the following features of _murex_'s bespoke readline
library:

* hint text - blue status text below the prompt (the colour is configurable)
* syntax highlighting (albeit there isn’t much syntax to highlight in the
    example). This can also be turned off if your preference is to have colours
    disabled
* tab-completion in gridded mode (seen when typing `cd`)
* tab-completion in list view (seen when selecting a process name to `kill`
    where the process ID was substituted when selected)
* regex searching through the tab-completion suggestions (seen in both `cd` and
    `kill` - enabled by pressing `[CTRL+f]`)
* line editing using $EDITOR (`vi` in the example - enabled by pressing `[ESC]`
    followed by `[v]`)
* readline’s warning before pasting multiple lines of data into the buffer and
    the preview option that’s available as part of the aforementioned warning
* and VIM keys (enabled by pressing `[ESC]`)

## Spellchecker

_murex_ supports inline spellchecking, where errors are underlined. For example

[![asciicast](https://asciinema.org/a/408024.svg)](https://asciinema.org/a/408024)

This might require some manual steps to enable, please see the [spellcheck user guide](spellcheck.md)
for more details.

## Hint Text

The **hint text** is a (typically) blue status line that appears directly below
your prompt. The idea behind the **hint text** is to provide clues to you as
type instructions into the prompt; but without adding distractions. It is there
to be used if you want it while keeping out of the way when you don't want it.

### Disable Hint Text

It is enabled by default but can be disabled if you prefer a more minimal
prompt:

```
» config: set shell hint-text-enabled false
```

### Hint Text Colour

By default the **hint text** will appear blue. This is also customizable:

```
» config get shell hint-text-formatting
{BLUE}
```

The formatting config takes a string and supports [ANSI constants](ansi.md).

It is also worth noting that if colour is disabled then the **hint text** will
not be coloured even if **hint-text-formatting** includes colour codes:

```
» config: set shell color false
```

(please note that **syntax highlighting** is unaffected by the above config)

### Custom Hint Text Statuses

There is a lot of behavior hardcoded into _murex_ like displaying the full path
to executables and the values of variables. However if there is no status to be
displayed then _murex_ can fallback to a default **hint text** status. This
default is a user defined function. At time of writing this document the author
has the following function defined:

```
» config: get shell hint-text-func
{
    trypipe <!null> {
        git status --porcelain -b -> set gitstatus
        #$gitstatus -> head -n1 -> sed -r 's/^## //;s/\.\.\./ => /' -> set gitbranch
        $gitstatus -> head -n1 -> regexp 's/^## //' -> regexp 's/\.\.\./ => /' -> set gitbranch
        let gitchanges=${ out $gitstatus -> sed 1d -> wc -l }
        !if { $gitchanges } then { ({GREEN}) } else { ({RED}) }
        (Git{BLUE}: $gitbranch ($gitchanges pending). )
    }
    catch {
        ({YELLOW}Git{BLUE}: Not a git repository. )
    }

    if { $SSH_AGENT_PID } then {
        ({GREEN}ssh-agent{BLUE}: $SSH_AGENT_PID. )
    } else {
        ({RED}ssh-agent{BLUE}: No env set. )
    }

    if { pgrep: vpnc } then {
        ({YELLOW}VPN{BLUE}: vpnc is active. )
    }

    if { ps aux -> regexp m/openvpn --errors-to-stderr --log/ } then {
        ({YELLOW}VPN{BLUE}: openvpn is active. )
    }

    trypipe <!null> {
        open: main.tf -> format json -> [ terraform ] -> [ 0 ] -> [ required_version ] -> sed -r 's/\s0\./ /' -> set tfmod
        terraform: version -> head -n1 -> regexp (f/Terraform v0\.([0-9.]+)$) -> set tfver
        if { = tfmod >= tfver } then { ({GREEN}) } else { ({RED}) }
        (Terraform{BLUE}: $tfver; required $tfmod. )
    }

    if { $AWS_SESSION_TOKEN } then {
        set aws_expiration
        set int date=${ date +%s }

        if { os linux } then {
            set int aws_expiration=${ date -d $AWS_SESSION_EXPIRATION +%s }
        } else {
            set int aws_expiration=${ date -j -f "%FT%R:%SZ" $AWS_SESSION_EXPIRATION +%s }
        }

        = (($aws_expiration-$date)/60) -> format int -> set aws_session_time
        if { = aws_session_time < 1 } then { ({RED}) } else { ({GREEN}) }
        (awscon{BLUE}: $AWS_SESSION_NAME => $aws_session_time mins. )
    }
}
```

...which produces a colorized status that looks something like the following:

```
Git: develop => origin/develop [ahead 1] (9 pending). ssh-agent: 34607.
```

## Autocompletion

Autocompletion happen when you press **{TAB}** and will differ slightly depending
on what is defined in `autocomplete` and whether you use the traditional
[POSIX pipe token](../parser/pipe-posix.md), `|`, or the [arrow pipe](../parser/pipe-arrow.md),
`->`.

The `|` token will behave much like any other shell however `->` will offer
suggestions with matching data types (as seen in `runtime --methods`). This is
a way of helping highlight commands that naturally follow after another in a
pipeline. Which is particularly important in _murex_ as it introduces data
types and dozens of new builtins specifically for working with data structures
in an intelligent and readable yet succinct way.

You can add your own commands and functions to _murex_ as methods by defining
them with `method`. For example if we were to add `jq` as a method:

```
method: define jq {
    "Stdin":  "json",
    "Stdout": "@Any"
}
```

### Tab Completion: `Grid`

This is where the completion suggestions are arranged in a grid. This is the
default.

### Tab Completion: `ListView`

This is where the completion suggestions are arranged in a list with a
description along the side.

## Syntax Completion

Like with most IDEs, _murex_ will auto close brackets et al.

[![asciicast](https://asciinema.org/a/408029.svg)](https://asciinema.org/a/408029)

## Syntax Highlighting

Pipelines in the interactive terminal are syntax highlighted. This is similar
to what one expects from an IDE.

Syntax highlighting can be disabled by running:

```
» config: set shell syntax-highlighting off
```