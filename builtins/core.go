package builtins

// This is where you can add or remove built in functions. Imports here require underscoring.
// Each builtin package should include a README.md with details about its use. However the code should also be readable
// so take a look through the .go files if you're still undecided about whether to include a builtin or not.
//
// My recommendation is to keep everything in this file included, but disable any of the files marked `optional`.
import (
	_ "github.com/lmorg/murex/builtins/core/events"
	_ "github.com/lmorg/murex/builtins/core/httpclient"
	_ "github.com/lmorg/murex/builtins/core/io"
	_ "github.com/lmorg/murex/builtins/core/management"
	_ "github.com/lmorg/murex/builtins/core/mkarray"
	_ "github.com/lmorg/murex/builtins/core/structs"
	_ "github.com/lmorg/murex/builtins/core/textmanip"
	_ "github.com/lmorg/murex/builtins/core/typemgmt"
	_ "github.com/lmorg/murex/builtins/types/binary"
	_ "github.com/lmorg/murex/builtins/types/generic"
	_ "github.com/lmorg/murex/builtins/types/json"
	_ "github.com/lmorg/murex/builtins/types/string"
)
