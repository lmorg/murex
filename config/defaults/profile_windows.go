// +build windows

package defaults

func init() {
	murexProfile = append(murexProfile, `

autocomplete set go { [{
    "Flags": [ "build", "clean", "doc", "env", "bug", "fix", "fmt", "generate", "get", "install", "list", "run", "test", "tool", "version", "vet", "help" ]
}] }
`)
}
