# `xml`

> Extensible Markup Language (XML) (experimental)

## Description

XML is a structured data-type within Murex.

[Wikipedia](https://en.wikipedia.org/wiki/XML#Applications) describes XML usage as:

> XML has come into common use for the interchange of data over the Internet.
> Hundreds of document formats using XML syntax have been developed, including
> RSS, Atom, Office Open XML, OpenDocument, SVG, COLLADA, and XHTML. XML also
> provides the base language for communication protocols such as SOAP and XMPP.
> It is one of the message exchange formats used in the Asynchronous JavaScript
> and XML (AJAX) programming technique.

## Examples

```
<person>
    <firstName>John</firstName>
    <lastName>Smith</lastName>
    <isAlive>true</isAlive>
    <age>27</age>
    <address>
        <city>New York</city>
        <postalCode>10021-3100</postalCode>
        <state>NY</state>
        <streetAddress>21 2nd Street</streetAddress>
    </address>
    <phoneNumbers>
        <number>212 555-1234</number>
        <type>home</type>
    </phoneNumbers>
    <phoneNumbers>
        <number>646 555-4567</number>
        <type>office</type>
    </phoneNumbers>
    <phoneNumbers>
        <number>123 456-7890</number>
        <type>mobile</type>
    </phoneNumbers>
    <children/>
    <spouse/>
</person>
```

## Default Associations

* **Extension**: `atom`
* **Extension**: `rss`
* **Extension**: `svg`
* **Extension**: `xht`
* **Extension**: `xhtml`
* **Extension**: `xml`
* **MIME**: `+xml`
* **MIME**: `application/x-xml`
* **MIME**: `application/xml`
* **MIME**: `application/xml-dtd`
* **MIME**: `application/xml-external-parsed-entity`
* **MIME**: `text/x-xml`
* **MIME**: `text/xml`
* **MIME**: `text/xml-external-parsed-entity`


## Supported Hooks

* `Marshal()`
    Writes minified XML when no TTY detected and indented XML when stdout is a TTY
* `ReadArray()`
    experimental; Works with XML arrays. Maps are converted into arrays
* `ReadArrayWithType()`
    experimental; Works with XML arrays. Maps are converted into arrays. Elements data-type in Murex mirrors the XML type of the element (if known)
* `ReadIndex()`
    experimental; Works against all properties in XML
* `ReadMap()`
    experimental; Works with XML maps
* `ReadNotIndex()`
    experimental; Works against all properties in XML
* `Unmarshal()`
    Supported
* `WriteArray()`
    experimental; Works with XML arrays

## See Also

* [Define Type: `cast`](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Get Nested Element: `[[ Element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [Open File: `open`](../commands/open.md):
  Open a file with a preferred handler
* [Prettify Objects: `pretty`](../commands/pretty.md):
  Prettifies data documents to make it human readable
* [Reformat Data Type: `format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`csv`](../types/csv.md):
  CSV files (and other character delimited tables)
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`toml`](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml`](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [index](../parser/item-index.md):
  Outputs an element from an array, map or table

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/types/xml/xml_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/xml/xml_doc.yaml).