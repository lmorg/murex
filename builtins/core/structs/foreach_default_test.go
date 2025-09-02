package structs_test

import (
    "testing"

    "github.com/lmorg/murex/test"
)

func TestForEachDefaultVars(t *testing.T) {
    tests := []test.MurexTest{
        {
            Block:  `a [1..6] -> foreach i { out $i }`,
            Stdout: `^1\n2\n3\n4\n5\n6\n$`,
            Stderr: `^$`,
        },
        {
            Block:  `a [1..6] -> foreach i { echo $i }`,
            Stdout: `^1\n2\n3\n4\n5\n6\n$`,
            Stderr: `^$`,
        },
    }

    test.RunMurexTestsRx(tests, t)
}

