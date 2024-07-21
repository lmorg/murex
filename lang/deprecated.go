package lang

import (
	"fmt"
	"os"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/lang/ref"
)

func DeprecatedBuiltin(p *Process) {
	Deprecated(fmt.Sprintf("WARNING: The builtin `%s`", p.Name.String()), p.FileRef)
}

func Deprecated(message string, fileRef *ref.File) {
	fileRef = _fileRefUnknown(fileRef)

	s := fmt.Sprintf("!!! %s has been deprecated and will be removed in the v%d.0 release\n!!!        : Module: %s\n!!!        : File:   %s\n!!!        : Line:   %d\n!!!        : Column: %d\n",
		message, app.Major+1,
		fileRef.Source.Module, fileRef.Source.Filename, fileRef.Line, fileRef.Column)
	os.Stderr.WriteString(s)
}

func _fileRefUnknown(fileRef *ref.File) *ref.File {
	if fileRef == nil {
		fileRef = &ref.File{}
	}

	if fileRef.Source == nil {
		fileRef.Source = &ref.Source{
			Filename: "unknown",
			Module:   "unknown",
			DateTime: time.Now(),
		}
	}

	return fileRef
}
