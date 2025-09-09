//go:build !windows
// +build !windows

package lang_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestVarGlobal(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				$GLOBAL.MurexTestVarGlobal00 = "fggfhsrt345u567jyujmdfgbfghbn"
				$MurexTestVarGlobal00
			`,
			Stdout: "fggfhsrt345u567jyujmdfgbfghbn",
		},
		{
			Block: `
				set GLOBAL.MurexTestVarGlobal01 = "sdfp23io4j3409asLKJHD2E9OP8I2340kjhlkj"
				$MurexTestVarGlobal01
			`,
			Stdout: "sdfp23io4j3409asLKJHD2E9OP8I2340kjhlkj",
		},
		{
			Block: `
				$GLOBAL.MurexTestVarGlobal02 = "abc"
				$MurexTestVarGlobal02 -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "str",
		},
		{
			Block: `
				$GLOBAL.MurexTestVarGlobal03 = 123
				$MurexTestVarGlobal03 -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "num",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarEnv(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				$ENV.MurexTestVarEnv00 = "lskdjflsakdjfoiwjef;oweijflsd;kjfweo;ij"
				/usr/bin/env -> regexp m/MurexTestVarEnv00/
			`,
			Stdout: "MurexTestVarEnv00=lskdjflsakdjfoiwjef;oweijflsd;kjfweo;ij\n",
		},
		{
			Block: `
				set ENV.MurexTestVarEnv01 = "ertyrtysdf;sldk;flkp;o342--04ik"
				/usr/bin/env -> regexp m/MurexTestVarEnv01/
			`,
			Stdout: "MurexTestVarEnv01=ertyrtysdf;sldk;flkp;o342--04ik\n",
		},
		{
			Block: `
				$ENV.MurexTestVarEnv02 = "abc"
				$MurexTestVarEnv02 -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "str",
		},
		{
			Block: `
				$ENV.MurexTestVarEnv03 = 123
				$MurexTestVarEnv03 -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "str",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarDotType(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				$TestVarDotType00 = %{a:1, b:2, c:3}
				$TestVarDotType00 -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "^json$",
		},
		{
			Block: `
				$TestVarDotType01 = %{a:1, b:2, c:3}
				$TestVarDotType01.b -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "^num$",
		},
		{
			Block: `
				$TestVarDotType02 = %{1:a, 2:b, 3:c}
				$TestVarDotType02.2 -> debug -> [[ /data-type/murex ]]
				$TestVarDotType02.2 = 10
				$TestVarDotType02.2 -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "^strstr$",
		},
		{
			Block: `
				$TestVarDotType03 = %{a:1, b:2, c:3}
				$TestVarDotType03.b -> debug -> [[ /data-type/murex ]]
				$TestVarDotType03.b = "abc"
				$TestVarDotType03.b -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "^numnum$",
			Stderr: "cannot convert 'abc' to a floating point number",
		},
		{
			Block: `
				$TestVarDotType04 = %{1:a, 2:b, 3:c, 4: [1, 2, 3]}
				$TestVarDotType04.4 -> debug -> [[ /data-type/murex ]]
				$TestVarDotType04.5 = 10
				$TestVarDotType04.5 -> debug -> [[ /data-type/murex ]]
			`,
			Stdout: "^jsonnum$",
		},
	}

	test.RunMurexTestsRx(tests, t)
}
