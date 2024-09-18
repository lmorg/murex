# Rosetta Stone

> A tabulated list of Bashism's and their equivalent Murex syntax



Below is a reference table of common Bash code and how it could be written in
Murex.

It is also recommended that you read the language [tour](../tour.md)
if you want to learn more about shell scripting in Murex.

<h2>Table of Contents</h2>

<div id="toc">

- [Output \& error streams](#output--error-streams)
- [Quoting strings](#quoting-strings)
- [Process management](#process-management)
- [Comments](#comments)
- [File pattern matching](#file-pattern-matching)
- [Expressions](#expressions)
- [Variables](#variables)
- [Arrays](#arrays)
- [Objects](#objects)
- [Sub-shells](#sub-shells)
- [Common one-liners](#common-one-liners)
- [Footnotes](#footnotes)

</div>


### Output & error streams
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| [Write to stdout](../commands/out.md) | `echo "Hello Bash"` | `out "Hello Murex"` <br/><br/>`echo "Hello Murex"` [[1]](#footnotes)|
| [Write to stderr](commands/err.md) | `echo "Hello Bash" >2` | `err "Hello Murex"` |
| Write to file (truncate) | `echo "Hello Bash" > hello.txt` | `echo "Hello Murex" \|> hello.txt`|
| Write to file (append) | `echo "Hello Bash" >> hello.txt` | `echo "Hello Murex" >> hello.txt`|
| [Pipe commands](../parser/pipe-arrow.md) | `echo "Hello Bash" \| grep Bash` | `echo "Hello Murex" \| grep Murex` <br/><br/> `out "Hello Murex" -> regexp m/Murex/` |
| [Redirect errors to stdout](../parser/pipe-err.md) | `curl murex.rocks 2>&1 \| less` | `curl <!out> murex.rocks \| less` |
| Redirect output to stderr | `uname -a >&2` | `uname <err> -a` |
| Ignore stderr output | `echo something 2>/dev/null` | `echo <!null> something` |
| Output [ANSI colors and styles](../user-guide/ansi_doc.md) | `echo -e "\n\032[0m\033[1mComplete!\033[0m\n"` | `out "{GREEN}{BOLD}Complete!{RESET}"` |

### Quoting strings
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| [Infixing](../parser/double-quote.md) | `echo "Hello $SHELL"` | `out "Hello $SHELL"` |
| [String literals](../parser/single-quote.md) | `echo 'Hello' $SHELL` | `out 'Hello' $SHELL` |
| [Nesting quotes](../parser/brace-quote.md) | `echo 'Hello \'Bob\''` | `out %(Hello 'Bob')` |

### Process management
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| [Exit number](../commands/exitnum.md) | `$?` | `exitnum` |
| [Background jobs](../commands/bg.md) | `command &` | `bg { command }` |
| [Job control](../commands/fid-list.md) | `ps`,<br/>`jobs`,<br/>`bg pid`,<br/>`fg pid` | `fid-list`,<br/>`jobs`,<br/>`bg fid`,<br/>`fg fid` |
| Happy paths | `command && command` | `command && command` <br/><br/> `try {command; command}` |
| Unhappy paths | `command \|\| command` | `command \|\| command` <br/><br/> `try {command}; catch {command}` |
| Pipe fail | `set -o pipefail` | `runmode trypipe module` <br/><br/> `runmode trypipe function` <br/><br/> `trypipe { commands }`

### Comments
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| Single line | `# comment` | `# comment` |
| Multiple lines | `:<<EOC`<br/>`line 1`<br/>`line 2`<br/>`EOC` | `/#`<br/>`line 1`<br/>`line 2`<br/>`#/` |
| Mid-line | n/a | eg `out foo/#comment#/bar`

### File pattern matching
(also known as "wildcards")
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| [Globbing](../commands/g.md) | eg `ls *.txt` | eg `ls *.txt` (in the interactive terminal) <br/><br/> `g pattern` <br/><br/> eg `ls @{g *.txt}` |
| [Regexp](../commands/rx.md) | n/a | `rx pattern` <br/><br/> eg `ls @{rx '*\\.txt'}` |
| [File type matching](../commands/f.md) | n/a | `f flags` <br/><br/> eg `f +s` (only return symlinks) |
| Chaining | n/a | eg `f +f \| g *.txt \| !g murex.*` <br/> (returns only files with the extension "txt" that aren't called "murex") |

### Expressions
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| Assignment | `foobar = $((1 + 2 * 3))` | `foobar = 1 + 2 * 3` [[2]](#footnotes) |
| Comparison, string | `[ "$(command parameters...)" == "value" ]` | `command(parameters...) == "value"` [[2]](#footnotes) [[7]](#footnotes) <br/><br/> `${command parameters...} == "value"` [[2]](#footnotes) [[5]](#footnotes) |
| Comparison, numeric | `[ $integer -eq 5 ]` | `$number == 5` [[2]](#footnotes) |
| Arithmetic | `echo $(( 1+2*3 ))` | `1 + 2 * 3` [[2]](#footnotes) <br/><br/> `out (1+2*3)` [[2]](#footnotes) [[5]](#footnotes) |
| Supported data types | 1. String,<br/>2. Integer<br/>(all variables are strings) | 1. String,<br/>2. Integer,<br/>3. Float (default number type),<br/>4. Boolean<br/>5. Array,<br/>6. Object,<br/>7. Null<br/>(all variables can be treated as strings and/or their primitive) |

### Variables
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| [Printing a variable](../parser/scalar.md) | `echo "$foobar"` | `out $foobar` [[5]](#footnotes)<br/><br/>`$foobar` <br/><br/> (variables don't need to be quoted in Murex) |
| [Assign a local variable](../commands/set.md) | `local foo="bar"` | `$foo = "bar"` [[2]](#footnotes) [[6]](#footnotes)<br/><br/>`out "bar" \| set $foo` |
| [Assign a global variable](../commands/global.md) | `foo="bar"` | `$GLOBAL.foo = "bar"` [[6]](#footnotes)<br/><br/>`out "bar" \| global $foo` |
| [Assign an environmental variable](../commands/export.md) | `export foo="bar"` | `export foo = "bar"` [[1]](#footnotes) [[2]](#footnotes) [[3]](#footnotes)<br/><br/>`$ENV.foo = "bar"` [[6]](#footnotes)<br/><br/>`out "bar" \| export $foo` [[3]](#footnotes) |
| [Assign with a default value](../parser/null-coalescing.md) | `FOOBAR="${VARIABLE:-default}"` | `$foobar = $variable ?? "default"` |

### Arrays
(eg arrays, lists)
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| Creating an array | `array_name=(value1 value2 value3)` | `%[value1 value2 value3]` <br/><br/>`%[value1, value2, value3]` <br/><br/> eg `array_name = %[1, 2, 3]`, <br/> eg `%[hello world] \| foreach { ... }`|
| Accessing an array element | `${array_name[0]}` | `$array_name[0]` (immutable) <br/><br/>`$array_name.0` (mutable) [[5]](#footnotes) <br/><br/> `array \| [0]` |
| Printing multiple elements | `echo ${array_name[1]} ${array_name[0]}` | `@array_name[1 0]` <br/><br/> `array \| [1 0]` |
| Printing a range of elements | n/a | `@array_name[1..3]` <br/><br/>`array \| [1..3]` |
| [Printing all elements](../parser/array.md) | `echo ${array_name[*]}` | `@array_name` |
| [Iterating through an array](../commands/foreach.md) | `for item in array; do;`<br/>&nbsp;&nbsp;&nbsp;&nbsp;`$item`<br/>`done;` | `array \| foreach item { $item }` <br/><br/> eg `%[Tom Richard Sally] \| foreach name { out "Hello $name" }` |

### Objects
(eg JSON objects, maps, hashes, dictionaries)
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| Creating an object | n/a | `%{ key: value, array: [1, 2, 3] }` [[2]](#footnotes) <br/><br/> eg `object_name = %{ key: val, arr: [1,3,3] }` <br/> eg `%{ a:1, b:2, c:3 } \| formap { ... }` |
| Accessing an element | n/a | `$object_name[key]` (immutable) <br/><br/> `$object_name.key` [[5]](#footnotes) (mutable) <br/><br/> `object \| [key]` |
| Printing multiple elements | n/a | `$object_name[key1 key2]` <br/><br/> `object \| [key1 key2]` |
| Accessing a nested element | n/a | `$object_name[[.path.to.element]]` (immutable) [[4]](#footnotes)<br/><br/> `$object_name.path.to.element` (mutable)<br/><br/> `object \| [[.path.to.element]]` [[4]](#footnotes)<br/><br/>
| [Iterating through an map](../commands/formap.md) | n/a | `object \| formap key value { $key; $value }` <br/><br/> eg `%{Bob: {age: 10}, Richard: {age: 20}, Sally: {age: 30} } \| formap name person { out "$name is $person[age] years old" }` |

### Sub-shells
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| Sub-shell, string | `"$(commands)"` <br/><br/> eg `"echo $(echo "Hello world")"` | `${commands}` [[5]](#footnotes) <br/><br/> eg `out ${out Hello world}` |
| Sub-shell, arrays | `$(commands)` <br/><br/> eg `$(echo 1 2 3)` | `@{commands}` [[5]](#footnotes) <br/><br/> eg `out @{ %[1,2,3] }` |
| In-lined functions | n/a | `function(parameters...)` [[7]](#footnotes) <br/><br/> eg `out uname(-a)` |

### Common one-liners
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| Add `$PATH` entries | `export PATH="$PATH:/usr/local/bin:$HOME/bin"` | The same Bash code works in Murex too. However you can also take advantage of Murex treating `$PATH` as an array <br/><br/>`%[ @PATH /usr/local/bin "$HOME/bin" ] \| format paths \| export $PATH` |
| Iterate directories | `for i in $(find . -maxdepth 1 -mindepth 1 -type d); do`<br/>&nbsp;&nbsp;&nbsp;&nbsp;`echo $i`<br/>`done`| `f +d \| foreach $dir {`<br/>&nbsp;&nbsp;&nbsp;&nbsp;`out $i`<br/>`}` |
| If `$dir` exists... | `if [ -d "$dir" ]; then`<br/>&nbsp;&nbsp;&nbsp;&nbsp;`# exists`<br/>`fi` | `if { g $dir \| f +d } then {`<br/>&nbsp;&nbsp;&nbsp;&nbsp;`# exists`<br/>`}` |
| Print current directory | `result=${PWD##*/}; result=${result:-/}; printf '%s' "${PWD##*/}"` ([read more](https://stackoverflow.com/a/1371283)) | `$PWD[-1]` |
### Footnotes

1. Supported for compatibility with traditional shells like Bash.
2. Unlike Bash, whitespace (or the absence of) is optional.
3. Environmental variables can only be stored as a string. This is a limitation of all major operating systems.
4. Path separator can be any 1 byte wide character, eg `/`. The path separator is defined by the first character in a path.
5. Murex uses `${}` for sub-shells and `$()` for variables, the reverse of what Bash and others use. The reason for this difference is because `{}` always denotes a code block and `()` denotes a sub-expression or string. So `${foobar}` makes more sense as a sub-shell syntax executing a block, while `$(foobar)` makes more sense as the syntax for a scalar.
6. When assigning a variable where the right hand side is an expression, eg `$foo = "bar"`, the dollar prefix is optional. The `set`, `global` and `export` keywords are considered deprecated.
7. The `command(parameters...)` only works for commands who's names match the following regexp pattern: `[._a-zA-Z0-9]+`. Which is exclusively uppercase and lowercase English letters, numbers, full stop, and underscore.

## See Also

* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Function / Module Defaults (`runmode`)](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Named Pipes](../user-guide/namedpipes.md):
  A detailed breakdown of named pipes in Murex
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [Truncate File (`>`)](../parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`&&` And Logical Operator](../parser/logical-and.md):
  Continues next operation if previous operation passes
* [`>>` Append File](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists
* [`>>` Append File](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`string` (stringing)](../types/str.md):
  string (primitive)
* [`||` Or Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails
* [index](../parser/item-index.md):
  Outputs an element from an array, map or table

<hr/>

This document was generated from [gen/user-guide/rosetta-stone_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/rosetta-stone_doc.yaml).