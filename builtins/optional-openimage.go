package builtins

// This is an optional package that renders images inside the terminal. It has additional dependencies:
//
//    go get -u golang.org/x/crypto/ssh/terminal
//    go get -u github.com/disintegration/imaging
//    go get -u golang.org/x/image/bmp
//    go get -u golang.org/x/image/tiff
//    go get -u golang.org/x/image/webp
//
// However these are included in the `vendor` directory
import _ "github.com/lmorg/murex/builtins/optional/openimage" // compile optional builtins
