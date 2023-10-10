package management

import (
	"fmt"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineMethod("debug", cmdDebug, types.Any, types.Json)
}

func cmdDebug(p *lang.Process) error {
	if p.IsMethod {
		return cmdDebugMethod(p)
	}

	block, err := p.Parameters.Block(0)
	if err == nil {
		return cmdDebugBlock(p, block)
	}

	p.Stdout.SetDataType(types.Boolean)

	v, err := p.Parameters.Bool(0)

	if err != nil {
		_, err = p.Stdout.Write([]byte(fmt.Sprint(debug.Enabled)))
		return err
	}
	debug.Enabled = v
	if !v {
		p.Stdout.Writeln(types.FalseByte)
		p.ExitNum = 1
		return nil
	}

	_, err = p.Stdout.Writeln(types.TrueByte)
	return err
}

func cmdDebugMethod(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(types.Json)

	var (
		j = make(map[string]interface{})
		b []byte
	)

	obj, err := lang.UnmarshalData(p, dt)

	j["Process"] = p.Previous.Dump()
	j["Data-Type"] = map[string]string{
		"Murex":             dt,
		"Go":                fmt.Sprintf("%T", obj),
		"UnmarshalData Err": fmt.Sprint(err),
	}

	b, err = json.Marshal(j, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdDebugBlock(p *lang.Process, block []rune) (err error) {
	v := debug.Enabled
	debug.Enabled = true
	defer func() {
		debug.Enabled = v
	}()

	fork := p.Fork(lang.F_NO_STDIN)
	p.ExitNum, err = fork.Execute(block)
	return
}
