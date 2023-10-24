//go:build opt_type_hcl
// +build opt_type_hcl

package optional

// Uses a 3rd party library: github.com/hashicorp/hcl
// (included in vendor directory)
import _ "github.com/lmorg/murex/builtins/types/hcl" // compile data type
