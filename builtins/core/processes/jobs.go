package processes

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
)

func cmdJobs(p *lang.Process) error {
	var dt, dtLine string
	if p.Stdout.IsTTY() {
		dt = types.Generic
		dtLine = types.Columns
	} else {
		dt = types.JsonLines
		dtLine = types.Json
	}
	p.Stdout.SetDataType(dt)

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	//if p.Stdout.IsTTY() {
	b, err := lang.MarshalData(p, dtLine, []interface{}{
		"PID",
		"State",
		"Background",
		"Process",
		"Parameters",
	})
	if err != nil {
		return err
	}
	_, err = p.Stdout.Writeln(b)
	if err != nil {
		return err
	}
	//}

	procs := lang.GlobalFIDs.ListAll()
	for i := range procs {
		if procs[i].IsBackground || procs[i].State == state.Stopped {
			b, err := lang.MarshalData(p, dtLine, []interface{}{
				procs[i].Id,
				procs[i].State.String(),
				procs[i].IsBackground,
				procs[i].Name,
				getParams(&procs[i]),
			})
			if err != nil {
				return err
			}
			err = aw.Write(b)
			if err != nil {
				return err
			}
		}
	}

	return aw.Close()
}
