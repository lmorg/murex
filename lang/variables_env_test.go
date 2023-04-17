package lang_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestVarEnv(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				$ENV.MurexTestVarEnv01 = "lskdjflsakdjfoiwjef;oweijflsd;kjfweo;ij"
				env -> regexp m/MurexTestVarEnv01/
			`,
			Stdout: "MurexTestVarEnv01=lskdjflsakdjfoiwjef;oweijflsd;kjfweo;ij\n",
		},
		{
			Block: `
				set ENV.MurexTestVarEnv02 = "ertyrtysdf;sldk;flkp;o342--04ik"
				env -> regexp m/MurexTestVarEnv02/
			`,
			Stdout: "MurexTestVarEnv02=ertyrtysdf;sldk;flkp;o342--04ik\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarGlobal(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				$GLOBAL.MurexTestVarGlobal01 = "fggfhsrt345u567jyujmdfgbfghbn"
				$MurexTestVarGlobal01
			`,
			Stdout: "fggfhsrt345u567jyujmdfgbfghbn",
		},
		{
			Block: `
				set GLOBAL.MurexTestVarGlobal02 = "sdfp23io4j3409asLKJHD2E9OP8I2340kjhlkj"
				$MurexTestVarGlobal02
			`,
			Stdout: "sdfp23io4j3409asLKJHD2E9OP8I2340kjhlkj",
		},
	}

	test.RunMurexTests(tests, t)
}
