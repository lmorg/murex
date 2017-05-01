package builtins

// This is where you can add or remove built in functions. Imports here require underscoring.
// Each builtin package should include a README.md with details about its use. However the code should also be readable
// so take a look through the .go files if you're still undecided about whether to include a builtin or not.
//
// Shell author's recommendation: keep everything included ;-). If you really feel the need to disable something then
// first take a look at `misc`.

import (
	_ "github.com/lmorg/murex/lang/builtins/encoders"
	_ "github.com/lmorg/murex/lang/builtins/httpclient"
	_ "github.com/lmorg/murex/lang/builtins/io"
	_ "github.com/lmorg/murex/lang/builtins/management"
	_ "github.com/lmorg/murex/lang/builtins/misc"
	_ "github.com/lmorg/murex/lang/builtins/structs"
	_ "github.com/lmorg/murex/lang/builtins/textmanip"
	_ "github.com/lmorg/murex/lang/builtins/typemgmt"
)
