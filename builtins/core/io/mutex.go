// +build ignore

package io

import (
	"errors"
	"sync"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

var mutexes map[string]sync.Mutex

func init() {
	proc.GoFunctions["mutex"] = cmdMutex

	mutexes = make(map[string]murexMutex)
}

func cmdMutex(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	method, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	name, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	switch method {
	case "lock":
		mutexes[name].Lock()

	case "unlock":
		mutexes[name].Unlock()

	case "wait":
		mutexes[name].Lock()
		mutexes[name].Unlock()

	default:
		return errors.New("That isn't a valid parameter for mutex.")
	}

	return nil
}
