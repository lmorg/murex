```
» set: example="foo\nbar"

» out: $example
foo
bar

» out: @example
foo bar
```

In this example the second command is passing `foo\nbar` (`\n` escaped as a new
line) to `out`. The third command is passing an array of two values: `foo` and
`bar`.

The string and array tokens also works for subshells

```
» out: ${ ja: [Mon..Fri] }
["Mon","Tue","Wed","Thu","Fri"]

» out: @{ ja: [Mon..Fri] }
Mon Tue Wed Thu Fri
```