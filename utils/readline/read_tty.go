//go:build !js
// +build !js

package readline

func read(b []byte) (int, error) {
	return replica.Read(b)
}
