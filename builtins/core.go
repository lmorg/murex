package builtins

// This is where you can add or remove built in functions. Imports here require underscoring.
// Each builtin package should include a README.md with details about its use. However the code should also be readable
// so take a look through the .go files if you're still undecided about whether to include a builtin or not.
//
// My recommendation is to keep everything in this file included, but disable any of the files marked `optional`.

import (
	_ "github.com/lmorg/murex/builtins/encoders"
	_ "github.com/lmorg/murex/builtins/httpclient"
	_ "github.com/lmorg/murex/builtins/io"
	_ "github.com/lmorg/murex/builtins/management"
	_ "github.com/lmorg/murex/builtins/mkarray"
	_ "github.com/lmorg/murex/builtins/structs"
	_ "github.com/lmorg/murex/builtins/textmanip"
	_ "github.com/lmorg/murex/builtins/time"
	_ "github.com/lmorg/murex/builtins/typemgmt"
)
