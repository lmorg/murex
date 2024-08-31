# Spellcheck

> How to enable inline spellchecking

## Description

Murex supports inline spellchecking, where errors are underlined. For example

[![asciicast](https://asciinema.org/a/408024.svg)](https://asciinema.org/a/408024)

However to use this there needs to be a few satisfied prerequisites, not all of
which will be enabled by default:

## CLI Spellchecker (3rd Party Software)

A CLI spellchecker needs to be installed. The recommendation is `aspell`. This
might already be installed by default with your OS or has been included as a
dependency with another application. You can check if `aspell` is installed by
running the following:

```
which aspell
```

If that returns no data, then you will need to install `aspell` yourself.
Please consult your OS docs for how to install software.

For help debugging issues with `aspell`, please see the last section in this
document.

## Murex Config

### ANSI Escape Sequences

ANSI escape sequences need to be enabled (which they are by default). This
option is found in `config` under **shell**, **color**.

```
config: set shell color true
```

### Spellcheck Enable

Spellcheck needs to be enabled. This option can be found in `config` under
**shell**, **spellcheck-enabled**.

To enable this run:

```
config set shell spellcheck-enabled true
```

...or add the above line to your Murex profile, `~/.murex_profile` to make
the change persistent.

> Please note that this option will automatically be enabled if `aspell` is
> installed.

### Spellcheck Murex Code

This shouldn't need tweaking if you're running `aspell` but other spellcheckers
will require updated code. The default will look something like this:

```
» config get shell spellcheck-func
{ -> aspell list }
```

The default should be good enough for most people but should you want to run an
alternative spellchecker then follow the instructions in the next section:

## How To Write Your Own `spellcheck-func`

You might legitimately want to run a different spellchecker and if so you'll
need to write your own **spellcheck-func**. Fortunately this is simple:

The function reads the command line from stdin, if the spellchecker reads lines
from parameters rather than stdin you'll need to write something equivalent to
the following

```
{
    # This is a theoretical example. It will not work generically.
    -> set line
    newspellchecker --check "$line"
}
```

The output of the function must me an array containing the misspelt words only.
That array can be JSON just as long as you have set stdout's data type to
`json`. Similarly, other supported Murex data types can be used too. However
in general you might just want to go with a misspelling per line as it's pretty
POSIX friendly and thus most spellcheckers are likely to support it. eg

```
» out "a list of misspelt words: qwert fuubar madeupword" -> aspell list
qwert
fuubar
madeupword
```

## User Dictionary

Murex has it's own user dictionary, which is held as a JSON array:

```
» config: get shell spellcheck-user-dictionary
["murex"]
```

You can add words to a user dictionary via:

```
» config eval shell spellcheck-user-dictionary { -> append "myword" }
```

or

```
» config eval shell spellcheck-user-dictionary { -> alter --merge / (["word1", "word2", "word3"]) }
```

> Don't forget to record these in your Murex profile, `~/.murex_profile` to
> make the changes persistent.

### Ignored By Default

Sometimes commands are not valid words in ones native language. Thus any words
that fall into the following categories are ignored by default:

* words that are also the names of commands found in `$PATH`
* words that are the names of Murex functions (defined via `function`)
* words that are builtins (eg `config` and `jsplit`)
* any global aliases
* also any words that are also the names of global variables

## Common Problems With `aspell`

### `Error: No word lists can be found for the language "en_NZ".`

The `en_NZ` portion of the error will differ depending on your language.

If this error arises then it means `aspell` is installed but it doesn't have
the dictionary for your language. This is an easy fix in most OSs. For example
in Ubuntu:

```
apt install aspell-en
```

(you may need to change `-en` with your specific language code)

Please consult your operating systems documentation for how to install software
and what the package names are for `aspell` and its corresponding dictionaries.

## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Alter Data Structure (`alter`)](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Code Block Parsing](../user-guide/code-block.md):
  Overview of how code blocks are parsed
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Profile Files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`{ Curly Brace }`](../parser/curly-brace.md):
  Initiates or terminates a code block

## Other Integrations

* [ChatGPT](../integrations/chatgpt.md):
  How to enable ChatGPT hints
* [Cheat.sh](../integrations/cheatsh.md):
  Cheatsheets provided by cheat.sh
* [Kitty Integrations](../integrations/kitty.md):
  Get more out of Kitty terminal emulator
* [Makefiles / `make`](../integrations/make.md):
  `make` integrations
* [Man Pages (POSIX)](../integrations/man-pages.md):
  Linux/UNIX `man` page integrations
* [Spellcheck](../integrations/spellcheck.md):
  How to enable inline spellchecking
* [Terminology Integrations](../integrations/terminology.md):
  Get more out of Terminology terminal emulator
* [`direnv` Integrations](../integrations/direnv.md):
  Directory specific environmental variables
* [`yarn` Integrations](../integrations/yarn.md):
  Working with `yarn` and `package.json`
* [iTerm2 Integrations](../integrations/iterm2.md):
  Get more out of iTerm2 terminal emulator


<hr/>

This document was generated from [gen/integrations/spellcheck_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/integrations/spellcheck_doc.yaml).