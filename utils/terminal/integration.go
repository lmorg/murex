package terminal

import (
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/utils/envvars"
)

type PromptKind string

const (
	PromptPrimary      PromptKind = "i"
	PromptSecondary    PromptKind = "s"
	PromptRight        PromptKind = "r"
	PromptContinuation PromptKind = "c"
)

type PromptOptions struct {
	Kind PromptKind
}

// PromptBegin marks a prompt's start using the shell's session identifier for OSC 133.
// Prompt kinds and session identifiers use implementation-specific OSC 133 extensions.
func PromptBegin(options ...PromptOptions) {
	var option PromptOptions
	if len(options) > 0 {
		option = options[0]
	}
	switch option.Kind {
	case "", PromptPrimary, PromptSecondary, PromptRight, PromptContinuation:
	default:
		debug.Log("term: invalid prompt kind")
		return
	}

	switch Application() {
	case TERM_VSCODE:
		switch option.Kind {
		case PromptSecondary, PromptContinuation:
			emitOSC(633, "F")
		case PromptRight:
			emitOSC(633, "H")
		default:
			emitOSC(633, "A")
		}
	default:
		fields := []string{"A"}
		if option.Kind != "" {
			fields = append(fields, "k="+string(option.Kind))
		}
		emitOSC(133, withSessionID(fields)...)
	}
}

// PromptEnd marks a prompt's end; the optional kind selects VS Code's G/I markers.
func PromptEnd(kind ...PromptKind) {
	var promptKind PromptKind
	if len(kind) > 0 {
		promptKind = kind[0]
	}
	switch promptKind {
	case "", PromptPrimary, PromptSecondary, PromptRight, PromptContinuation:
	default:
		debug.Log("term: invalid prompt kind")
		return
	}
	switch Application() {
	case TERM_VSCODE:
		switch promptKind {
		case PromptSecondary, PromptContinuation:
			emitOSC(633, "G")
		case PromptRight:
			emitOSC(633, "I")
		default:
			emitOSC(633, "B")
		}
	default:
		emitOSC(133, "B")
	}
}

// RightPromptProperty emits the implementation-specific OSC 133 right-prompt property.
// VS Code has no equivalent property marker, so it emits nothing there.
// Murex also has no support for right prompts
func RightPromptProperty() {
	switch Application() {
	case TERM_VSCODE:
		return
	default:
		emitOSC(133, "P", "k=r")
	}
}

// CommandBegin marks the start of command execution.
func CommandBegin(commandLine string) {
	switch Application() {
	case TERM_VSCODE:
		vscodeCommandLine(commandLine)
		emitOSC(633, "C")
	case TERM_TTYPHOON:
		envVars := make(map[string]any)
		envvars.All(envVars)
		ttyphoonOutputBlockBegin(commandLine, envVars)
	default:
		emitOSC(133, "C")
	}
}

// CommandFinished marks completion. An exit code of -1 omits the status.
// OSC 133 uses the shell's session identifier automatically.
func CommandFinished(exitCode int) {
	fields := []string{"D"}
	if exitCode != -1 {
		fields = append(fields, strconv.Itoa(exitCode))
	}
	switch Application() {
	case TERM_VSCODE:
		emitOSC(633, fields...)
		//vscodeEnvironment()
	case TERM_TTYPHOON:
		ttyphoonOutputBlockEnd(exitCode)
	default:
		emitOSC(133, withSessionID(fields)...)
	}
}

// WorkingDirectory reports a directory via OSC 633 in VS Code or an OSC 7 file URI.
// With no argument it reports the process's current directory.
func WorkingDirectory(path ...string) {
	var directory string
	var err error
	if len(path) > 0 {
		directory, err = filepath.Abs(path[0])
	} else {
		directory, err = os.Getwd()
	}
	if err != nil {
		debug.Log("term: resolving working directory:", err)
		return
	}
	switch Application() {
	case TERM_VSCODE:
		emitVSCode("P", "Cwd="+directory)
	default:
		hostname, err := os.Hostname()
		if err != nil {
			debug.Log("term: resolving hostname:", err)
			return
		}
		uriPath := filepath.ToSlash(directory)
		if len(uriPath) > 0 && uriPath[0] != '/' {
			uriPath = "/" + uriPath
		}
		uri := url.URL{Scheme: "file", Host: hostname, Path: uriPath}
		emitOSC(7, uri.String())
	}
}
