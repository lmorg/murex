package cache

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func CreateHash(exe string, args []string, block []rune) string {
	argv := exe + " " + strings.Join(args, " ")
	code := []byte(string(block))

	hash := sha256.New()
	_, err := hash.Write(append([]byte(argv), code...))
	if err != nil {
		return argv
	}

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
