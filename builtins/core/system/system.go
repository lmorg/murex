package system

import (
	"runtime"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/posix"
)

func init() {
	lang.DefineFunction("os", cmdOs, types.String)
	lang.DefineFunction("cpuarch", cmdCpuArch, types.String)
	lang.DefineFunction("cpucount", cmdCpuCount, types.Integer)
}

func cmdOs(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Write([]byte(runtime.GOOS))
		return err
	}

	for _, os := range p.Parameters.StringArray() {
		if os == runtime.GOOS || (os == "posix" && posix.IsPosix()) {
			_, err := p.Stdout.Write(types.TrueByte)
			return err
		}
	}

	p.ExitNum = 1
	_, err := p.Stdout.Write(types.FalseByte)
	return err
}

func cmdCpuArch(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.String)
	_, err = p.Stdout.Write([]byte(runtime.GOARCH))
	return
}

func cmdCpuCount(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Integer)
	_, err = p.Stdout.Write([]byte(strconv.Itoa(runtime.NumCPU())))
	return
}
