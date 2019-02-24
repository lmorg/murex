# lmorg/readline

## Preface

This project began a few years prior to this commit history as an API for murex,
my alternative UNIX shell, because I wasn't satisfied with the state of existing
Go packages for readline (at that time they were either bugger and/or poorly
maintained, or lacked features I desired). The state of things for readline in
Go may have changed since then however own package had also matured and grown to
include many more features that has arisen during the development of my shell.
So it seemed only fair to give back to the community considering it was other
peoples readline libraries that allowed me rapidly prototype my shell during
it's early days of development.

## Apology

I get this README isn't very interesting nor helpful at the moment. I promise I
will embellish this a little more with some useful documentation and fancy GIFs
(etc) as and when I get time. However for now, I would recommend the following:

* Reading the code in `/cmd/main.go` as this is a working example implementation

* https://godoc.org/github.com/lmorg/readline