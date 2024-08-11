package expressions_test

import (
	"embed"
	"fmt"
	"testing"
	"time"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func TestParseBlock(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `out 1;out 2;out 3;out 4;out 5`,
			Stdout: "1\n2\n3\n4\n5\n",
		},
		{
			Block:  `out 1;out 2;out 3;out 4;out 5;`,
			Stdout: "1\n2\n3\n4\n5\n",
		},
		{
			Block:  "out 1\nout 2\nout 3\nout 4\nout 5",
			Stdout: "1\n2\n3\n4\n5\n",
		},
		{
			Block:  "out 1\nout 2\nout 3\nout 4\nout 5\n\n",
			Stdout: "1\n2\n3\n4\n5\n",
		},
		/*{
			Block:  `${err 1|err 2|err 3|err 4|err 5} ? msort`,
			Stdout: "1\n2\n3\n4\n5\n",
		},*/
		{
			Block:  "out:1\nout:2\nout:3\nout:4\nout:5",
			Stdout: "1\n2\n3\n4\n5\n",
		},
	}

	test.RunMurexTests(tests, t)
}

//go:embed testcode/*.mx
var testcode embed.FS

func TestParseBlockExampleRealCode(t *testing.T) {
	dir, err := testcode.ReadDir("testcode")
	if err != nil {
		// not a bug in murex
		panic(err)
	}

	count.Tests(t, len(dir))

	for i := range dir {
		name := dir[i].Name()

		b, err := testcode.ReadFile("testcode/" + name)
		if err != nil {
			// not a bug in murex
			panic(err)
		}

		block := []rune(string(b))
		blk := expressions.NewBlock(block)
		err = blk.ParseBlock()
		if err != nil {
			// this _is_ a bug in murex!
			t.Errorf("testcode failed to parse: `%s`", name)
			t.Logf("  Error returned: %v", err)
		}
	}
}

func TestParseBlockSubBlocks(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `function a {$ARGS};a (${out foo bar},)`,
			Stdout: `["a","foo bar,"]`,
		},
		{
			Block:  `function a {$PARAMS};a { bob }`,
			Stdout: `["{ bob }"]`,
		},
		{
			Block:  `function a {$PARAMS};a { { bob } }`,
			Stdout: `["{ { bob } }"]`,
		},
		{
			Block:  `function a {$PARAMS};a ${ out { { bob } } }`,
			Stdout: `["{ { bob } }"]`,
		},
		{
			Block:  `function a {$PARAMS};a {({({4})})}{({({4})})}`,
			Stdout: `["{({({4})})}{({({4})})}"]`,
		},
		{
			Block:  `function a {$PARAMS};a ${ out {({({5})})}{({({5})})} }`,
			Stdout: `["{({({5})})}{({({5})})}"]`,
		},
		/*{
			Block:  "function a {$PARAMS};a ${\n\nout ({\n(\n{\n(\n{\n5\n}\n)\n}\n)\n}\n{\n(\n{\n(\n{\n5\n}\n)\n}\n)\n}\n\n})",
			Stdout: "[\"{\n(\n{\n(\n{\n5\n}\n)\n}\n)\n}\n{\n(\n{\n(\n{\n5\n}\n)\n}\n)\n}\"]",
		},*/
		{
			Block:  `function a {$PARAMS};a ${ out ${ out ${ out bob } } }`,
			Stdout: `["bob"]`,
		},
		{
			Block:  `function a {$PARAMS};a ${ ${ ${ out bob } } }`,
			Stdout: `["bob"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockExistingCodeBugFixes1(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `config: eval shell safe-commands {
				-> alter --merge / ([
					"builtins", "jobs"
				])
			}`,
			Stdout: ``,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockEscapedCrLf(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				out \
				bob`,
			Stdout: "bob\n",
		},
		/////
		{
			Block: `
				out \ # comment
				bob`,
			Stdout: "bob\n",
		},
		{
			Block: `
				out \	# comment
				bob`,
			Stdout: "bob\n",
		},
		/////
		{
			Block: `
				out # comment \
				bob`,
			Stdout: "bob\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockPipeOverwriteFile1(t *testing.T) {
	filename := fmt.Sprintf("%s-%d.testfile", t.Name(), time.Now().UnixNano())
	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`> %s; open %s; rm %s`, filename, filename, filename),
			Stdout: ``,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockPipeOverwriteFile2(t *testing.T) {
	filename := fmt.Sprintf("%s-%d.testfile", t.Name(), time.Now().UnixNano())
	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`%%(%s) |> %s; open %s; rm %s`, filename, filename, filename, filename),
			Stdout: filename,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockPipeAppendFile1(t *testing.T) {
	filename := fmt.Sprintf("%s-%d.testfile", t.Name(), time.Now().UnixNano())
	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`%%(%s) >> %s; %%(%s) >> %s; open %s; rm %s`, filename, filename, filename, filename, filename, filename),
			Stdout: filename + filename,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockPipeAppendFile2(t *testing.T) {
	filename := fmt.Sprintf("%s-%d.testfile", t.Name(), time.Now().UnixNano())
	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`echo %s >> %s; echo %s >> %s; open %s; rm %s`, filename, filename, filename, filename, filename, filename),
			Stdout: fmt.Sprintf("%s\n%s\n", filename, filename),
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockLogicOperators(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `out 1 && out 2 && false && out 3`,
			Stdout:  "1\n2\n",
			ExitNum: 1,
		},
		{
			Block:   `out 1 && out 2 && false || out 3`,
			Stdout:  "1\n2\n3\n",
			ExitNum: 0,
		},
		//
		{
			Block:   `out 1 || out 2 || false || out 3`,
			Stdout:  "1\n",
			ExitNum: 0,
		},
		{
			Block:   `out 1 || out 2 || false && out 3`,
			Stdout:  "1\n",
			ExitNum: 0,
		},
		/////
		{
			Block:   `out 1 && out 2 && true && out 3`,
			Stdout:  "1\n2\n3\n",
			ExitNum: 0,
		},
		{
			Block:   `out 1 && out 2 && true || out 3`,
			Stdout:  "1\n2\n",
			ExitNum: 0,
		},
		//
		{
			Block:   `out 1 || out 2 || true || out 3`,
			Stdout:  "1\n",
			ExitNum: 0,
		},
		{
			Block:   `out 1 || out 2 || true && out 3`,
			Stdout:  "1\n",
			ExitNum: 0,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockAmpersand(t *testing.T) {
	count.Tests(t, 1)

	block := "&foobar"

	tree, err := expressions.ParseBlock([]rune(block))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err.Error())
	}

	if len(*tree) != 1 {
		t.Fatalf("Invalid tree length: %v", len(*tree))
	}

	if string((*tree)[0].Command) != block {
		t.Error("Invalid tree:")
		t.Logf("  Command:     '%s'", string((*tree)[0].Command))
		t.Logf("  Named Pipes: %v", json.LazyLogging(((*tree)[0].NamedPipes)))
		t.Logf("  Parameters:  %v", ((*tree)[0].Parameters))
	}
}

func TestParseBlockHyphen(t *testing.T) {
	count.Tests(t, 1)

	block := "-abc"

	tree, err := expressions.ParseBlock([]rune(block))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err.Error())
	}

	if len(*tree) != 1 {
		t.Fatalf("Invalid tree length: %v", len(*tree))
	}

	if string((*tree)[0].Command) != block {
		t.Error("Invalid tree:")
		t.Logf("  Command:     '%s'", string((*tree)[0].Command))
		t.Logf("  Named Pipes: %v", json.LazyLogging(((*tree)[0].NamedPipes)))
		t.Logf("  Parameters:  %v", ((*tree)[0].Parameters))
	}
}

func TestParseBlockCommentHash(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				out: 1
				# out 2
				out 3
			`,
			Stdout: "1\n3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseBlockArrayPanicBugFix(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `TestParseBlockArrayPanicBugFix0 = @TestParseBlockArrayPanicBugFix0[0]`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
		{
			Block:   `@TestParseBlockArrayPanicBugFix1[1]`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
		{
			Block:   `echo @TestParseBlockArrayPanicBugFix2[22]`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
		{
			Block:   `echo @3[333]`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
		{
			Block:   `echo @44[44]`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
		{
			Block:   `echo @555[5]`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
		{
			Block:   `echo @6666[]`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestAlterOp(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `tout json %{a:1, b:2, c:3} ~> %{b:4}`,
			Stdout: `{"a":1,"b":4,"c":3}`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestIssue854(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj== TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj`,
			Stderr:  `cannot EqualTo`,
			ExitNum: 1,
		},
		{
			Block:   `"TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj"== TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj`,
			Stderr:  `cannot EqualTo`,
			ExitNum: 1,
		},
		{
			Block:   `TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj== "TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj"`,
			Stderr:  `cannot EqualTo`,
			ExitNum: 1,
		},
		{
			Block:   `"TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj"== "TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj"`,
			Stdout:  `true`,
			ExitNum: 0,
		},
		/////
		{
			Block:   `TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj == TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj`,
			Stderr:  `cannot EqualTo`,
			ExitNum: 1,
		},
		{
			Block:   `"TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj" == TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj`,
			Stderr:  `cannot EqualTo`,
			ExitNum: 1,
		},
		{
			Block:   `TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj == "TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj"`,
			Stderr:  `cannot EqualTo`,
			ExitNum: 1,
		},
		{
			Block:   `"TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj" == "TestIssue854_alksdlsdkfjsldfkjsldfkjsldfkjsldkfj"`,
			Stdout:  `true`,
			ExitNum: 0,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
