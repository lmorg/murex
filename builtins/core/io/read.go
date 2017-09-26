package io

import (
	"github.com/lmorg/murex/lang/proc"
)

func init() {
	proc.GoFunctions["read"] = cmdRead
}

func cmdRead(p *proc.Process) (err error) {

	return
}
