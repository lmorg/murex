package io

import (
	"compress/gzip"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
	"time"
)

func init() {
	proc.GoFunctions["text"] = proc.GoFunction{Func: cmdText, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["open"] = proc.GoFunction{Func: cmdOpen, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["pt"] = proc.GoFunction{Func: cmdPipeTelemetry, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdText(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)
	for _, filename := range p.Parameters.StringArray() {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}

		if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
			gz, err := gzip.NewReader(file)
			if err != nil {
				file.Close()
				return err
			}
			_, err = io.Copy(p.Stdout, gz)
			file.Close()
			if err != nil {
				return err
			}

		} else {
			_, err = io.Copy(p.Stdout, file)
			file.Close()
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func cmdOpen(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)
	for _, filename := range p.Parameters.StringArray() {
		file, err := os.Open(filename)
		if err != nil {
			file.Close()
			return err
		}
		_, err = io.Copy(p.Stdout, file)
		if err != nil {
			file.Close()
			return err
		}

		file.Close()
	}

	return nil
}

func cmdPipeTelemetry(p *proc.Process) error {
	p.Stdout.SetDataType(p.Stdin.GetDataType())
	quit := false
	stats := func() {
		written, _ := p.Stdin.Stats()
		_, read := p.Stdout.Stats()
		os.Stderr.WriteString(
			fmt.Sprintf("Pipe telemetry: `%s` written %s -> pt -> `%s` read %s\n",
				p.Previous.Name,
				utils.HumanBytes(written),
				p.Next.Name,
				utils.HumanBytes(read)),
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
