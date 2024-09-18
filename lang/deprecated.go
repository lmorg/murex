package lang

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/lang/ref"
)

func FeatureDeprecatedBuiltin(p *Process) {
	FeatureDeprecated(fmt.Sprintf("The builtin `%s`", p.Name.String()), p.FileRef)
}

func FeatureDeprecated(message string, fileRef *ref.File) {
	s := fmt.Sprintf("%s has been deprecated and will be removed in the v%d.0 release",
		message, app.Major+1)

	FeatureWarning(s, fileRef)
}

func FeatureWarning(message string, fileRef *ref.File) {
	fileRef = _fileRefUnknown(fileRef)

	message = strings.ReplaceAll(message, "\n", "\n!!!        > ")

	s := fmt.Sprintf("!!! WARNING: %s\n!!!        : Module: %s\n!!!        : File:   %s\n!!!        : Line:   %d\n!!!        : Column: %d\n",
		message,
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
