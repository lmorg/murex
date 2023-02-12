package io

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/tty"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/humannumbers"
)

func init() {
	lang.DefineMethod("pt", cmdPipeTelemetry, types.Any, types.Any)
	lang.DefineMethod(">", cmdWriteFile, types.Any, types.Null)
	lang.DefineMethod("fwrite", cmdWriteFile, types.Any, types.Null)
	lang.DefineMethod(">>", cmdAppendFile, types.Any, types.Null)
	lang.DefineMethod("fappend", cmdAppendFile, types.Any, types.Null)
}

func cmdPipeTelemetry(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)
	quit := make(chan bool)
	stats := func() {
		written, _ := p.Stdin.Stats()
		_, read := p.Stdout.Stats()
		tty.Stderr.WriteString(
			fmt.Sprintf("Pipe telemetry: `%s` written %s -> pt -> `%s` read %s (Data type: %s)\n",
				p.Previous.Name.String(),
				humannumbers.Bytes(written),
				p.Next.Name.String(),
				humannumbers.Bytes(read),
				dt),
		)
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-quit:
				return
			default:
				stats()
			}

		}
	}()

	_, err := io.Copy(p.Stdout, p.Stdin)
	quit <- true
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
