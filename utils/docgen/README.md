# docgen

## `docgen` is a tool for auto-generating technical documents

The purpose of this tool is to automate documentation as
much as possible and have the output a static CMS which
can be hosted on GitHub or S3 (depending on whether your
templates are markdown or HTML).

It was written originally for the Murex project (https://github.com/lmorg/murex)
however it is flexible enough to be useful in other projects
as well.

All config if defined in YAML files and template files
(using the templating engine from Go's (golang) core libs)

## Usage

    docgen -config docgen.yaml

## Flags

    -config   takes a string parameter for a path to a YAML file
    -verbose  optional flag to enable log and warning messages
    -warning  optional flag to enable warning messages
    -panic    optional flag to enable stack a trace on error
    -readonly optional flag to prevent files getting written to disk
              (use -readonly to test your config, templates, etc)

## Don't Panic!

You may notice there are a lot of `panic()`'s in the codebase.
This is deliberate because by we want any errors to fail the
utility (essentially raising all errors as exceptions). However
the panics are caught in `main()` by default (ie if `-panic`
isn't set) so you get friendly error messages as standard but a
stack trace on all errors for debugging if needed.

I get this isn't the typical Go idiom but it makes more sense
for this particular application given it serves one function
and any errors beyond that should fail the program.
