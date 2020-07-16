# _murex_ Shell Docs

## Command Reference: `test`

> _murex_'s test framework - define tests, run tests and debug shell scripts

## Description

`test` is used to define tests, run tests and debug _murex_ shell scripts.

## Usage

Define an inlined test

    test: define test-name { json-properties }
    
Define a state report

    test: state name { code block }
    
Define a unit test

    test: unit function|private|open|event test-name { json-properties }
    
Enable or disable boolean test states (more options available in `config`)

    test: config [ enable|!enable ] [ verbose|!verbose ] [ auto-report|!auto-report ]
    
Disable test mode

    !test
    
Execute a function with testing enabled

    test: run { code-block }
    
Execute unit test(s)

    test: run-unit package[/module[/test-name]|*

## Synonyms

* `test`
* `!test`


## See Also

* [commands/`<>` (murex named pipe)](../commands/namedpipe.md):
  Reads from a _murex_ named pipe
* [commands/`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [parser/namedpipe](../parser/namedpipe.md):
  