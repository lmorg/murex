package mxjson

import "errors"

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
	s   []int
	len int
}

func newPair() *pair {
	p := new(pair)
	p.s = make([]int, 100)
	return p
}

func (p *pair) IsOpen() bool {
	return p.len != 0
}

func (p *pair) Open(pos int) {
	if len(p.s) == p.len {
		p.s = append(p.s, pos)
	} else {
		p.s[p.len] = pos
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
	value   interface{}
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

	var v interface{}
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
		v = []interface{}{}

	case objArrayBoolean:
		v = []bool{}

	case objArrayNumber:
		v = []float64{}

	case objArrayString:
		v = []string{}

	case objArrayArray:
		v = [][]interface{}{}

	case objArrayMap:
		v = []interface{}{}

	case objMap:
		v = make(map[string]interface{})

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
	return no.nest[no.len].objType
}

func (no *nestedObject) SetObjType(objType objectType) {
	no.nest[no.len].objType = objType
}

func (no *nestedObject) SetValue(value interface{}) {
	switch no.nest[no.len].objType {
	case objUndefined:
		panic("Undef object type - We should know what the data type is by now!")

	case objBoolean:
		no.nest[no.len].value = value.(bool)

	case objNumber:
		no.nest[no.len].value = value.(float64)

	case objString:
		no.nest[no.len].value = value.(string)

	case objArrayUndefined:
		//panic("Undef array - We should know what the array is by now!")
		no.nest[no.len].value = append(no.nest[no.len].value.([]interface{}), value)

	case objArrayBoolean:
		no.nest[no.len].value = append(no.nest[no.len].value.([]bool), value.(bool))

	case objArrayNumber:
		no.nest[no.len].value = append(no.nest[no.len].value.([]float64), value.(float64))

	case objArrayString:
		no.nest[no.len].value = append(no.nest[no.len].value.([]string), value.(string))

	case objArrayArray:
		no.nest[no.len].value = append(no.nest[no.len].value.([][]interface{}), value.([]interface{}))

	case objArrayMap:
		no.nest[no.len].value = append(no.nest[no.len].value.([]interface{}), value)

	case objMap:
		no.nest[no.len].value.(map[string]interface{})[no.nest[no.len].key.String()] = value

	default:
		panic("This condition shouldn't arise!")
	}
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
