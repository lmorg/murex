package userdic

import (
	"encoding/json"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestGet(t *testing.T) {
	count.Tests(t, 1)

	a := GetSpellcheckUserDic()

	if len(dictionary) != len(a) {
		t.Error("len doesn't match")
	}
}

func TestRead(t *testing.T) {
	count.Tests(t, 1)

	v, err := ReadSpellcheckUserDic()
	if err != nil {
		t.Error(err)
	}

	if len(v.([]string)) != len(dictionary) {
		t.Error("len doesn't match")
	}
}

func TestWrite(t *testing.T) {
	count.Tests(t, 1)

	i := len(dictionary)
	a := append(dictionary, "test")
	b, err := json.Marshal(&a)
	if err != nil {
		t.Error(err)
	}

	err = WriteSpellcheckUserDic(string(b))
	if err != nil {
		t.Error(err)
	}

	if i+1 != len(dictionary) {
		t.Errorf("Before and after len don't match: %d+1.(before) != %d.(after)", i, len(dictionary))
	}
}
