package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
)

type parseObjectKvT struct {
	Interface any
	IsNull    bool
	Runes     []rune
	ValueSet  bool
}

func (kv *parseObjectKvT) Value() any {
	if kv.Interface != nil || kv.IsNull {
		return kv.Interface
	}

	return string(kv.Runes)
}

type parseObjectT struct {
	keyValue [2]parseObjectKvT
	stage    objStageT
	obj      map[string]interface{}
	tree     *ParserT
}

type objStageT int

const (
	OBJ_STAGE_KEY   objStageT = 0
	OBJ_STAGE_VALUE objStageT = 1
)

func newParseObjectT(tree *ParserT) *parseObjectT {
	o := new(parseObjectT)
	o.obj = make(map[string]interface{})
	o.tree = tree
	return o
}

func (o *parseObjectT) WriteKeyValuePair() error {
	ko, vo := o.IsKeyUndefined(), o.IsValueUndefined()

	if ko && vo {
		return nil // empty
	}

	if ko {
		return raiseError(o.tree.expression, nil, o.tree.charPos,
			fmt.Sprintf("object keys cannot be null no undefined. Key/Value expected before '%s' (pos: %d)",
				string(o.tree.expression[o.tree.charPos]), o.tree.charPos))
	}

	if vo {
		return raiseError(o.tree.expression, nil, o.tree.charPos,
			fmt.Sprintf("object values cannot be undefined. Value expected before '%s' (pos: %d)",
				string(o.tree.expression[o.tree.charPos]), o.tree.charPos))
	}

	key, err := types.ConvertGoType(o.keyValue[OBJ_STAGE_KEY].Value(), types.String)
	if err != nil {
		return err
	}

	value := o.keyValue[OBJ_STAGE_VALUE].Value()

	o.obj[key.(string)] = value

	o.keyValue = [2]parseObjectKvT{}
	o.stage = 0

	return nil
}

const (
	objErrUnexpectedKey   = "unexpected object key:\n* new values should follow a colon\n* barewords cannot include whitespace"
	objErrUnexpectedValue = "unexpected object value:\n* new keys should follow a comma, or new line.\n* objects should be terminated with a closing curly bracket '}'\n* barewords cannot include whitespace"
)

func (o *parseObjectT) AppendRune(r ...rune) error {
	if o.ExpectNextStage() {
		if o.stage == OBJ_STAGE_KEY {
			return raiseError(o.tree.expression, nil, o.tree.charPos, objErrUnexpectedKey)
		}

		return raiseError(o.tree.expression, nil, o.tree.charPos, objErrUnexpectedValue)
	}

	o.keyValue[o.stage].Runes = append(o.keyValue[o.stage].Runes, r...)
	return nil
}

func (o *parseObjectT) UpdateInterface(v interface{}) error {
	if o.ExpectNextStage() {
		if o.stage == OBJ_STAGE_KEY {
			return raiseError(o.tree.expression, nil, o.tree.charPos, objErrUnexpectedKey)
		}

		return raiseError(o.tree.expression, nil, o.tree.charPos, objErrUnexpectedValue)
	}

	if v == nil {
		o.keyValue[o.stage].IsNull = true
	} else {
		o.keyValue[o.stage].Interface = v
	}
	o.keyValue[o.stage].ValueSet = true

	return nil
}

func (o *parseObjectT) IsKeyUndefined() bool {
	return o.keyValue[OBJ_STAGE_KEY].Interface == nil &&
		len(o.keyValue[OBJ_STAGE_KEY].Runes) == 0
}

func (o *parseObjectT) IsValueUndefined() bool {
	return o.keyValue[OBJ_STAGE_VALUE].Interface == nil &&
		!o.keyValue[OBJ_STAGE_VALUE].IsNull &&
		len(o.keyValue[OBJ_STAGE_VALUE].Runes) == 0
}

func (o *parseObjectT) ExpectNextStage() bool {
	return o.keyValue[o.stage].ValueSet
}
