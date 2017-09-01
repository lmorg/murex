# Install Instructions

Assuming you already have Go (Golang) installed, you can download the
source just by running the following from the command line

    go get -u github.com/lmorg/murex
    cd $GOPATH/src/github.com/lmorg/murex
    go build github.com/lmorg/murex

Test the binary (requires Bash and `timeout`):

    test/regression_test.sh

(A Dockerfile is also included for your convenience. The file is located
in `test/docker` and includes a [README.md](test/docker/README.md) with
more information).

Then to start the shell:

    ./murex

## Required dependencies

Dependencies should be managed by `go get` however for your information
below is a list of packages used by _murex_:

* `github.com/chzyer/readline` used for interactive mode (REPL)

* `github.com/Knetic/govaluate` evaluates the math formulas. This is
exposed via `eval` / `=` and `let`

## Optional dependencies

* `labix.org/v2/mgo/bson`  adds support for BSON (binary JSON) (as used
by MongoDB). This is disabled by default due to a requirement for `bzr`
to exist in $PATH

* `github.com/abesto/sexp` adds support for s-expressions and canonical
s-expressions

* `gopkg.in/yaml.v2` adds support for YAML

* `github.com/BurntSushi/toml` adds support for TOML

* `github.com/hashicorp/hcl` adds support for HCL. Disabled by default
because of converting from JSON to HCL and back to JSON fails to produce
consistent output and there is likely very little demand for HCL anyway

* Image previewing requires a few dependencies:

    1. `github.com/disintegration/imaging`
    2. `golang.org/x/crypto/ssh/terminal`
    3. `golang.org/x/image/bmp`
    4. `golang.org/x/image/tiff`
    5. `golang.org/x/image/webp`

If you wish do disable any of these then delete the appropriate files in
the `builtins` directory of this project or append `// +build ignore` to
the `.go` file if you wish to preserve the change in subsequent updates
from git.