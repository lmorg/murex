// +build windows

package defaults

// DefaultMurexProfile is basically just the contents of the example murex_profile but wrapped up in Go code so it can
// be compiled into the portable executable. This is also done to make things a little more user friendly out of the box
// ie people don't need to create their own ~/.murex_profile nor `source` the file in /examples.
const DefaultMurexProfile string = `

autocomplete set go { [{
    "Flags": [ "build", "clean", "doc", "env", "bug", "fix", "fmt", "generate", "get", "install", "list", "run", "test", "tool", "version", "vet", "help" ]
}] }
`
