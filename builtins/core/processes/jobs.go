package processes

import (
	"github.com/lmorg/murex/lang"
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

	b, err := lang.MarshalData(p, dtLine, []any{
		"JobID",
		"FunctionID",
		"State",
		"Background",
		"Process",
		"Parameters",
	})
	if err != nil {
		return err
	}
	err = aw.Write(b)
	if err != nil {
		return err
	}

	for _, job := range lang.Jobs.List() {
		b, err := lang.MarshalData(p, dtLine, []any{
			job.JobId,
			job.Process.Id,
			job.Process.State.String(),
			job.Process.Background.Get(),
			job.Process.Name.String(),
			getParams(job.Process),
		})
		if err != nil {
			return err
		}
		err = aw.Write(b)
		if err != nil {
			return err
		}
	}

	return aw.Close()
}
