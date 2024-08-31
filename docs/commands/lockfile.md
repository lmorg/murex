# Lock Files (`lockfile`)

> Create and manage lock files

## Description

`lockfile` is used to create and manage lock files

## Usage

Create a lock file with the name `identifier`

```
lockfile lock identifier
```

Delete a lock file with the name `identifier`

```
lockfile unlock identifier
```

Wait until lock file with the name `identifier` has been deleted

```
lockfile wait identifier
```

Output the the file name and path of a lock file with the name `identifier`

```
lockfile path identifier -> <stdout>
```

## Examples

```
lockfile lock example
out "lock file created: ${lockfile path example}"

bg {
    sleep 10
    lockfile unlock example
}

out "waiting for lock file to be deleted (sleep 10 seconds)...."
lockfile wait example
out "lock file gone!"
```

## Synonyms

* `lockfile`


## See Also

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character

<hr/>

This document was generated from [builtins/core/io/lockfile_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/lockfile_doc.yaml).