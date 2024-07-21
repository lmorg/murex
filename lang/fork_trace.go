//go:build trace
// +build trace

package lang

func forkCheckForNils(fork *Fork) {
	switch {
	case fork.FileRef == nil:
		panic("fork.FileRef == nil in (fork *Fork).Execute()")
	case fork.FileRef.Source == nil:
		panic("fork.FileRef.Source == nil in (fork *Fork).Execute()")
	case fork.FileRef.Source.Module == "":
		panic("missing module name in (fork *Fork).Execute()")
	case fork.Name.String() == "":
		panic("missing function name in (fork *Fork).Execute()")
	}
}
