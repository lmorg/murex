{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>

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
- [Footnotes](#footnotes)

</div>

{{ end }}
### Output & error streams
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| [Write to STDOUT](../commands/out.md) | `echo "Hello Bash"` | `out "Hello Murex"` <br/><br/>`echo "Hello Murex"` [[1]](#footnotes)
| [Write to STDERR](commands/err.md) | `echo "Hello Bash" >2` | `err "Hello Murex"` |
| Write to file (truncate) | `echo "Hello Bash" > hello.txt` | `echo "Hello Murex" \|> hello.txt`
| Write to file (append) | `echo "Hello Bash" >> hello.txt` | `echo "Hello Murex" >> hello.txt`
| [Pipe commands](../parser/pipe-arrow.md) | `echo "Hello Bash \| grep Bash` | `echo "Hello Murex \| grep Murex` <br/><br/> `out "Hello Murex" -> regexp m/Murex/` |
| [Redirect errors to STDOUT](../parser/pipe-err.md) | `curl murex.rocks 2>&1 \| less` | `curl murex.rocks ? less` <br/><br/> `curl <!out> murex.rocks \| less` |
| Redirect output to STDERR | `uname -a >&2` | `uname <err> -a` |

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
| Comparison, string | `[ "$(command)" == "value" ]` | `${command} == "value"` [[2]](#footnotes) [[5]](#footnotes) |
| Comparison, numeric | `[ $integer -eq 5 ]` | `$number == 5` [[2]](#footnotes) |
| Arithmetic | `echo $(( 1+2*3 ))` | `1 + 2 * 3` [[2]](#footnotes) <br/><br/> `out ${1+2*3}` [[2]](#footnotes) [[5]](#footnotes) |
| Supported data types | 1. String,<br/>2. Integer<br/>(all variables are strings) | 1. String,<br/>2. Integer,<br/>3. Float (default number type),<br/>4. Boolean<br/>5. Array,<br/>6. Object,<br/>7. Null<br/>(all variables can be treated as strings and/or their primitive) |

### Variables
| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| [Assign a local variable](../commands/set.md) | `local foo="bar"` | `$foo = "bar"` [[2]](#footnotes) [[6]](#footnotes)<br/><br/>`out "bar" \| set foo` |
| [Assign a global variable](../commands/global.md) | `foo="bar"` | `$GLOBAL.foo = "bar"` [[6]](#footnotes)<br/><br/>`out "bar" \| global foo` |
| [Assign an environmental variable](../commands/export.md) | `export foo="bar"` | `export foo = "bar"` [[1]](#footnotes) [[2]](#footnotes) [[3]](#footnotes)<br/><br/>`$ENV.foo = "bar"` [[6]](#footnotes)<br/><br/>`out "bar" \| export foo` [[3]](#footnotes) |
| [Printing a variable](../parser/string.md) | `echo "$foobar"` | `out $foobar` [[5]](#footnotes)<br/><br/>`$foobar` <br/><br/> (variables don't need to be quoted in Murex)|

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

### Footnotes

1. Supported for compatibility with traditional shells like Bash.
2. Unlike Bash, whitespace (or the absence of) is optional.
3. Environmental variables can only be stored as a string. This is a limitation of current operating systems.
4. Path separator can be any 1 byte wide character, eg `/`. The path separator is defined by the first character in a path.
5. Murex uses `${}` for subshells and `$()` for variables, the reverse of what Bash and others use. The reason for this difference is because `{}` always denotes a code block and `()` denotes strings. So `${foobar}` makes more sense as a subshell executing the command `foobar`, while `$(foobar)` makes more sense as the variable `$foobar`.
6. When assigning a variable where the right hand side is an expression, eg `$foo = "bar"`, the dollar prefix is optional. The `set`, `global` and `export` keywords are considered deprecated.