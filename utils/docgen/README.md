# docgen

## `docgen` is a tool for auto-generating technical documents

The purpose of this tool is to automate documentation as
much as possible and have the output a static CMS which
can be hosted on GitHub or S3 (depending on whether your
templates are markdown or HTML).

It was written originally for the _murex_ project (https://github.com/lmorg/murex)
however it is flexible enough to be useful in other projects
as well.

All config if defined in YAML files and template files
(using the templating engine from Go's (golang) core libs)

## Usage

    docgen -config docgen.yaml

## Flags

    -config   takes a string parameter for a path to a YAML file
    -verbose  optional flag to enable log messages
    -debug    optional flag to enable stack a trace on error

## Don't Panic!

You may notice there are a lot of `panic()`'s in the codebase.
This is deliberate because by we want any errors to fail the
utility (essentially raising all errors as exceptions).
However the panics are caught in `main()` if by default if
`-debug` isn't set (see "flags" section above) so you get
friendly error messages as standard but a stack trace on all
errors for debugging.

I get this isn't the typical Go idiom but it makes more sense
for this particular application given it serves one function
and any errors beyond that should fail the program.
