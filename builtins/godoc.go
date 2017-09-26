// Package builtins is the gatekeeper to the various modules, additional data types and builtin functions within murex
//
// The builtins are split into several categories:
//     * core - builtin functions required by murex
//     * optional - builtin functions that might add value to murex but are not required
//     * preview - interfaces for rendering various document types within the CLI
//     * stdio - pipe interfaces (eg stdout, network pipes, file writers, etc)
//     * types - murex data types (marshallers et al)
//
// You can specify which packages to enable by creating a file in this package importing the required builtin.
// Or see one of the existing files for reference (eg core.go)
package builtins
