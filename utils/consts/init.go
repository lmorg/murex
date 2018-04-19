package consts

import "os"

func init() {
	createDirIfNotExist(TempDir)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		//if err != nil {
		//	panic(err)
		//}
	}
}
