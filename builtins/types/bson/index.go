package bson

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
	"labix.org/v2/mgo/bson"
)

func readIndex(p *proc.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = bson.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	return define.IndexTemplateObject(p, params, &jInterface, bson.Marshal)
}
