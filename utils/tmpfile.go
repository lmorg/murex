package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/lmorg/murex/utils/consts"
)

type TempFile struct {
	FileName string
	reader   *os.File
}

func NewTempFile(reader io.Reader, ext string) (*TempFile, error) {
	if ext != "" {
		ext = "." + ext
	}

	fileId := strconv.Itoa(time.Now().Nanosecond())

	h := md5.New()
	_, err := h.Write([]byte(fileId))
	if err != nil {
		return nil, err
	}

	name := consts.TempDir + hex.EncodeToString(h.Sum(nil)) + "-" + strconv.Itoa(os.Getpid()) + ext

	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		return nil, err
	}

	tmp := new(TempFile)
	tmp.FileName = name

	return tmp, nil
}

func (tmp *TempFile) Close() {
	if tmp.reader != nil {
		tmp.reader.Close()
	}

	os.Remove(tmp.FileName)
}

func (tmp *TempFile) Read(p []byte) (int, error) {
	if tmp.reader == nil {
		file, err := os.Open(tmp.FileName)
		if err != nil {
			return 0, err
		}
		tmp.reader = file
	}

	return tmp.reader.Read(p)
}
