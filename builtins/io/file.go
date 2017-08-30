package io

import (
	"compress/gzip"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/data"
	"github.com/lmorg/murex/utils"
	"github.com/mattn/go-sixel"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func init() {
	proc.GoFunctions["text"] = cmdText
	proc.GoFunctions["open"] = cmdOpen
	proc.GoFunctions["img"] = cmdImg
	proc.GoFunctions["pt"] = cmdPipeTelemetry
	proc.GoFunctions[">"] = cmdWriteFile
	proc.GoFunctions[">>"] = cmdAppendFile
}

var rxExt *regexp.Regexp = regexp.MustCompile(`\.([a-zA-Z]+)(\.gz|)$`)

func cmdText(p *proc.Process) error {
	filename, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	var ext string
	match := rxExt.FindAllStringSubmatch(filename, -1)
	if len(match) > 0 && len(match[0]) > 1 {
		ext = strings.ToLower(match[0][1])
	}

	dt := data.GetExtType(ext)
	if dt == "" {
		p.Stdout.SetDataType(types.String)
	} else {
		p.Stdout.SetDataType(dt)
	}

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

func cmdImg(p *proc.Process) error {
	filename, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	enc := sixel.NewEncoder(os.Stdout)
	err = enc.Encode(img)

	return err
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

func cmdWriteFile(p *proc.Process) error {
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

func cmdAppendFile(p *proc.Process) error {
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
