package io

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func init() {
	lang.GoFunctions["tmp"] = cmdTempFile
}

func cmdTempFile(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	ext, _ := p.Parameters.String(0)
	if ext != "" {
		ext = "." + ext
	}

	fileId := time.Now().String() + ":" + strconv.Itoa(int(p.Id))

	h := md5.New()
	_, err := h.Write([]byte(fileId))
	if err != nil {
		return err
	}

	name := consts.TempDir + hex.EncodeToString(h.Sum(nil)) + "-" + strconv.Itoa(os.Getpid()) + ext

	file, err := os.Create(name)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, p.Stdin)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(name))
	return err
}
