# Write File (Truncate): `>`

> Writes stdin to disk - overwriting contents if file already exists

## Description

Redirects output to file.

If a file already exists, the contents will be truncated (overwritten).
Otherwise a new file is created.

## Usage

```
<stdin> |> [ -i | --ignore-pipeline-check ] filename

<stdin> |> [ -w | --wait-for-eof ] filename
```

## Examples

```
g * |> files.txt
```

## Flags

* `--ignore-pipeline-check`
    Don't check if _filename_ is a parameter for an earlier command in the pipeline
* `--wait-for-eof`
    Wait for stdin to return an EOF before opening _filename_
* `-i`
    Alias for `--ignore-pipeline`
* `-w`
    Alias for `--wait-for-eof`

## Detail

### Race Condition Detection

If no flags are specified, then `|>` will check if the filename supplied is
used in any parameters for other commands in the pipeline. If it has been, then
`|>` will wait for an **EOF** (End Of File) from stdin before opening _filename_.

This is to allow pipelines like the following:

```
open example.log | regexp m/error/ |> example.log
```

Under traditional shells and Murex's normal scheduler, all commands in a
pipeline would run concurrently. This leads to a race condition where `|>`
opens (and thus truncates) a file before other commands can read from it.

However by default, `|>` will check the pipeline to look for any other
references of _filename_ and if it exists, it will wait for an EOF before
`|>` truncates _filename_.

This wait for EOF behaviour can be forced with the `--wait-for-eof` / `-w`
flag.

Alternatively, if you want to force `|>` to run concurrently then you can
disable the pipeline check with the `--ignore-pipeline-check` / `-i` flag.

#### High Memory Usage

> WARNING! Waiting for EOF will cause `|>` to cache the pipeline into RAM.
> If your pipeline is parsing multi-gigabyte or larger files then you may
> experience performance issues.

For large datasets, it might be preferable to write to a temporary file first.

```
open example.log | regexp m/error/ |> example.log.tmp
mv example.log.tmp example.log
```

The move operation should be instantaneous on most filesystems because your
operating system will just alter filesystem metadata rather than move the file
contents.

### Flag Without A Filename

If you specify a flag without a filename, eg `|> --wait-for-eof`, then it is
assumed that the flag _is_ the filename.

### Syntactic Sugar

While `|>` is referred to as an operator, it's actually a pipe followed by a
builtin:

```
out "foobar" | > example.txt
```

Thus you can actually use `>` by itself.

### Creating An Empty File

If `>` is at the start of a pipeline then it is treated as null input. This a
convenient shortcut to create an empty file or blank an existing file.

**Create a new empty file:**

```
> newfile
```

**Clear a large log file without deleting the file itself:**

```
> large.log
```

### Appending A File

To append a file (ie write at the end of the file without overwriting its
contents) use `>>` instead.

## Synonyms

* `>`
* `|>`
* `fwrite`


## See Also

* [Create Named Pipe: `pipe`](../commands/pipe.md):
  Manage Murex named pipes
* [Create Temporary File: `tmp`](../commands/tmp.md):
  Create a temporary file and write to it
* [Globbing: `g`](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [Read / Write To A Named Pipe: `<pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [Write File (Append): `>>`](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`?` stderr Pipe](../deprecated/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (removed 8.0)
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [builtins/core/io/write_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/write_doc.yaml).