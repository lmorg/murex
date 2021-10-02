// +build !js

package readline

import "os"

func read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}
