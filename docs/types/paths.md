# `paths` 

> Structured array for working with `$PATH` style data

## Description

The `path` type Turns file and directory paths into structured objects.

The root directory (typically `/`) is counted as a directory. If a path is
relative rather than absolute then `/` will be excluded from outputted string.

## Examples

**Creating a PATH:**

```
» %[/bin, /usr/bin, "$JAVA_HOME/bin"] -> format paths
/bin:/usr/bin:/opt/java/bin
```

**Splitting a PATH:**

```
» $PATH -> :paths: format json
[
    "/bin",
    "/usr/bin",
    "/opt/java/bin"
]
```

**Appending to `$PATH`:**

```
» $PATH -> :paths: append /sbin -> export PATH
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

* [PWD](../variables/PWD.md):
  
* [`MUREX_EXE` (path)](../variables/MUREX_EXE.md):
  Absolute path to running shell
* [`PWDHIST` (json)](../variables/PWDHIST.md):
  History of each change to the sessions working directory
* [`path` ](../types/path.md):
  Structured object for working with file and directory paths

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory
