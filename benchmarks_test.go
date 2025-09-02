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

func BenchmarkAForeachNoVarN(b *testing.B) {
	lang.InitEnv()

	block := fmt.Sprintf(`a [1..%d] -> foreach ! { out "iteration (unknown) of %d" }`, b.N, b.N)

	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
	if err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkAForeachDotVarN(b *testing.B) {
	lang.InitEnv()

	block := fmt.Sprintf(`a [1..%d] -> foreach ! { out "$.i of %d" }`, b.N, b.N)

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

func BenchmarkAForeachParallel8Empty(b *testing.B) {
    lang.InitEnv()

    block := fmt.Sprintf(`a [1..%d] -> foreach i --parallel 8 --unordered { out $i }`, b.N)

    _, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
    if err != nil {
        b.Error(err.Error())
    }
}

func BenchmarkAForeachParallel8Out(b *testing.B) {
    lang.InitEnv()

    block := fmt.Sprintf(`a [1..%d] -> foreach i --parallel 8 --unordered { out $i }`, b.N)

    _, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
    if err != nil {
        b.Error(err.Error())
    }
}

func BenchmarkAForeachParallelOrdered(b *testing.B) {
    lang.InitEnv()

    block := fmt.Sprintf(`a [1..%d] -> foreach i --parallel 8 --ordered { out $i }`, b.N)

    _, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
    if err != nil {
        b.Error(err.Error())
    }
}

// Table-driven benchmarks to compare ordered vs unordered across sizes and parallelism.
func BenchmarkForeachParallelMatrix(b *testing.B) {
    sizes := []int{1000, 10000, 100000}
    pars := []int{1, 2, 4, 8, 16}

    for _, size := range sizes {
        for _, par := range pars {
            for _, ordered := range []bool{true, false} {
                name := fmt.Sprintf("size=%d/parallel=%d/%s", size, par, map[bool]string{true: "ordered", false: "unordered"}[ordered])
                b.Run(name, func(b *testing.B) {
                    lang.InitEnv()
                    orderFlag := "--ordered"
                    if !ordered {
                        orderFlag = "--unordered"
                    }
                    block := fmt.Sprintf(`a [1..%d] -> foreach i --parallel %d %s { out $i }`, size, par, orderFlag)
                    b.ResetTimer()
                    for i := 0; i < b.N; i++ {
                        _, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
                        if err != nil {
                            b.Fatal(err)
                        }
                    }
                })
            }
        }
    }
}

// External exec case to stress OS process scheduling under parallel foreach.
func BenchmarkForeachParallelExecTrue(b *testing.B) {
    sizes := []int{200, 1000}
    pars := []int{1, 4, 8}
    for _, size := range sizes {
        for _, par := range pars {
            for _, ordered := range []bool{true, false} {
                name := fmt.Sprintf("exec/size=%d/parallel=%d/%s", size, par, map[bool]string{true: "ordered", false: "unordered"}[ordered])
                b.Run(name, func(b *testing.B) {
                    lang.InitEnv()
                    orderFlag := "--ordered"
                    if !ordered {
                        orderFlag = "--unordered"
                    }
                    block := fmt.Sprintf(`a [1..%d] -> foreach i --parallel %d %s { exec true }`, size, par, orderFlag)
                    b.ResetTimer()
                    for i := 0; i < b.N; i++ {
                        _, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
                        if err != nil {
                            b.Fatal(err)
                        }
                    }
                })
            }
        }
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
