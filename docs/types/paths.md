# `paths`

> Structured array for working with `$PATH` style data

## Description

The `path` type Turns file and directory paths into structured objects.

The root directory (typically `/`) is counted as a directory. If a path is
relative rather than absolute then `/` will be excluded from outputted string.

## Examples

### Creating a PATH

```
» %[/bin, /usr/bin, "$JAVA_HOME/bin"] -> format paths
/bin:/usr/bin:/opt/java/bin
```

### Splitting a PATH

```
» $PATH -> :paths: format json
[
    "/bin",
    "/usr/bin",
    "/opt/java/bin"
]
```

### Appending to $PATH

As a statement:

```
» $PATH -> append /sbin -> export PATH
» $PATH
/bin:/usr/bin:/opt/java/bin:/sbin
```

As an expression:

```
» $PATH <~ %[ "/sbin" ]
» $PATH
/bin:/usr/bin:/opt/java/bin:/sbin
```

## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Each element is a directory branch. Root, `/`, is treated as it's own element
* `ReadArrayWithType()`
    Same as `ReadArray()`
* `ReadIndex()`
    Returns a directory branch or filename if last element is a file
* `ReadMap()`
    Not currently supported
* `ReadNotIndex()`
    Supported
* `Unmarshal()`
    Supported
* `WriteArray()`
    Each element is a directory branch

## See Also

* [MUREX_EXE](../variables/murex_exe.md):
  Absolute path to running shell
* [PWD](../variables/pwd.md):
  Current working directory
* [PWDHIST](../variables/pwdhist.md):
  History of each change to the sessions working directory
* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays
* [`path`](../types/path.md):
  Structured object for working with file and directory paths
* [assign-merge](../types/assign-merge.md):
  

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/types/paths/paths_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/paths/paths_doc.yaml).