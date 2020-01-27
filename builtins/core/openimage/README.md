# Package `openimage`

This renders bitmap image data on your terminal. It contains the following commands:

* `open-image`

As well as defines extensions and mime-types for the following murex data-types:

* `image`

It has the following additional dependencies:

    go get -u golang.org/x/crypto/ssh/terminal
    go get -u github.com/disintegration/imaging
    go get -u golang.org/x/image/bmp
    go get -u golang.org/x/image/tiff
    go get -u golang.org/x/image/webp

However these are included in the `vendor` directory.
