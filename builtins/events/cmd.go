package events

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
)

func init() {
	proc.GoFunctions["event"] = cmdEvent
	proc.GoFunctions["!event"] = cmdUnevent
}

var rxNameInterruptSyntax *regexp.Regexp = regexp.MustCompile(`^([-_a-zA-Z0-9]+)=(.*)$`)

func cmdEvent(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	et, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if events[et] == nil {
		return fmt.Errorf("No event-type known for `%s`.\nRun `runtime --events` to view which events are compiled in.", et)
	}

	nameInterrupt, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	split := rxNameInterruptSyntax.FindAllStringSubmatch(nameInterrupt, 1)
	if len(split) != 1 || len(split[0]) != 3 {
		return errors.New("Invalid syntax: " + nameInterrupt + ". Expected: `name=interrupt`.")
	}

	name := split[0][1]
	interrupt := ansi.ExpandConsts(split[0][2])

	block, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}

	err = events[et].Add(name, interrupt, block)
	return err
}

func cmdUnevent(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	et, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if events[et] == nil {
		return fmt.Errorf("No event-type known for `%s`.\nRun `runtime --events` to view which events are compiled in.", et)
	}

	name, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	err = events[et].Remove(name)
	return err
}
