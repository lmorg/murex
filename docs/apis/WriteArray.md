# _murex_ Shell Docs

## API Reference: `WriteArray()` (type)

> Write a data type, one array element at a time

## Description

This is a function you would write when programming a _murex_ data-type.

It's called by builtins to allow them to write data structures one array
element at a time.

The purpose of this function is to allow builtins to support sequential writes
(where possible) and also create a standard interface for builtins, thus
allowing them to be data-type agnostic.

### A Collection of Functions

`WriteArray()` should return a `struct` that satisfies the following
`interface{}`:

```go
package stdio

// ArrayWriter is a simple interface types can adopt for buffered writes of formatted arrays in structured types (eg JSON)
type ArrayWriter interface {
	Write([]byte) error
	WriteString(string) error
	Close() error
}
```

## Usage

Registering your `WriteArray()`

```go
// To avoid confusion, this should only happen inside func init()
stdio.RegisterWriteArray(/* your type name */, /* your writeArray func */)
```

## Examples

Example `WriteArray()` function:

```go
package string

import (
	"github.com/lmorg/murex/lang/stdio"
)

type arrayWriter struct {
	writer stdio.Io
}

func newArrayWriter(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{writer: writer}
	return w, nil
}

func (w *arrayWriter) Write(b []byte) error {
	_, err := w.writer.Writeln(b)
	return err
}

func (w *arrayWriter) WriteString(s string) error {
	_, err := w.writer.Writeln([]byte(s))
	return err
}

func (w *arrayWriter) Close() error { return nil }
```

## Detail

Since not all data types will be stream-able (for example `json`), some types
may need to cache the array and then to write it once the array writer has been
closed.

```go
package json

import (
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/json"
)

type arrayWriter struct {
	array  []string
	writer stdio.Io
}

func newArrayWriter(writer stdio.Io) (stdio.ArrayWriter, error) {
	w := &arrayWriter{writer: writer}
	return w, nil
}

func (w *arrayWriter) Write(b []byte) error {
	w.array = append(w.array, string(b))
	return nil
}

func (w *arrayWriter) WriteString(s string) error {
	w.array = append(w.array, s)
	return nil
}

func (w *arrayWriter) Close() error {
	b, err := json.Marshal(w.array, w.writer.IsTTY())
	if err != nil {
		return err
	}

	_, err = w.writer.Write(b)
	return err
}
```

## See Also

* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [apis/`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [apis/`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [apis/`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin