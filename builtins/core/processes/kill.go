package processes

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("fid-kill", cmdFidKill, types.Null)
	lang.DefineFunction("fid-killall", cmdKillAll, types.Null)
}

func cmdFidKill(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	for i := 0; i < p.Parameters.Len(); i++ {
		fid, err := p.Parameters.Uint32(i)
		if err != nil {
			return err
		}

		process, err := lang.GlobalFIDs.Proc(fid)
		if err != nil {
			return err
		}

		if process.Kill != nil {
			process.Kill()
		} else {
			return fmt.Errorf("fid `%d` cannot be killed. `Kill` method == `nil`", fid)
		}
	}

	return nil
}

func cmdKillAll(*lang.Process) error {
	fids := lang.GlobalFIDs.ListAll()
	for _, p := range fids {
		if p.Kill != nil /*&& !p.HasTerminated()*/ {
			procName := p.Name.String()
			procParam, _ := p.Parameters.String(0)
			if procName == "exec" {
				procName = procParam
				procParam, _ = p.Parameters.String(1)
			}
			if len(procParam) > 10 {
				procParam = procParam[:10]
			}
			lang.ShellProcess.Stderr.Write([]byte(fmt.Sprintf("!!! Sending kill signal to fid %d: %s %s !!!\n", p.Id, procName, procParam)))
			p.Kill()
		}
	}

	return nil
}
