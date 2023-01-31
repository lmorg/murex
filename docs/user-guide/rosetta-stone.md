# User Guide: Rosetta Stone

> A tabulated list of Bashism's and their equivalent _murex_ syntax

Below is a reference table of common Bash code and how it could be written in
_murex_.

It is also recommended that you read the language [tour](../GUIDE.quick-start.md)
if you want to learn more about shell scripting in _murex_.

| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| [Write to STDOUT](../commands/out.md) | `echo "Hello Bash"` | `out "Hello Murex"` <br/><br/>`echo "Hello Murex"` [[1]](#footnotes)
| [Write to STDERR](commands/err.md) | `echo "Hello Bash" >2` | `err "Hello Murex"` |
| Write to file (truncate) | `echo "Hello Bash" > hello.txt` | `echo "Hello Murex" \|> hello.txt`
| Write to file (append) | `echo "Hello Bash" >> hello.txt` | `echo "Hello Murex" >> hello.txt`
| [Pipe commands](../parser/pipe-arrow.md) | `echo "Hello Bash \| grep Bash` | `echo "Hello Murex \| grep Murex` <br/><br/> `out "Hello Murex" -> regexp m/Murex/` |
| [Redirect errors to STDOUT](../parser/pipe-err.md) | `curl murex.rocks 2>&1 \| less` | `curl murex.rocks ? less` <br/><br/> `curl <!out> murex.rocks \| less` |
| Redirect output to STDERR | `uname -a >&2` | `uname <err> -a` |
| <br/> | | |
| **Quoting strings** | | |
| [Infixing](../parser/double-quote.md) | `echo "Hello $SHELL"` | `out "Hello $SHELL"` |
| [String literals](../parser/single-quote.md) | `echo 'Hello' $SHELL` | `out 'Hello' $SHELL` |
| [Nesting quotes](../parser/brace-quote.md) | `echo 'Hello \'Bob\''` | `out %(Hello 'Bob')` |
| <br/> | | |
| **Process management** | | |
| [Exit number](../commands/exitnum.md) | `$?` | `exitnum` |
| [Background jobs](../commands/bg.md) | `command &` | `bg { command }` |
| [Job control](../commands/fid-list.md) | `ps`,<br/>`jobs`,<br/>`bg pid`,<br/>`fg pid` | `fid-list`,<br/>`jobs`,<br/>`bg fid`,<br/>`fg fid` |
| Happy paths | `command && command` | `command && command` <br/><br/> `try {command; command}` |
| Unhappy paths | `command \|\| command` | `command \|\| command` <br/><br/> `try {command}; catch {command}` |
| Pipe fail | `set -o pipefail` | `runmode trypipe module` <br/><br/> `runmode trypipe function` <br/><br/> `trypipe { commands }`
| <br/> | | |
| **Comments** | | |
| Single line | `# comment` | `# comment` |
| Multiple lines | n/a | `/#`<br/>`line 1`<br/>`line 2`<br/>`#/` |
| Mid-line | n/a | eg `out foo/#comment#/bar`
| <br/> | | |
| **File pattern matching**<br/>(also known as "wildcards") | | |
| [Globbing](../commands/g.md) | eg `ls *.txt` | eg `ls *.txt` (in the interactive terminal) <br/><br/> `g pattern` <br/><br/> eg `ls @{g *.txt}` |
| [Regexp](../commands/rx.md) | n/a | `rx pattern` <br/><br/> eg `ls @{rx '*\\.txt'}` |
| [File type matching](../commands/f.md) | n/a | `f flags` <br/><br/> eg `f +s` (only return symlinks) |
| Chaining | n/a | eg `f +f \| g *.txt \| !g murex.*` <br/> (returns only files with the extension "txt" that aren't called "murex") |
| <br/> | | |
| **Expressions** | | |
| Assignment | `foobar = $((1 + 2 * 3))` | `foobar = 1 + 2 * 3` [[2]](#footnotes) |
| Comparison, string | `[ "$(command)" == "value" ]` | `${command} == "value"` [[2]](#footnotes) |
| Comparison, numeric | `[ $integer -eq 5 ]` | `$number == 5` [[2]](#footnotes) |
| Arithmetic | `echo $(( 1+2*3 ))` | `1 + 2 * 3` [[2]](#footnotes) <br/><br/> `out ${1+2*3}` [[2]](#footnotes) |
| Supported data types | 1. String,<br/>2. Integer<br/>(all variables are strings) | 1. String,<br/>2. Integer,<br/>3. Float (default number type),<br/>4. Boolean<br/>5. Array,<br/>6. Object,<br/>7. Null<br/>(all variables can be treated as strings and/or their primitive) |
| <br/>  | | |
| **Variables**<br/> | | |
| [Assign a local variable](../commands/set.md) | `local foo="bar"` | `foo = "bar"` [[2]](#footnotes)<br/><br/>`set str foo = "bar"` [[2]](#footnotes) <br/><br/>`out "bar" \| set foo` |
| [Assign a global variable](../commands/global.md) | `foo="bar"` | `global str foo = "bar"` [[2]](#footnotes)<br/><br/>`out "bar" \| global foo` |
| [Assign an environmental variable](../commands/export.md) | `export foo="bar"` | `export foo = "bar"` [[2]](#footnotes) [[3]](#footnotes)<br/><br/>`out "bar" \| export foo` [[3]](#footnotes) |
| [Printing a variable](../parser/string.md) | `echo "$foobar"` | `out $foobar` <br/><br/>`$foobar` <br/><br/> (variables don't need to be quoted in _murex_)|
| <br/> | | |
| **Arrays**<br/>(eg arrays, lists) | | |
| Creating an array | `array_name=(value1 value2 value3)` | `%[value1 value2 value3]` <br/><br/>`%[value1, value2, value3]` <br/><br/> eg `array_name = %[1, 2, 3]`, <br/> eg `%[hello world] \| foreach { ... }`|
| Accessing an array element | `${array_name[0]}` | `$array_name[0]` <br/><br/>`array \| [0]` |
| Printing multiple elements | `echo ${array_name[1]} ${array_name[0]}` | `@array_name[1 0]` <br/><br/>`array \| [1 0]` |
| Printing a range of elements | n/a | `@array_name[1..3]` <br/><br/>`array \| [1..3]` |
| [Printing all elements](../parser/array.md) | `echo ${array_name[*]}` | `@array_name` |
| [Iterating through an array](../commands/foreach.md) | `for item in array; do;`<br/>&nbsp;&nbsp;&nbsp;&nbsp;`$item`<br/>`done;` | `array \| foreach item { $item }` <br/><br/> eg `%[Tom Richard Sally] \| foreach name { out "Hello $name" }` |
| <br/> | | |
| **Objects**<br/>(eg JSON objects, maps, hashes, dictionaries) | | |
| Creating an object | n/a | `%{ key: value, array: [1, 2, 3] }` [[2]](#footnotes) <br/><br/> eg `object_name = %{ key: val, arr: [1,3,3] }` <br/> eg `%{ a:1, b:2, c:3 } \| formap { ... }` |
| Accessing an element | n/a | `$object_name[key]` <br/><br/>`object \| [key]` |
| Printing multiple elements | n/a | `$object_name[key1 key2]` <br/><br/> `object \| [key1 key2]` |
| Accessing a nested element | n/a | `$object_name[.path.to.element]]` [[4]](#footnotes)<br/><br/> `object \| [[.path.to.element]]` [[4]](#footnotes)<br/><br/>
| [Iterating through an map](../commands/formap.md) | n/a | `object \| formap key value { $key; $value }` <br/><br/> eg `%{Bob: {age: 10}, Richard: {age: 20}, Sally: {age: 30} } \| formap name person { out "$name is $person[age] years old" }` |
| <br/> | | |
| **Sub-shells**<br/>| | |
| Sub-shell, string | `"$(commands)"` <br/><br/> eg `"echo $(echo "Hello world")"` | `${commands}` <br/><br/> eg `out ${out Hello world}` |
| Sub-shell, arrays | `$(commands)` <br/><br/> eg `$(echo 1 2 3)` | `@{commands}` <br/><br/> eg `out @{ %[1,2,3] }` |

### Footnotes

1. Supported for compatibility with traditional shells like Bash.
2. Unlike Bash, whitespace (or the absence of) is optional.
3. Environmental variables can only be stored as a string. This is a limitation of current operating systems.
4. Path separator can be any 1 byte wide character, eg `/`. The path separator is defined by the first character in a path.

## See Also

* [And (`&&`) Logical Operator](../parser/logical-and.md):
  Continues next operation if previous operation passes
* [Append Pipe (`>>`) Token](../parser/pipe-append.md):
  Redirects STDOUT to a file and append its contents
* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [Murex Named Pipes](../user-guide/namedpipes.md):
  A detailed breakdown of named pipes in _murex_
* [Or (`||`) Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails
* [String (`$`) Token](../parser/string.md):
  Expand values as a string
* [`>>` (append file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [`runmode`](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [`try`](../commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error