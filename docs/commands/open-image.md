# Render Image In Terminal (`open-image`)

> Renders bitmap image data on your terminal

## Description

`open-image` is an optional builtin which will render images (JPEG, GIF,
PNG, BMP, TIFF and WebP) to the terminal using block characters and ANSI
colour sequences.

## Usage

```
<stdin> -> open-image -> <stdout>

open-image file-path -> <stdout>
```

## Examples

### As a method

```
» cat example.png -> open-image
```

### As a function

```
» open-image example.png
```

## Detail

`open-image` will fail if stdout is not a TTY.

## Synonyms

* `open-image`


## See Also

* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [`qr`](../optional/qr.md):
  Creates a QR code from stdin

<hr/>

This document was generated from [builtins/core/openimage/open-image_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/openimage/open-image_doc.yaml).