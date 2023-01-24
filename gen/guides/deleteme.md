| Description   | Bash          | Murex  |
|---------------|---------------|--------|
| Assign a local variable | `local foo="bar"` | `foo = "bar"` [1]</br> `set foo = "bar"` [1] </br> `out "bar" \| set foo` |
| Assign a global variable | `foo="bar"` | `global foo = "bar"` [1]</br> `out "bar" \| global foo` |
| Assign an environmental variable | `export foo="bar"` | `export foo = "bar"` [1]</br> |`out "bar" \| export foo`
||||
| Print a variable | `echo $foo` | `out $foo` </br> `$foo` |
||||
| Print to STDOUT | `echo "Hello Bash"` | `out "Hello Murex"` </br> `echo "Hello Murex"` [2]
| Print to STDERR | `echo "Hello Bash" >2` | `err "Hello Murex"` |
| Redirect errors to STDOUT | `uname -a 2>&1` | `uname -a ? next-command` </br> `uname <!out> -a` |
| Redirect output to STDERR | `uname -a >&2` | `uname <err> -a` |
||||
| String, infixing | `echo "Hello $SHELL"` | `out "Hello $SHELL"` |
| String, literal  | `echo 'Hello' $SHELL` | `out 'Hello' $SHELL` |
| String, nested quotes | `echo 'Hello \'Bob\''` | `out %(Hello 'Bob')` |
||||
| Subshell, string | `echo "$(echo "Hello world")"` | `out ${out Hello world}` |
| Subshell, arrays | `echo $(echo 1 2 3)` | `out @{a: [1..3]}` |
||||
| Creating an array | `array_name=(value1 value2 value3)` | `array_name = %[value1 value2 value3]` [1] |
| Accessing an array element | `${array_name[0]}` | `$array_name[0]` |
| Printing multiple elements | `echo ${array_name[1]} ${array_name[0]}` | `out @array_name[1 0]` |
| Printing all elements | `echo ${array_name[*]}` | `out @array_name` |

### Footnotes

1. Unlike Bash, whitespace (or the absence of) is optional
2. Supported for compatibility with traditional shells like Bash
