package hintsummary

import "os"

func readlink(path string) string {
	/*f, err := os.Stat(path)
	if err != nil {
		return err.Error()
	}

	if f.Mode()&os.ModeSymlink != 0 {
		return path
	}*/

	ln, err := os.Readlink(path)
	if err != nil {
		return path
	}

	//if ln[0] != consts.PathSlash[0] {
	return ln //path + " => " + ln
	//}
}
