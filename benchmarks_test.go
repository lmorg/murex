package main

import (
	"fmt"
	"testing"

	"github.com/lmorg/murex/lang"
)

func BenchmarkAForeachN(b *testing.B) {
	lang.InitEnv()

	block := fmt.Sprintf(`a [1..%d] -> foreach i { out "iteration $i of %d" }`, b.N, b.N)

	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
	if err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkJaForeachN(b *testing.B) {
	lang.InitEnv()

	block := fmt.Sprintf(`ja [1..%d] -> foreach i { out "iteration $i of %d" }`, b.N, b.N)

	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
	if err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkAForeachForkN(b *testing.B) {
	lang.InitEnv()

	block := fmt.Sprintf(`a [1..%d] -> foreach i { exec printf "iteration $i of %d\n" }`, b.N, b.N)

	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
	if err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkCsvIndexNTimes(b *testing.B) {
	lang.InitEnv()

	block := []rune(`tout csv "murex,foo,bar\n1,2,3\na,b,c\nz,y,x\n" -> [ :foo ]`)

	for i := 0; i < b.N; i++ {
		_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute(block)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkCsvForkNTimes(b *testing.B) {
	lang.InitEnv()

	block := []rune(`tout csv "murex,foo,bar\n1,2,3\na,b,c\nz,y,x\n" -> grep foo`)

	for i := 0; i < b.N; i++ {
		_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute(block)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkForkPipelineNTimes(b *testing.B) {
	lang.InitEnv()

	block := []rune(`exec printf "the\nquick\nbrown\nfox\n" -> tr '[:lower:]' '[:upper:]' -> tr '[:upper:]' '[:lower:]'`)

	for i := 0; i < b.N; i++ {
		_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute(block)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkMathLibsExpr(b *testing.B) {
	lang.InitEnv()

	block := fmt.Sprintf(`a [1..%d] -> foreach i { 1 * %d }`, b.N, b.N)

	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
	if err != nil {
		b.Error(err.Error())
	}
}
