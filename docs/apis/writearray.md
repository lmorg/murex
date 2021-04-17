# _murex_ Shell Docs

## API Reference: `WriteArray()` 

> Write a data type, one array element at a time

## Description

This is a function you would write when programming a _murex_ data-type.

It's called by builtins to allow them to write data structures one array
element at a time.

## Usage

Registering your WriteArray()

```go
// To avoid confusion, this should only happen inside func init()
stdio.RegisterWriteArray(/* your type */, /* your writeArray func */)
```

## Examples

Example WriteArray() function:

```go
package string

import (
	"github.com/lmorg/murex/lang/proc/stdio"
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
	"github.com/lmorg/murex/lang/proc/stdio"
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

## Parameters

1. `string`: array element to write

## See Also

* [apis/readearray](../apis/readearray.md):
  
* [apis/readmap](../apis/readmap.md):
  