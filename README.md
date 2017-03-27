# murex
(I'm not sold on that name either. However I am open to suggestions)

## Description

Murex is a cross-platform 

## Dependencies
```
go get github.com/chzyer/readline
go get github.com/kr/pty
```

## Build
```
go build github.com/lmorg/murex
```

Test the binary (requires Bash):
```
test/regression_test.sh
```

## Language guides

Please read the guides:

1. [GUIDE.syntax.md](./GUIDE.syntax.md) - this is recommended first as it gives an overview
if the shell scripting languages syntax and data types.

2. [GUIDE.control-structures.md](./GUIDE.control-structures.md) - this will list how to use if
statements and iteration like for loops.

3. [GUIDE.builtin-functions.md](./GUIDE.builtin-functions.md) - lastly this will list some of the
builtin functions available for this shell.