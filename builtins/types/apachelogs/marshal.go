package apachelogs

import (
	"github.com/lmorg/apachelogs"
	"github.com/lmorg/murex/lang/proc"
)

func unmarshal(p *proc.Process) (interface{}, error) {
	var log []apachelogs.AccessLine

	p.Stdin.ReadLine(func(b []byte) {
		line, err, _ := apachelogs.ParseAccessLine(string(b))
		if err != nil {
			return
		}

		log = append(log, *line)
	})

	return log, nil
}
