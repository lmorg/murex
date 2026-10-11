package terminal

import (
	"encoding/json"

	"github.com/lmorg/murex/debug"
)

// ContinuationPrompt reports the secondary prompt text to VS Code (experimental).
func ContinuationPrompt(text string) {
	emitVSCode("P", "ContinuationPrompt="+text)
}

// Prompt reports the prompt text to VS Code (implementation-specific).
func Prompt(text string) {
	emitVSCode("P", "Prompt="+text)
}

// SetMark creates a VS Code navigation mark (experimental); id may be empty.
func SetMark(id string, hidden bool) {
	fields := []string{"SetMark"}
	if id != "" {
		fields = append(fields, "Id="+id)
	}
	if hidden {
		fields = append(fields, "Hidden")
	}
	emitVSCode(fields...)
}

// Environment reports only the supplied variables as JSON to VS Code (experimental).
// Callers should explicitly allowlist variables to avoid exposing credentials.
func vscodeEnvironment(variables map[string]string) {
	filtered := make(map[string]string, len(variables))
	for key, value := range variables {
		if key != "VSCODE_NONCE" {
			filtered[key] = value
		}
	}
	payload, err := json.Marshal(filtered)
	if err != nil {
		debug.Log("term: encoding environment:", err)
		return
	}
	emitVSCode(withNonce([]string{"EnvJson", string(payload)})...)
}

// EnvironmentStart begins a VS Code environment update batch (experimental).
func EnvironmentStart(clear bool) {
	clearValue := "0"
	if clear {
		clearValue = "1"
	}
	emitVSCode(withNonce([]string{"EnvSingleStart", clearValue})...)
}

// EnvironmentEntry adds or updates a variable in the pending batch (experimental).
func EnvironmentEntry(key, value string) {
	if key == "VSCODE_NONCE" {
		return
	}
	emitVSCode(withNonce([]string{"EnvSingleEntry", key, value})...)
}

// EnvironmentDelete removes a variable from the pending batch (experimental).
func EnvironmentDelete(key, value string) {
	if key == "VSCODE_NONCE" {
		return
	}
	emitVSCode(withNonce([]string{"EnvSingleDelete", key, value})...)
}

// EnvironmentEnd completes the pending environment batch (experimental).
func EnvironmentEnd() {
	emitVSCode(withNonce([]string{"EnvSingleEnd"})...)
}
