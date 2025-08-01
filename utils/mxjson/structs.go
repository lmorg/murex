package mxjson

import (
	"errors"
	"fmt"
)

// parser state

type parserState int

const (
	stateBeginKey parserState = 0 + iota
	stateEndKey
	stateBeginVal
	stateEndVal
	stateExpectNext
)

type objectType int

const (
	objUndefined objectType = 0 + iota
	objBoolean
	objNumber
	objString
	objArrayUndefined
	objArrayBoolean
	objArrayNumber
	objArrayString
	objArrayArray
	objArrayMap
	objMap
)

// quote pairs

type quote struct {
	open bool
	pos  int
}

func (q *quote) IsOpen() bool {
	return q.open
}

func (q *quote) Open(pos int) {
	q.open = true
	q.pos = pos
}

func (q *quote) Close() {
	q.open = false
}

// bracket pairs

type pair struct {
	pos []int
	len int
}

func newPair() *pair {
	p := new(pair)
	p.pos = make([]int, 100)
	return p
}

func (p *pair) IsOpen() bool {
	return p.len != 0
}

func (p *pair) Open(pos int) {
	if len(p.pos) == p.len {
		p.pos = append(p.pos, pos)
	} else {
		p.pos[p.len] = pos
	}
	p.len++
}

func (p *pair) Close() error {
	if p.len == 0 {
		return errors.New("close found with no open")
	}
	p.len--
	return nil
}

// lazy string

type str struct {
	b   []byte
	len int
}

func newStr() *str {
	s := new(str)
	s.b = make([]byte, 1024*1024)
	return s
}

func (s *str) Append(b byte) {
	if len(s.b) == s.len {
		s.b = append(s.b, b)
	} else {
		s.b[s.len] = b
	}
	s.len++
}

func (s *str) Get() []byte {
	if s.len == 0 {
		return nil
	}

	b := s.b[:s.len]
	s.len = 0
	return b
}

func (s *str) String() string {
	if s.len == 0 {
		return ""
	}

	b := s.b[:s.len]
	s.len = 0
	return string(b)
}

// object type

type nestedObject struct {
	nest []item
	len  int
}
type item struct {
	key     *str
	value   any
	objType objectType
}

func newObjs() (no nestedObject) {
	no.nest = make([]item, 10)
	for i := range no.nest {
		no.nest[i].key = newStr()
	}
	no.len = -1
	return
}

func (no *nestedObject) New(objType objectType) {
	no.len++

	var v any
	switch objType {
	case objUndefined:
		panic("Undef object type - We should know what the data type is by now!")

	case objBoolean:
		v = false

	case objNumber:
		v = float64(0)

	case objString:
		v = ""

	case objArrayUndefined:
		//panic("Undef array - We should know what the array is by now!")
		v = []any{}

	case objArrayBoolean:
		v = []bool{}

	case objArrayNumber:
		v = []float64{}

	case objArrayString:
		v = []string{}

	case objArrayArray:
		v = [][]any{}

	case objArrayMap:
		v = []any{}

	case objMap:
		v = make(map[string]any)

	default:
		panic("This condition shouldn't arise!")
	}

	if len(no.nest) == no.len {
		no.nest = append(no.nest, item{
			key:     newStr(),
			objType: objType,
			value:   v,
		})
	} else {
		no.nest[no.len].objType = objType
		no.nest[no.len].value = v
	}
}

func (no *nestedObject) GetKeyPtr() *str {
	return no.nest[no.len].key
}

func (no *nestedObject) GetObjType() objectType {
	if no.len < 0 || len(no.nest) == 0 || no.len > len(no.nest) {
		return objUndefined
	}
	return no.nest[no.len].objType
}

func (no *nestedObject) SetObjType(objType objectType) {
	no.nest[no.len].objType = objType
}

func (no *nestedObject) SetValue(value any) error {
	if no.len < 0 {
		return fmt.Errorf("unable to marshal '%v' into parent structure. This might be due to the an incorrect file", value)
	}

	switch no.nest[no.len].objType {
	case objUndefined:
		return fmt.Errorf("undef object type. We should know what the data type is by now. Please file a bug at https://github.com/lmorg/murex/issues")

	case objBoolean:
		no.nest[no.len].value = value.(bool)

	case objNumber:
		no.nest[no.len].value = value.(float64)

	case objString:
		no.nest[no.len].value = value.(string)

	case objArrayUndefined:
		//panic("Undef array - We should know what the array is by now!")
		no.nest[no.len].value = append(no.nest[no.len].value.([]any), value)

	case objArrayBoolean:
		no.nest[no.len].value = append(no.nest[no.len].value.([]bool), value.(bool))

	case objArrayNumber:
		no.nest[no.len].value = append(no.nest[no.len].value.([]float64), value.(float64))

	case objArrayString:
		no.nest[no.len].value = append(no.nest[no.len].value.([]string), value.(string))

	case objArrayArray:
		no.nest[no.len].value = append(no.nest[no.len].value.([][]any), value.([]any))

	case objArrayMap:
		no.nest[no.len].value = append(no.nest[no.len].value.([]any), value)

	case objMap:
		no.nest[no.len].value.(map[string]any)[no.nest[no.len].key.String()] = value

	default:
		return fmt.Errorf("switch statement unexpectedly failed. This condition shouldn't arise, so please file a bug at https://github.com/lmorg/murex/issues")
	}

	return nil
}

func (no *nestedObject) MergeDown() {
	switch no.len {
	case -1:
		panic("This condition shouldn't arise!")

	case 0:
		no.len--

	default:
		no.len--
		no.SetValue(no.nest[no.len+1].value)
	}
}
