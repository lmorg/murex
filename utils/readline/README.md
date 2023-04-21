# lmorg/readline

## Preface

This project began a few years prior to this git commit history as an API for
_[murex](https://github.com/lmorg/murex)_, an alternative UNIX shell, because
I wasn't satisfied with the state of existing Go packages for readline (at that
time they were either bugger and/or poorly maintained, or lacked features I
desired). The state of things for readline in Go may have changed since then
however own package had also matured and grown to include many more features
that has arisen during the development of Murex. So it seemed only fair to
give back to the community considering it was other peoples readline libraries
that allowed me rapidly prototype Murex during it's early stages of
development.

## `readline` In Action

[![asciicast](https://asciinema.org/a/232714.svg)](https://asciinema.org/a/232714)

This is a very rough and ready recording but it does demonstrate a few of the
features available in `readline`. These features include:
* hint text (the blue status text below the prompt - however the colour is 
  configurable)
* syntax highlighting (albeit there isn't much syntax to highlight in the
  example)
* tab-completion in gridded mode (seen when typing `cd`)
* tab-completion in list view (seen when selecting an process name to `kill`
  and the process ID was substituted when selected)
* regex searching through the tab-completion suggestions (seen in both `cd` and
  `kill` - enabled by pressing `[CTRL+f]`)
* line editing using `$EDITOR` (`vi` in the example - enabled by pressing
  `[ESC]` followed by `[v]`)
* `readline`'s warning before pasting multiple lines of data into the buffer
* the preview option that's available as part of the aforementioned warning
* and VIM keys (enabled by pressing `[ESC]`)

## Example Code

Using `readline` is as simple as:

```go
func main() {
    rl := readline.NewInstance()

    for {
        line, err := rl.Readline()
        
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        fmt.Printf("You just typed: %s\n", line)
    }
}
```

However I suggest you read through the examples in `/examples` for help using
some of the more advanced features in `readline`.

The source for `readline` will also be documented in godoc: https://godoc.org/github.com/lmorg/readline

## Version Information

Because the last thing a developer wants is to do is fix breaking changes after
updating modules, I will make the following guarantees:

* The version string will be based on Semantic Versioning. ie version numbers
  will be formatted `(major).(minor).(patch)` - for example `2.0.1`

* `major` releases _will_ have breaking changes. Be sure to read CHANGES.md for
  upgrade instructions

* `minor` releases will contain new APIs or introduce new user facing features
  which may affect useability from an end user perspective. However `minor`
  releases will not break backwards compatibility at the source code level and
  nor will it break existing expected user-facing behavior. These changes will
  be documented in CHANGES.md too

* `patch` releases will be bug fixes and such like. Where the code has changed
  but both API endpoints and user experience will remain the same (except where
  expected user experience was broken due to a bug, then that would be bumped
  to either a `minor` or `major` depending on the significance of the bug and
  the significance of the change to the user experience)

* Any updates to documentation, comments within code or the example code will
  not result in a version bump because they will not affect the output of the
  go compiler. However if this concerns you then I recommend pinning your
  project to the git commit hash rather than a `patch` release

My recommendation is to pin to either the `minor` or `patch` release and I will
endeavour to keep breaking changes to an absolute minimum.

## License Information

`readline` is licensed Apache 2.0. All the example code and documentation in
`/examples` is public domain.