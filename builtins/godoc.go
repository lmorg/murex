// Package builtins is the gatekeeper to the various modules, additional data types and builtin functions within murex
//
// The builtins are split into several categories:
//     * core          - builtin functions required by murex
//     * events        - event hooks for murex code
//     * imports_build - optional builtins you wish to compile
//     * imports_src   - optional builtins available to compile (these are just include files that need to be copied to `builtins/imports_build` if you wish to compile them)
//     * optional      - builtin functions that might add value to murex but are not required. This is the source files rather than the includes
//     * pipes         - different supported methods for murex pipes
//     * types         - murex data types (marshallers et al)
//
// You can specify which packages to enable by creating a file in this package importing the required builtin.
// Or see one of the existing files for reference (eg core.go)
package builtins
