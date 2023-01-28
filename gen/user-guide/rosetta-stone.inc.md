| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| Write to STDOUT | `echo "Hello Bash"` | `out "Hello Murex"` </br></br>`echo "Hello Murex"` [1]
| Write to STDERR | `echo "Hello Bash" >2` | `err "Hello Murex"` |
| Write to file (truncate) | `echo "Hello Bash" > hello.txt` | `echo "Hello Murex" \|> hello.txt`
| Write to file (append) | `echo "Hello Bash" >> hello.txt` | `echo "Hello Murex" >> hello.txt`
| Pipe commands | `echo "Hello Bash \| grep Bash` | `echo "Hello Murex \| grep Murex` </br></br> `out "Hello Murex" -> regexp m/Murex/` |
| Redirect errors to STDOUT | `uname -a 2>&1` | `uname -a ? next-command` </br></br>`uname <!out> -a` |
| Redirect output to STDERR | `uname -a >&2` | `uname <err> -a` |
| </br> | | |
| **Quoting strings** | | |
| | | |
| Infixing | `echo "Hello $SHELL"` | `out "Hello $SHELL"` |
| String literals  | `echo 'Hello' $SHELL` | `out 'Hello' $SHELL` |
| Nesting quotes | `echo 'Hello \'Bob\''` | `out %(Hello 'Bob')` |
| </br> | | |
| **Process management** | | |
| | | |
| Exit number | `$?` | `exitnum` |
| Background jobs | `command &` | `bg { command }` |
| Job control | `ps`,</br>`jobs`,</br>`bg pid`,</br>`fg pid` | `fid-list`,</br>`jobs`,</br>`bg fid`,</br>`fg fid` |
| Happy paths | `command && command` | `command && command` </br></br> `try {command; command}` |
| Unhappy paths | `command \|\| command` | `command \|\| command` </br></br> `try {command}; catch {command}` |
| Pipe fail | `set -o pipefail` | `runmode trypipe module` </br></br> `runmode trypipe function` </br></br> `trypipe { commands }`
| </br> | | |
| **Comments** | | |
| | | |
| Single line | `# comment` | `# comment` |
| Multiple lines | n/a | `/#`</br>`line 1`</br>`line 2`</br>`#/` |
| Mid-line | n/a | eg `out foo/#comment#/bar`
| </br> | | |
| **File pattern matching** | also known as "wildcards" | |
| | | |
| Globbing | eg `ls *.txt` | eg `ls *.txt` (in the interactive terminal) </br></br> `g pattern` </br></br> eg `ls @{g *.txt}` |
| Regexp | n/a | `rx pattern` </br></br> eg `ls @{rx '*\\.txt'}` |
| File type matching | n/a | `f flags` </br></br> eg `f +s` (only return symlinks) |
| Chaining | n/a | eg `f +f \| g *.txt \| !g murex.*` </br> (returns only files with the extension "txt" that aren't called "murex") |
| </br> | | |
| **Expressions** | | |
| | | |
| Assignment | `foobar = $((1 + 2 * 3))` | `foobar = 1 + 2 * 3` [2] |
| Comparison, string | `[ "$(command)" == "value" ]` | `${command} == "value"` [2] |
| Comparison, numeric | `[ $integer -eq 5 ]` | `$number == 5` [2] |
| Arithmetic | `echo $(( 1+2*3 ))` | `1 + 2 * 3` [2] </br></br> `out ${1+2*3}` [2] |
| Supported data types | 1. String,</br>2. Integer</br>(all variables are strings) | 1. String,</br>2. Integer,</br>3. Float (default number type),</br>4. Boolean</br>5. Array,</br>6. Object,</br>7. Null</br>(all variables can be treated as strings and/or their primitive) |
| </br>  | | |
| **Variables** | | |
| | | |
| Assign a local variable | `local foo="bar"` | `foo = "bar"` [2]</br></br>`set str foo = "bar"` [2] </br></br>`out "bar" \| set foo` |
| Assign a global variable | `foo="bar"` | `global str foo = "bar"` [2]</br></br>`out "bar" \| global foo` |
| Assign an environmental variable | `export foo="bar"` | `export foo = "bar"` [2] [3]</br></br>`out "bar" \| export foo` [3] |
| Printing a variable | `echo "$foobar"` | `out $foobar` </br></br>`$foobar` </br></br> (variables don't need to be quoted in _murex_)|
| </br> | | |
| **Arrays** | eg arrays, lists | |
| | | |
| Creating an array | `array_name=(value1 value2 value3)` | `%[value1 value2 value3]` </br></br>`%[value1, value2, value3]` </br></br> eg `array_name = %[1, 2, 3]`, </br> eg `%[hello world] \| foreach { ... }`|
| Accessing an array element | `${array_name[0]}` | `$array_name[0]` </br></br>`array \| [0]` |
| Printing multiple elements | `echo ${array_name[2]} ${array_name[0]}` | `@array_name[1 0]` </br></br>`array \| [1 0]` |
| Printing a range of elements | n/a | `@array_name[1..3]` </br></br>`array \| [1..3]` |
| Printing all elements | `echo ${array_name[*]}` | `@array_name` |
| Iterating through an array | `for item in array; do;`</br>&nbsp;&nbsp;&nbsp;&nbsp;`$item`</br>`done;` | `array \| foreach item { $item }` </br></br> eg `%[Tom Richard Sally] \| foreach name { out "Hello $name" }` |
| </br> | | |
| **Objects** | eg JSON objects, maps, hashes, dictionaries | |
| | | |
| Creating an object | n/a | `%{ key: value, array: [1, 2, 3] }` [2] </br></br> eg `object_name = %{ key: value, array: [1,3,3] }` </br> eg `%{Bob: {age: 10}, Richard: {age: 20}, Sally: {age: 30} } \| formap { ... }` |
| Accessing an element | n/a | `$object_name[key]` </br></br>`object \| [key]` |
| Printing multiple elements | n/a | `$object_name[key1 key2]` </br></br>`object \| [key1 key2]` |
| Accessing a nested element | n/a | `$object_name[.path.to.element]]` [4]</br></br> `object \| [[.path.to.element]]` [4]</br></br>
| Iterating through an map | n/a | `object \| formap key value { $key; $value }` </br></br> eg `%{Bob: {age: 10}, Richard: {age: 20}, Sally: {age: 30} } \| formap name person { out "$name is $person[age] years old" }` |
| </br> | | |
| **Sub-shells** | | |
| | | |
| Sub-shell, string | `"$(commands)"` </br></br> eg `"echo $(echo "Hello world")"` | `${commands}` </br></br> eg `out ${out Hello world}` |
| Sub-shell, arrays | `$(commands)` </br></br> eg `$(echo 1 2 3)` | `@{commands}` </br></br> eg `out @{ %[1,2,3] }` |
| | | |

### Footnotes

1. Supported for compatibility with traditional shells like Bash.
2. Unlike Bash, whitespace (or the absence of) is optional.
3. Environmental variables can only be stored as a string. This is a limitation of current operating systems.
4. Path separator can be any 1 byte wide character, eg `/`. The path separator is defined by the first character in a path.