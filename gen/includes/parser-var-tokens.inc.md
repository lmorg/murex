**ASCII variable names:**

```
» $example = "foobar"
» out $example
foobar
```

**Unicode variable names:**

Variable names can be non-ASCII however they have to be surrounded by
parenthesis. eg

```
» $(比如) = "举手之劳就可以使办公室更加环保，比如，使用再生纸。"
» out $(比如)
举手之劳就可以使办公室更加环保，比如，使用再生纸。
```

**Infixing inside text:**

Sometimes you need to denote the end of a variable and have text follow on.

```
» $partial_word = "orl"
» out "Hello w$(partial_word)d!"
Hello world!
```

**Variables are tokens:**

Please note the new line (`\n`) character. This is not split using `$`:

```
» $example = "foo\nbar"
```

Output as a string:

```
» out $example
foo
bar
```

Output as an array:

```
» out @example
foo bar
```

The string and array tokens also works for subshells:

```
» out ${ %[Mon..Fri] }
["Mon","Tue","Wed","Thu","Fri"]

» out @{ %[Mon..Fri] }
Mon Tue Wed Thu Fri
```

> `out` will take an array and output each element, space delimited. Exactly
> the same how `echo` would in Bash.

**Variable as a command:**

If a variable is used as a commend then Murex will just print the content of
that variable.

```
» $example = "Hello World!"

» $example
Hello World!
```
