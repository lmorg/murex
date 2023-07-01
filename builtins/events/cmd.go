package events

import (
	"fmt"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
)

func init() {
	lang.DefineFunction("event", cmdEvent, types.Null)
	lang.DefineFunction("!event", cmdUnevent, types.Null)
}

var rxNameInterruptSyntax = regexp.MustCompile(`^([-_a-zA-Z0-9]+)=(.*)$`)

func cmdEvent(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	et, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if events[et] == nil {
		return fmt.Errorf("no event-type known for `%s`.\nRun `runtime --events` to view which events are compiled in", et)
	}

	nameInterrupt, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	split := rxNameInterruptSyntax.FindAllStringSubmatch(nameInterrupt, 1)
	if len(split) != 1 || len(split[0]) != 3 {
		return fmt.Errorf("invalid syntax: `%s`. Expected: `name=interrupt`", nameInterrupt)
	}

	name := split[0][1]
	interrupt := ansi.ExpandConsts(split[0][2])

	block, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}

	err = events[et].Add(name, interrupt, block, p.FileRef)
	return err
}

func cmdUnevent(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	et, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if events[et] == nil {
		return fmt.Errorf("no event-type known for `%s`.\nRun `runtime --events` to view which events are compiled in", et)
	}

	name, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	err = events[et].Remove(name)
	return err
}
