package builtins

// This is where you can add or remove built in functions. Imports here require underscoring.
// Each builtin package should include a README.md with details about its use. However the code should also be readable
// so take a look through the .go files if you're still undecided about whether to include a builtin or not.
//
// My recommendation is to keep everything in this file included, but disable any of the files marked `optional`.
import (
	_ "github.com/lmorg/murex/builtins/core/arange"     // working with ranges within arrays (`@[..]`)
	_ "github.com/lmorg/murex/builtins/core/datatools"  // utilities for manipulating structured data
	_ "github.com/lmorg/murex/builtins/core/events"     // rudimentary event system
	_ "github.com/lmorg/murex/builtins/core/httpclient" // builtins for http
	_ "github.com/lmorg/murex/builtins/core/io"         // OS IO builtins
	_ "github.com/lmorg/murex/builtins/core/management" // murex management builtins
	_ "github.com/lmorg/murex/builtins/core/mkarray"    // array management builtins
	_ "github.com/lmorg/murex/builtins/core/random"     // random data builtin
	_ "github.com/lmorg/murex/builtins/core/structs"    // control structures
	_ "github.com/lmorg/murex/builtins/core/textmanip"  // text manipulation builtins
	_ "github.com/lmorg/murex/builtins/core/typemgmt"   // type handling and management builtins
	_ "github.com/lmorg/murex/builtins/types/binary"    // basic data type for handing binary data
	_ "github.com/lmorg/murex/builtins/types/generic"   // generic (string) data type
	_ "github.com/lmorg/murex/builtins/types/json"      // JSON data type
	_ "github.com/lmorg/murex/builtins/types/string"    // string data type (soon to be depreciated)
)
