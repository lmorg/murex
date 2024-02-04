package lang

import (
	"fmt"
	"os"
)

func Deprecated(p *Process) {
	message := fmt.Sprintf("!!! WARNING: The builtin `%s` has been deprecated and will be removed in the next release\n!!!        : Module: %s\n!!!        : File:   %s\n!!!        : Line:   %d\n!!!        : Column: %d\n",
		p.Name.String(), p.FileRef.Source.Module, p.FileRef.Source.Filename, p.FileRef.Line, p.FileRef.Column)
	os.Stderr.WriteString(message)
}
