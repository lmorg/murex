package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/functions"
	"github.com/lmorg/murex/test/count"
)

func TestBlockPropertiesCrLf(t *testing.T) {
	count.Tests(t, 1)
	block := []rune("out 1\nout 2\nout 3")
	blk := NewBlock(block)
	err := blk.ParseBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(blk.Functions) != 3 {
		t.Fatalf("expecting 3 functions, instead got %d", len(blk.Functions))
	}

	for i := range blk.Functions {
		if !blk.Functions[i].Properties.NewChain() {
			t.Errorf("function %d != NewChain() <func>", i)
		}

		if blk.Functions[i].Properties != functions.P_NEW_CHAIN {
			t.Errorf("function %d != P_NEW_CHAIN <int>", i)
		}
	}
}

func TestBlockPropertiesCrLfNoParams(t *testing.T) {
	count.Tests(t, 1)
	block := []rune("1\n2\n3")
	blk := NewBlock(block)
	err := blk.ParseBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(blk.Functions) != 3 {
		t.Fatalf("expecting 3 functions, instead got %d", len(blk.Functions))
	}

	for i := range blk.Functions {
		if !blk.Functions[i].Properties.NewChain() {
			t.Errorf("function %d != NewChain() <func>", i)
		}

		if blk.Functions[i].Properties != functions.P_NEW_CHAIN {
			t.Errorf("function %d != P_NEW_CHAIN <int>", i)
		}
	}
}

func TestBlockPropertiesCrLfx2(t *testing.T) {
	count.Tests(t, 1)
	block := []rune("out 1\n\nout 2\n\nout 3")
	blk := NewBlock(block)
	err := blk.ParseBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(blk.Functions) != 3 {
		t.Fatalf("expecting 3 functions, instead got %d", len(blk.Functions))
	}

	for i := range blk.Functions {
		if !blk.Functions[i].Properties.NewChain() {
			t.Errorf("function %d != NewChain() <func>", i)
		}

		if blk.Functions[i].Properties != functions.P_NEW_CHAIN {
			t.Errorf("function %d != P_NEW_CHAIN <int>", i)
		}
	}
}

func TestBlockPropertiesSemiColon(t *testing.T) {
	count.Tests(t, 1)
	block := []rune("out 1;out 2;out 3")
	blk := NewBlock(block)
	err := blk.ParseBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(blk.Functions) != 3 {
		t.Fatalf("expecting 3 functions, instead got %d", len(blk.Functions))
	}

	for i := range blk.Functions {
		if !blk.Functions[i].Properties.NewChain() {
			t.Errorf("function %d != NewChain() <func>", i)
		}

		if blk.Functions[i].Properties != functions.P_NEW_CHAIN {
			t.Errorf("function %d != P_NEW_CHAIN <int>", i)
		}
	}
}

func TestBlockPropertiesSemiColonCrLF(t *testing.T) {
	count.Tests(t, 1)
	block := []rune("out 1;\nout 2;\nout 3")
	blk := NewBlock(block)
	err := blk.ParseBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(blk.Functions) != 3 {
		t.Fatalf("expecting 3 functions, instead got %d", len(blk.Functions))
	}

	for i := range blk.Functions {
		if !blk.Functions[i].Properties.NewChain() {
			t.Errorf("function %d != NewChain() <func>", i)
		}

		if blk.Functions[i].Properties != functions.P_NEW_CHAIN {
			t.Errorf("function %d != P_NEW_CHAIN <int>", i)
		}
	}
}

func TestBlockPropertiesPosixPipe(t *testing.T) {
	count.Tests(t, 1)
	block := []rune("out 1|out 2|out 3")
	blk := NewBlock(block)
	err := blk.ParseBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(blk.Functions) != 3 {
		t.Fatalf("expecting 3 functions, instead got %d", len(blk.Functions))
	}

	i := 0

	if !blk.Functions[i].Properties.NewChain() {
		t.Errorf("function %d != NewChain() <func>", i)
	}

	if !blk.Functions[i].Properties.PipeOut() {
		t.Errorf("function %d != PipeOut() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_NEW_CHAIN|functions.P_PIPE_OUT {
		t.Errorf("function %d != P_NEW_CHAIN|P_PIPE_OUT <int>", i)
	}

	i = 1

	if !blk.Functions[i].Properties.Method() {
		t.Errorf("function %d != Method() <func>", i)
	}

	if !blk.Functions[i].Properties.PipeOut() {
		t.Errorf("function %d != PipeOut() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_METHOD|functions.P_PIPE_OUT {
		t.Errorf("function %d != P_METHOD|P_PIPE_OUT <int>", i)
	}

	i = 2

	if !blk.Functions[i].Properties.Method() {
		t.Errorf("function %d != Method() <func>", i)
	}

	if blk.Functions[i].Properties.PipeOut() {
		t.Errorf("function %d == PipeOut() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_METHOD {
		t.Errorf("function %d != P_METHOD <int>", i)
	}
}

func TestBlockPropertiesArrowPipe(t *testing.T) {
	count.Tests(t, 1)
	block := []rune("out 1->out 2->out 3")
	blk := NewBlock(block)
	err := blk.ParseBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(blk.Functions) != 3 {
		t.Fatalf("expecting 3 functions, instead got %d", len(blk.Functions))
	}

	i := 0

	if !blk.Functions[i].Properties.NewChain() {
		t.Errorf("function %d != NewChain() <func>", i)
	}

	if !blk.Functions[i].Properties.PipeOut() {
		t.Errorf("function %d != PipeOut() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_NEW_CHAIN|functions.P_PIPE_OUT {
		t.Errorf("function %d != P_NEW_CHAIN|P_PIPE_OUT <int>", i)
	}

	i = 1

	if !blk.Functions[i].Properties.Method() {
		t.Errorf("function %d != Method() <func>", i)
	}

	if !blk.Functions[i].Properties.PipeOut() {
		t.Errorf("function %d != PipeOut() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_METHOD|functions.P_PIPE_OUT {
		t.Errorf("function %d != P_METHOD|P_PIPE_OUT <int>", i)
	}

	i = 2

	if !blk.Functions[i].Properties.Method() {
		t.Errorf("function %d != Method() <func>", i)
	}

	if blk.Functions[i].Properties.PipeOut() {
		t.Errorf("function %d == PipeOut() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_METHOD {
		t.Errorf("function %d != P_METHOD <int>", i)
	}
}

func TestBlockPropertiesQuestionMark(t *testing.T) {
	count.Tests(t, 1)
	block := []rune("out 1 ? out 2 ? out 3")
	blk := NewBlock(block)
	err := blk.ParseBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(blk.Functions) != 3 {
		t.Fatalf("expecting 3 functions, instead got %d", len(blk.Functions))
	}

	i := 0

	if !blk.Functions[i].Properties.NewChain() {
		t.Errorf("function %d != NewChain() <func>", i)
	}

	if !blk.Functions[i].Properties.PipeErr() {
		t.Errorf("function %d != PipeErr() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_NEW_CHAIN|functions.P_PIPE_ERR {
		t.Errorf("function %d != P_NEW_CHAIN|P_PIPE_ERR <int>", i)
	}

	i = 1

	if !blk.Functions[i].Properties.Method() {
		t.Errorf("function %d != Method() <func>", i)
	}

	if !blk.Functions[i].Properties.PipeErr() {
		t.Errorf("function %d != PipeErr() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_METHOD|functions.P_PIPE_ERR {
		t.Errorf("function %d != P_METHOD|P_PIPE_ERR <int>", i)
	}

	i = 2

	if !blk.Functions[i].Properties.Method() {
		t.Errorf("function %d != Method() <func>", i)
	}

	if blk.Functions[i].Properties.PipeErr() {
		t.Errorf("function %d == PipeErr() <func>", i)
	}

	if blk.Functions[i].Properties != functions.P_METHOD {
		t.Errorf("function %d != P_METHOD <int>", i)
	}
}
