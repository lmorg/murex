package io

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/humannumbers"
)

func init() {
	lang.DefineMethod("pt", cmdPipeTelemetry, types.Any, types.Any)
}

func cmdPipeTelemetry(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)
	quit := make(chan bool)
	stats := func() {
		written, _ := p.Stdin.Stats()
		_, read := p.Stdout.Stats()
		os.Stderr.WriteString(
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
