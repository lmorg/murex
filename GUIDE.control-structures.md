# Language Guide: Control Structures

## if

`if` supports 2 "modes":

1. Method If: `conditional -> if: { true } { false }`
2. Function If: `if: { conditional } { true } { false }`

The conditional is evaluated based on the output produced by the
function and the exit number. Any non-zero exit numbers are an automatic
"false". Any functions returning no data are also classed as a "false".
For a full list of conditions that are evaluated to determine a true or
false state of a function, please read the documentation on the `boolean`
data type in [GUIDE.syntax.md](GUIDE.syntax.md#boolean).

Please also note that while the last parameter is optional, if it is
left off and `if` or `!if` would have otherwise called it, then `if` /
`!if` will return a non-zero exit number. The significance of this is
important when using `if` or `!if` inside a `try` block.

#### Method If

This is where the conditional is evaluated from the result of the piped
function. The last parameter is optional.
```
# if / then
out: hello world | grep: world -> if: { out: world found }

# if / then / else
out: hello world | grep: world -> if: { out: world found } { out: world missing }

# if / else
out: hello world | grep: world -> !if: { out: world missing }
```

#### Function If

This is where the conditional is evaluated from the first parameter. The
last parameter is optional.
```
# if / then / else
if: { out: hello world | grep: world } { out: world found }

# if / then / else
if: { out: hello world | grep: world } { out: world found } { out: world missing }

# if / else
!if: { out: hello world | grep: world } { out: world missing }
```

## !if

`if` also supports an anti-alias which will "not" the conditional,
effectively reversing the "then" and "else" parameters. See `if` (above)
for examples.

## foreach

(description to follow)
```
fuction_with_listed_output -> foreach: variable { iteration } 
```

## while
(description to follow)
```
while: { conditional } { iteration } 
```

## try

This will force a different execution behavior. All pipelined processes
will become sequential (unlike normally when they run in parallel) and
any exit numbers not equal to zero (0) will terminate the code block.
This also includes `if` statements so be very careful to include an else
parameter, even if it's an empty block, so `if` doesn't raise an error.

If the try block fails then try will raise a non-zero exit number. If
you want to run an alternative block of code in an event of a failure
then combine with the `catch` method.
```
# try
try: { out: hello world | grep: foobar; out: other stuff }

# try / catch
try: { out: hello world | grep: foobar; out: other stuff }
    -> catch { out: `try` failed }
```

## catch

This works a little like the single parameter `!if` method except it
only checks the exit number (not stdin) and the stdin stream is simply
forwarded along the chain.

`catch` is typically used alongside `try` but it can also be used on its
own where you want to check the success of a routine while preserving
its stdout stream.

Use `!catch` to "else" the `try`

```
# try / catch
try: { out: hello world | grep: foobar; out: other stuff } -> catch { out: `try` failed }

# catch
out: hello world | grep: foobar -> catch { out: foobar not found }

# !catch
out: hello world | grep: world -> !catch { out: world found }

# else
try: { out: hello world | grep: foobar; out: other stuff }
    -> catch  { out: `try` failed }
    -> !catch { out: `try` succeeded }
```

`catch` also supports anti-alias (`!catch`) where the code block only
executes if the exit number equals zero.
