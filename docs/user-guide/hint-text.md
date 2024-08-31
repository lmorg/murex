# Hint Text

> A status bar for your shell

<h2>Table of Contents</h2>

<div id="toc">

- [Overview](#overview)
- [Configuring Hint Text Colour](#configuring-hint-text-colour)
- [Custom Hint Text Statuses](#custom-hint-text-statuses)
  - [Starship Example](#starship-example)
- [Disabling Hint Text](#disabling-hint-text)

</div>



## Overview

The **hint text** is a (typically) blue status line that appears directly below
your prompt. The idea behind the **hint text** is to provide clues to you as
type instructions into the prompt; but without adding distractions. It is there
to be used if you want it while keeping out of the way when you don't want it.




## Configuring Hint Text Colour

By default the **hint text** will appear blue. This is also customizable:

```
» config get shell hint-text-formatting
{BLUE}
```

The formatting config takes a string and supports [ANSI constants](ansi.md).

It is also worth noting that if colour is disabled then the **hint text** will
not be coloured even if **hint-text-formatting** includes colour codes:

```
» config set shell color false
```

(please note that **syntax highlighting** is unaffected by the above config)

## Custom Hint Text Statuses

There is a lot of behavior hardcoded into Murex like displaying the full path
to executables and the values of variables. However if there is no status to be
displayed then Murex can fallback to a default **hint text** status. This
default is a user defined function. At time of writing this document the author
has the following function defined:

```
config set shell hint-text-func {
    trypipe <!null> {
        git status --porcelain -b -> set gitstatus
        $gitstatus -> head -n1 -> regexp 's/^## //' -> regexp 's/\.\.\./ => /'
    }
    catch {
        out "Not a git repository."
    }
}
```

...which produces a colorized status that looks something like the following:

```
develop => origin/develop
```



## Disabling Hint Text

It is enabled by default but can be disabled if you prefer a more minimal
prompt:

```
» config set shell hint-text-enabled false
```

## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings

<hr/>

This document was generated from [gen/user-guide/hint-text_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/hint-text_doc.yaml).