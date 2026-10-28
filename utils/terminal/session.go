package terminal

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"sync"

	"github.com/lmorg/murex/debug"
)

var (
	vscodeNonce = captureNonce()
	sessionOnce sync.Once
	sessionID   string
)

func captureNonce() string {
	nonce := os.Getenv("VSCODE_NONCE")
	if err := os.Unsetenv("VSCODE_NONCE"); err != nil {
		debug.Log("terminal: clearing VS Code nonce:", err)
	}
	return nonce
}

func shellSessionID() string {
	sessionOnce.Do(func() {
		var identifier [16]byte
		if _, err := rand.Read(identifier[:]); err != nil {
			debug.Log("terminal: generating session identifier:", err)
			return
		}
		sessionID = hex.EncodeToString(identifier[:])
	})
	return sessionID
}

func withSessionID(fields []string) []string {
	if identifier := shellSessionID(); identifier != "" {
		fields = append(fields, "aid="+identifier)
	}
	return fields
}
