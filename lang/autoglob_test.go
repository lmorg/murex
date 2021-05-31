package lang

import "testing"

func TestAutoGlob(t *testing.T) {
	p := NewTestProcess()
	p.Name.Set("@g")
	p.Parameters.DefineParsed([]string{"cat", "*.go"})

	err := autoGlob(p)
	if err != nil {
		t.Errorf("Error returned from autoGlob: %s", err.Error())
	}

	if p.Name.String() != "cat" {
		t.Errorf("autoGlob didn't set the correct executable name")
		t.Logf("  Expecting: `cat`")
		t.Logf("  Actual:    `%s`", p.Name.String())
	}

	name, err := p.Parameters.String(0)
	if err != nil {
		t.Errorf("p.Parameters.String(0) produced an unexpected err: %s", err.Error())
	}
	if name == "cat" {
		t.Errorf("autoGlob didn't remove `cat` from the parameter stack")
	}
	if p.Parameters.Len() < 3 {
		t.Errorf("autoGlob probably didn't expand the globbed parameters because missing some markdown docs")
	}
}
