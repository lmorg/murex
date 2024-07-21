package cache

import (
	"crypto/sha256"
	"encoding/base64"
)

func CreateHash(cmdLine string, block []rune) string {
	code := []byte(string(block))
	toHash := append([]byte(cmdLine), code...)

	hash := sha256.New()
	_, err := hash.Write(toHash)
	if err != nil {
		return cmdLine
	}

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
