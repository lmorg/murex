package terminal

import (
	"fmt"
	"strings"
)

// vscodeCommandLine reports the exact command using the captured VS Code nonce, if available.
// Call immediately before CommandBegin when rich command detection is enabled.
func vscodeCommandLine(command string) {
	if command == "" {
		return
	}
	emitVSCode(withNonce([]string{"E", command})...)
}

// WindowsBackend reports the terminal backend type to VS Code.
func WindowsBackend(enabled bool) {
	emitVSCode("P", "IsWindows="+protocolBool(enabled))
}

// RichCommandDetection reports whether command boundaries and command lines are reliable.
func RichCommandDetection(enabled bool) {
	emitVSCode("P", "HasRichCommandDetection="+protocolBool(enabled))
}

func protocolBool(enabled bool) string {
	if enabled {
		return "True"
	}
	return "False"
}

func escapeOSC(value string) string {
	var escaped strings.Builder
	for _, character := range []byte(value) {
		switch {
		case character == '\\':
			escaped.WriteString("\\\\")
		case character <= 0x20 || character == ';' || character == 0x7f:
			fmt.Fprintf(&escaped, "\\x%02x", character)
		default:
			escaped.WriteByte(character)
		}
	}
	return escaped.String()
}

func emitVSCode(fields ...string) {
	switch Application() {
	case TERM_VSCODE:
		for index := range fields {
			fields[index] = escapeOSC(fields[index])
		}
		emitOSC(633, fields...)
	default:
		return
	}
}

func withNonce(fields []string) []string {
	if vscodeNonce != "" {
		fields = append(fields, vscodeNonce)
	}
	return fields
}
