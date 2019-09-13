package cmdtest

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func testRun(p *lang.Process) error {
	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	fork := p.Fork(lang.F_FUNCTION)
	fork.Name = "(test run)"

	err = fork.Config.Set("test", "enabled", true)
	if err != nil {
		return err
	}

	err = fork.Config.Set("test", "auto-report", true)
	if err != nil {
		return err
	}

	h := md5.New()
	_, err = h.Write([]byte(time.Now().String() + ":" + strconv.Itoa(int(p.Id))))
	if err != nil {
		return err
	}

	pipeName := "system_test_" + hex.EncodeToString(h.Sum(nil))

	err = lang.GlobalPipes.CreatePipe(pipeName, "std", "")
	if err != nil {
		return err
	}

	pipe, err := lang.GlobalPipes.Get(pipeName)
	if err != nil {
		return err
	}

	err = fork.Config.Set("test", "report-pipe", pipeName)
	if err != nil {
		return err
	}

	_, err = fork.Execute(block)
	if err != nil {
		return err
	}

	err = lang.GlobalPipes.Close(pipeName)
	if err != nil {
		return err
	}

	reportType, err := p.Config.Get("test", "report-format", types.String)
	if err != nil {
		return err
	}
	if reportType.(string) == "table" {
		p.Stderr.Writeln([]byte(consts.TestTableHeadings))
	}

	_, err = io.Copy(p.Stderr, pipe)
	return err
}
