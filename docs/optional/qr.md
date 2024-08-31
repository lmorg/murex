# `qr`

> Creates a QR code from stdin

## Description

`qr` is an optional builtin which generates a PNG format image based on the
input from stdin. `qr` must be run as a method.

## Usage

```
<stdin> -> qr -> <stdout>
```

## Examples

### Write PNG to disk

```
» out "Hello, World!" -> qr |> qr.png
```

### Display PNG in terminal

```
» out "Hello, World!" -> qr -> open-image
```

## Detail

`qr` sets stdout's data-type to be "image", which is defined in with the
`open-image` optional builtin. So if you have that disabled then you may
have to `cast` the output in some circumstances.

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Render Image In Terminal (`open-image`)](../commands/open-image.md):
  Renders bitmap image data on your terminal

<hr/>

This document was generated from [builtins/optional/qr/qr_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/optional/qr/qr_doc.yaml).