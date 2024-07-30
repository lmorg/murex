package expressions_test

import (
	"fmt"
	"os/user"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/utils/home"
)

func TestParseVarsScalar(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseVarsScalar0="0";$TestParseVarsScalar0`,
			Stdout: `0`,
		},
		{
			Block:  `TestParseVarsScalar1=1;$TestParseVarsScalar1`,
			Stdout: `1`,
		},
		{
			Block:  `TestParseVarsScalar2="foobar";TestParseVarsScalar2="> $TestParseVarsScalar2 <";$TestParseVarsScalar2`,
			Stdout: `> foobar <`,
		},
		{
			Block:  `TestParseVarsScalar3=3;TestParseVarsScalar3="> $TestParseVarsScalar3 <";$TestParseVarsScalar3`,
			Stdout: `> 3 <`,
		},
		/////
		{
			Block:  `TestParseVarsScalar4=4;%[1 2 3 $TestParseVarsScalar4]`,
			Stdout: `[1,2,3,4]`,
		},
		{
			Block:  `TestParseVarsScalar5=5;%{$TestParseVarsScalar5: "foobar"}`,
			Stdout: `{"5":"foobar"}`,
		},
		{
			Block:  `TestParseVarsScalar6=6;%{"foobar": $TestParseVarsScalar6}`,
			Stdout: `{"foobar":6}`,
		},
		/////
		{
			Block:  `TestParseVarsScalar7=1;TestParseVarsScalar7=2;TestParseVarsScalar7=3;$TestParseVarsScalar7`,
			Stdout: `3`,
		},
		{
			Block:  `TestParseVarsScalar8=1;TestParseVarsScalar8=2;TestParseVarsScalar8=3;;$TestParseVarsScalar8`,
			Stdout: `3`,
		},
		{
			Block:  `TestParseVarsScalar9=1;TestParseVarsScalar9=2;TestParseVarsScalar9=3;out bob;$TestParseVarsScalar9`,
			Stdout: "bob\n3",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseVarsArray(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseVarsArray0=%[1..3];$TestParseVarsArray0`,
			Stdout: `[1,2,3]`,
		},
		{
			Block:  `TestParseVarsArray1=%[1..3];TestParseVarsArray1="> $TestParseVarsArray1 <";$TestParseVarsArray1`,
			Stdout: `> [1,2,3] <`,
		},
		/////
		{
			Block:  `TestParseVarsArray2=%[1..3];@TestParseVarsArray2`,
			Stdout: `[1,2,3]`,
		},
		{
			Block:  `TestParseVarsArray3=%[1..3];TestParseVarsArray3="> @TestParseVarsArray3 <";$TestParseVarsArray3`,
			Stdout: `> @TestParseVarsArray3 <`,
		},
		{
			Block:  `TestParseVarsArray4=%[1..3];TestParseVarsArray4=". @TestParseVarsArray4 .";@TestParseVarsArray4`,
			Stdout: `[". @TestParseVarsArray4 ."]`,
		},
		/////
		{
			Block:  `TestParseVarsArray5=%[1..3];%[1 2 3 @TestParseVarsArray5]`,
			Stdout: `[1,2,3,1,2,3]`,
		},
		{
			Block:  `TestParseVarsArray6=%[1..3];%[1 2 3 [@TestParseVarsArray6]]`,
			Stdout: `[1,2,3,[1,2,3]]`,
		},
		/////
		{
			Block:  `TestParseVarsArray7=%[1..3];%{a: @TestParseVarsArray7}`,
			Stdout: `{"a":[1,2,3]}`,
		},
		{
			Block:  `TestParseVarsArray8=%[1..3];%{a: [@TestParseVarsArray8]}`,
			Stdout: `{"a":[1,2,3]}`,
		},
		{
			Block:  `TestParseVarsArray8=%[1..3];%{a: [[@TestParseVarsArray8]]}`,
			Stdout: `{"a":[[1,2,3]]}`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseVarsTilder(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseVarsTilder0= ~;$TestParseVarsTilder0`,
			Stdout: home.MyDir,
		},
		{
			Block:  `TestParseVarsTilder1="~";$TestParseVarsTilder1`,
			Stdout: home.MyDir,
		},
		{
			Block:  `TestParseVarsTilder2='~';$TestParseVarsTilder2`,
			Stdout: `~`,
		},
		{
			Block:  `%[~]`,
			Stdout: fmt.Sprintf(`["%s"]`, home.MyDir),
		},
		{
			Block:  `%{~: a}`,
			Stdout: fmt.Sprintf(`{"%s":"a"}`, home.MyDir),
		},
		{
			Block:  `%{a: ~}`,
			Stdout: fmt.Sprintf(`{"a":"%s"}`, home.MyDir),
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseVarsTilderPlusNamePositive(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Errorf("cannot run tests: %s", err.Error())
		return
	}

	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`TestParseVarsTilderPlusName0= ~%s;$TestParseVarsTilderPlusName0`, usr.Username),
			Stdout: home.MyDir,
		},
		{
			Block:  fmt.Sprintf(`TestParseVarsTilderPlusName1="~%s";$TestParseVarsTilderPlusName1`, usr.Username),
			Stdout: home.MyDir,
		},
		{
			Block:  fmt.Sprintf(`TestParseVarsTilderPlusName2='~%s';$TestParseVarsTilderPlusName2`, usr.Username),
			Stdout: fmt.Sprintf(`~%s`, usr.Username),
		},
		{
			Block:  `%[~` + usr.Username + `]`,
			Stdout: fmt.Sprintf(`["%s"]`, home.MyDir),
		},
		{
			Block:  `%{~` + usr.Username + `: a}`,
			Stdout: fmt.Sprintf(`{"%s":"a"}`, home.MyDir),
		},
		{
			Block:  `%{a: ~` + usr.Username + `}`,
			Stdout: fmt.Sprintf(`{"a":"%s"}`, home.MyDir),
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseVarsTilderPlusNameNegative(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `TestParseVarsTilderPlusName0= ~TestParseVarsTilderPlusName0;$TestParseVarsTilderPlusName0`,
			Stderr:  "cannot expand variable",
			ExitNum: 1,
		},
		{
			Block:   `TestParseVarsTilderPlusName1="~TestParseVarsTilderPlusName1";$TestParseVarsTilderPlusName1`,
			Stderr:  "cannot expand variable",
			ExitNum: 1,
		},
		{
			Block:   `TestParseVarsTilderPlusName2='~TestParseVarsTilderPlusName2';$TestParseVarsTilderPlusName2`,
			Stdout:  `~TestParseVarsTilderPlusName2`,
		},
		{
			Block:   `%[~TestParseVarsTilderPlusName3]`,
			Stderr:  "cannot expand variable",
			ExitNum: 1,
		},
		{
			Block:   `%{~TestParseVarsTilderPlusName4: a}`,
			Stderr:  "cannot expand variable",
			ExitNum: 1,
		},
		{
			Block:   `%{a: ~TestParseVarsTilderPlusName5}`,
			Stderr:  "cannot expand variable",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestParseVarsIndex(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseVarsIndex0=%[1..3];$TestParseVarsIndex0[1]`,
			Stdout: `2`,
		},
		{
			Block:  `TestParseVarsIndex1=%[1..3];TestParseVarsIndex1=$TestParseVarsIndex1[2];$TestParseVarsIndex1`,
			Stdout: `3`,
		},
		{
			Block:  `TestParseVarsIndex2=%[1..3];TestParseVarsIndex2="-$TestParseVarsIndex2[1]-";$TestParseVarsIndex2`,
			Stdout: `-2-`,
		},
		{
			Block:  `TestParseVarsIndex3=%[1..3];%[1 2 3 $TestParseVarsIndex3[1] 1 2 3]`,
			Stdout: `[1,2,3,2,1,2,3]`,
		},
		{
			Block:  `TestParseVarsIndex4=%[1..3];%{$TestParseVarsIndex4[1]:a}`,
			Stdout: `{"2":"a"}`,
		},
		{
			Block:  `TestParseVarsIndex5=%[1..3];%{a:$TestParseVarsIndex5[1]}`,
			Stdout: `{"a":2}`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseVarsElementSlash(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseVarsElement0=%[1..3];$TestParseVarsElement0[[/1]]`,
			Stdout: `2`,
		},
		{
			Block:  `TestParseVarsElement1=%[1..3];TestParseVarsElement1=$TestParseVarsElement1[[/2]];$TestParseVarsElement1`,
			Stdout: `3`,
		},
		{
			Block:  `TestParseVarsElement2=%[1..3];TestParseVarsElement2="-$TestParseVarsElement2[[/1]]-";$TestParseVarsElement2`,
			Stdout: `-2-`,
		},
		{
			Block:  `TestParseVarsElement3=%[1..3];%[1 2 3 $TestParseVarsElement3[[/1]] 1 2 3]`,
			Stdout: `[1,2,3,2,1,2,3]`,
		},
		{
			Block:  `TestParseVarsElement4=%[1..3];%{$TestParseVarsElement4[[/1]]:a}`,
			Stdout: `{"2":"a"}`,
		},
		{
			Block:  `TestParseVarsElement5=%[1..3];%{a:$TestParseVarsElement5[[/1]]}`,
			Stdout: `{"a":2}`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseVarsElementDot(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseVarsElement0=%[1..3];$TestParseVarsElement0[[.1]]`,
			Stdout: `2`,
		},
		{
			Block:  `TestParseVarsElement1=%[1..3];TestParseVarsElement1=$TestParseVarsElement1[[.2]];$TestParseVarsElement1`,
			Stdout: `3`,
		},
		{
			Block:  `TestParseVarsElement2=%[1..3];TestParseVarsElement2="-$TestParseVarsElement2[[.1]]-";$TestParseVarsElement2`,
			Stdout: `-2-`,
		},
		{
			Block:  `TestParseVarsElement3=%[1..3];%[1 2 3 $TestParseVarsElement3[[.1]] 1 2 3]`,
			Stdout: `[1,2,3,2,1,2,3]`,
		},
		{
			Block:  `TestParseVarsElement4=%[1..3];%{$TestParseVarsElement4[[.1]]:a}`,
			Stdout: `{"2":"a"}`,
		},
		{
			Block:  `TestParseVarsElement5=%[1..3];%{a:$TestParseVarsElement5[[.1]]}`,
			Stdout: `{"a":2}`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseVarsDotNotation(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseVarsDotNotation0=%[1..3];$TestParseVarsDotNotation0.1`,
			Stdout: `2`,
		},
		{
			Block:  `TestParseVarsDotNotation1=%[1..3];TestParseVarsDotNotation1=$TestParseVarsDotNotation1.2;$TestParseVarsDotNotation1`,
			Stdout: `3`,
		},
		{
			Block:  `TestParseVarsDotNotation2=%[1..3];TestParseVarsDotNotation2="-$TestParseVarsDotNotation2.1-";$TestParseVarsDotNotation2`,
			Stdout: `-2-`,
		},
		{
			Block:  `TestParseVarsDotNotation3=%[1..3];%[1 2 3 $TestParseVarsDotNotation3.1 1 2 3]`,
			Stdout: `[1,2,3,2,1,2,3]`,
		},
		{
			Block:  `TestParseVarsDotNotation4=%[1..3];%{$TestParseVarsDotNotation4.1:a}`,
			Stdout: `{"2":"a"}`,
		},
		{
			Block:  `TestParseVarsDotNotation5=%[1..3];%{a:$TestParseVarsDotNotation5.1}`,
			Stdout: `{"a":2}`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseVarsParen(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseVarsParen0="foobar";$(TestParseVarsParen0)`,
			Stdout: `^foobar$`,
		},
		{
			Block:  `TestParseVarsParen1=%[1 2 3];$(TestParseVarsParen1)`,
			Stdout: `[1,2,3]`,
		},
		{
			Block:  `TestParseVarsParen2=%[1 2 3];$(TestParseVarsParen2.1)`,
			Stdout: `2`,
		},
		/*{
			Block:  `TestParseVarsParen3=%[1 2 3];$(TestParseVarsParen3[2])`,
			Stdout: `3`,
		},*/
		/*{
			Block:  `TestParseVarsParen4=%[1 2 3];$(TestParseVarsParen4[[.1]])`,
			Stdout: `2`,
		},*/
		//
		{
			Block:  `TestParseVarsParen5=%[1..3];echo -$(TestParseVarsParen5.1)-`,
			Stdout: `^-2-\n$`,
		},
		{
			Block:  `TestParseVarsParen6=%[1..3];TestParseVarsParen6="-$(TestParseVarsParen6.1)-";$(TestParseVarsParen6)`,
			Stdout: `^-2-$`,
		},
		{
			Block:  `TestParseVarsParen7=%[1..3];%[1 2 3 $(TestParseVarsParen7.1) 1 2 3]`,
			Stdout: `[1,2,3,2,1,2,3]`,
		},
		{
			Block:  `TestParseVarsParen8=%[1..3];%{$(TestParseVarsParen8.1):a}`,
			Stdout: `{"2":"a"}`,
		},
		{
			Block:  `TestParseVarsParen9=%[1..3];%{a:$(TestParseVarsParen9.1)}`,
			Stdout: `{"a":2}`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
