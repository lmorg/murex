# Language Guide: Control Structures

## if

`if` supports 3 "modes" depending on the number of parameters:

1. Method If: `conditional -> if: { true }`
2. Function If: `if: { conditional } { true }`
3. If Else: `if: { conditional } { true } { false }`

The conditional is evaluated based on the output produced by the function
and the exit number. Any non-zero exit numbers are an automatic "false".
Any functions returning no data are also classed as a "false". For a full
list of conditions that are evaluated to determine a true or false state
of a function, please read the documentation on the `boolean` data type
in [GUIDE.syntax.md](GUIDE.syntax.md#boolean).

#### Method If

This is where the conditional is evaluated from the result of the
piped function.
```
out: hello world | grep: world -> if: { out: world found }
```

#### Function If

This is where the conditional is evaluated from the first parameter.
```
if: { out: hello world | grep: world } { out: world found }
```

#### If Else

Same as a Function If but with an Else block. This is the only usage of
`if` that doesn't set the exit status to zero (0) if the "if" condition
doesn't match "true".
```
if: { out: hello world | grep: world } { out: world found } { out: world missing }
```

## else / !if

`if` also supports an anti-alias where the conditional is "notted".
`else` is also an alias for `!if`.
```
out: hello world | grep: world -> if: { out: world found } -> else { out: world missing }
```

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
This also includes `if` statements so be very careful to use the three
parameter `if` which handles the "else" condition without raising an
error.

To handle errors in `try`, pipe the result into `else`.
```
try: { out: "hello world" -> match: "foobar"; out: "other stuff" } -> else { out: "error raised" } 
```