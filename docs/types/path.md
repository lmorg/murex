# `path`

> Structured object for working with file and directory paths

## Description

The `path` type Turns file and directory paths into structured objects.

The root directory (typically `/`) is counted as a directory. If a path is
relative rather than absolute then `/` will be excluded from outputted string.

## Examples

**Return the first two elements in a path:**

```
» $PWD[..2]
/Users/
```

**Check if path exists:**

```
» set path foobar="/dev/foobar"
» $foobar.2.Exists
```

**Example of `path` data structure:**

```
» set path foobar="/dev/foobar"
» $foobar -> format json
[
    {
        "Exists": true,
        "IsDir": true,
        "IsRelative": false,
        "Value": "/"
    },
    {
        "Exists": true,
        "IsDir": true,
        "IsRelative": false,
        "Value": "dev"
    },
    {
        "Exists": false,
        "IsDir": false,
        "IsRelative": false,
        "Value": "foobar"
    }
]
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

* [`MUREX_EXE` (path)](../variables/MUREX_EXE.md):
  Absolute path to running shell
* [`PWDHIST` (json)](../variables/PWDHIST.md):
  History of each change to the sessions working directory
* [`PWD` (str)](../variables/PWD.md):
  Current working directory
* [`paths`](../types/paths.md):
  Structured array for working with `$PATH` style data

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory
