package generic_test

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestReadArray(t *testing.T) {
	input := []byte("England Scotland Ireland\na       b        c\n1       2        3\n")

	expected := []string{
		"England Scotland Ireland",
		"a       b        c",
		"1       2        3",
	}

	test.ReadArrayTest(t, types.Generic, input, expected)
}

func TestReadMap(t *testing.T) {
	input := []byte("England Scotland Ireland\na       b        c\n1       2        3\n")

	expected := []test.ReadMapExpected{
		{
			Key:   "0",
			Value: "England",
			Last:  false,
		},
		{
			Key:   "1",
			Value: "Scotland",
			Last:  false,
		},
		{
			Key:   "2",
			Value: "Ireland",
			Last:  true,
		},
		// -- next row
		{
			Key:   "0",
			Value: "a",
			Last:  false,
		},
		{
			Key:   "1",
			Value: "b",
			Last:  false,
		},
		{
			Key:   "2",
			Value: "c",
			Last:  true,
		},
		// -- next row
		{
			Key:   "0",
			Value: "1",
			Last:  false,
		},
		{
			Key:   "1",
			Value: "2",
			Last:  false,
		},
		{
			Key:   "2",
			Value: "3",
			Last:  true,
		},
	}

	test.ReadMapOrderedTest(t, types.Generic, input, expected, config.InitConf)
}

func TestArrayWriter(t *testing.T) {
	input := []string{
		"England Scotland Ireland",
		"a       b        c",
		"1       2        3",
	}

	output := "England Scotland Ireland\na       b        c\n1       2        3\n"
	test.ArrayWriterTest(t, types.Generic, input, output)
}
