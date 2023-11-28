package hintsummary

import "os"

func readlink(path string) string {
	ln, err := os.Readlink(path)
	if err != nil {
		return path
	}

	return ln
}
