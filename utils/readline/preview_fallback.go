//go:build plan9 || windows || js
// +build plan9 windows js

package readline

func previewFile(filename string) []byte {
	return nil
}
