package io

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

func init() {
	lang.GoFunctions["pt"] = cmdPipeTelemetry
	lang.GoFunctions[">"] = cmdWriteFile
	lang.GoFunctions[">>"] = cmdAppendFile
	lang.GoFunctions["tmp"] = cmdTempFile
}

func cmdPipeTelemetry(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)
	quit := false
	stats := func() {
		written, _ := p.Stdin.Stats()
		_, read := p.Stdout.Stats()
		os.Stderr.WriteString(
			fmt.Sprintf("Pipe telemetry: `%s` written %s -> pt -> `%s` read %s (Data type: %s)\n",
				p.Previous.Name,
				utils.HumanBytes(written),
				p.Next.Name,
				utils.HumanBytes(read),
				dt),
		)
	}

	go func() {
		for !quit {
			time.Sleep(1 * time.Second)
			if quit {
				return
			}
			stats()
		}
	}()

	_, err := io.Copy(p.Stdout, p.Stdin)
	quit = true
	stats()
	return err
}

func cmdWriteFile(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, p.Stdin)
	return err
}

func cmdAppendFile(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, p.Stdin)
	return err
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
