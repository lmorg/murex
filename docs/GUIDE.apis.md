# _murex_ Language Guide

## API Reference

This section is a glossary of APIs.

These APIs are provided for reference for any developers wishing to write
their own builtins. However some APIs are still worth being aware of even
when just writing _murex_ scripts because they provide a background into
the internal logic of _murex_'s runtime.

### Pages

* [`Marshal()` ](apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`Unmarshal()` ](apis/unmarshal.md):
  Converts a structured file format into structured memory
* [`define.MarshalData()` ](apis/marshaldata.md):
  Converts structured memory into a _murex_ data-type (eg for stdio)
* [`define.UnmarshalData()` ](apis/unmarshaldata.md):
  Converts a _murex_ data-type into structured memory