package typemgmt

import (
	"fmt"
	"math"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("round", cmdRound, types.Float)
}

const (
	flagRoundDown = "--down"
	flagRoundUp   = "--up"
)

var roundArgs = &parameters.Arguments{
	Flags: map[string]string{
		flagRoundDown: types.Boolean,
		flagRoundUp:   types.Boolean,
		"-d":          flagRoundDown,
		"-u":          flagRoundUp,
	},
	AllowAdditional: true,
}

func cmdRound(p *lang.Process) error {
	p.Stdout.SetDataType(types.Number)

	flags, params, err := p.Parameters.ParseFlags(roundArgs)
	if err != nil {
		return err
	}
	if len(params) != 2 {
		return fmt.Errorf("invalid parameters. Expecting `round <value> <rounding>")
	}

	v, err := types.ConvertGoType(params[0], types.Float)
	input := v.(float64)
	if err != nil {
		return fmt.Errorf("cannot convert '%s' to float: %s", params[0], err.Error())
	}

	v, err = types.ConvertGoType(params[1], types.Float)
	if err != nil {
		return fmt.Errorf("cannot convert '%s' to float: %s", params[1], err.Error())
	}
	round := v.(float64)

	roundDown := flags[flagRoundDown] == types.TrueString
	roundUp := flags[flagRoundUp] == types.TrueString

	if roundUp && roundDown {
		return fmt.Errorf("you cannot use both %s/-d and %s/-u flags together", flagRoundDown, flagRoundUp)
	}

	switch {
	case round == 0:
		fallthrough
	case round == 1:
		switch {
		case roundDown:
			return roundWriter(p, roundDownInteger(input))
		case roundUp:
			return roundWriter(p, roundUpInteger(input))
		default:
			return roundWriter(p, roundNearestInteger(input))
		}

	case round == float64(int(round)):
		switch {
		case roundDown:
			return roundWriter(p, roundDownMultiple(int(input), int(round)))
		case roundUp:
			return roundWriter(p, roundUpMultiple(int(input), int(round)))
		default:
			return roundWriter(p, roundNearestMultiple(int(input), int(round)))
		}

	default:
		split := strings.SplitN(params[1], ".", 2)
		round := len(split[1])
		return roundWriter(p, roundNearestDecimalPlace(input, round))
	}
}

func roundWriter[Number int | float64](p *lang.Process, v Number) error {
	s, err := types.ConvertGoType(v, types.String)
	if err != nil {
		return fmt.Errorf("cannot convert %T for display: %s", v, err.Error())
	}

	_, err = p.Stdout.Write([]byte(s.(string)))
	return err
}

// round to the nearest integer
func roundNearestInteger(f float64) int {
	return int(math.Round(f))
}

// round down to the previous integer
func roundDownInteger(f float64) int {
	return int(f)
}

// round up to the next integer
func roundUpInteger(f float64) int {
	i := int(f)
	if f > float64(i) {
		return i + 1
	}
	return i
}

// round to the nearest multiple, eg:
// 2 rounded to 5 would equal 0
// 3 rounded to 5 would equal 5
func roundNearestMultiple(i, multiple int) int {
	remainder := math.Remainder(float64(i), float64(multiple))
	return i - int(remainder)
}

// round to the lowest multiple, eg:
// 2 rounded to 5 would equal 0
// 3 rounded to 5 would equal 0
// 5 rounded to 5 would equal 5
func roundDownMultiple(i, multiple int) int {
	round := i / multiple
	return round * multiple
}

// round to the highest multiple, eg:
// 2 rounded to 5 would equal 0
// 3 rounded to 5 would equal 0
// 5 rounded to 5 would equal 5
func roundUpMultiple(i, multiple int) int {
	round := i / multiple
	mod := i % multiple
	if mod == 0 {
		return i
	}
	return (round * multiple) + multiple
}

func roundNearestDecimalPlace(f float64, decPlaces int) float64 {
	ratio := math.Pow(10, float64(decPlaces))
	return math.Round(f*ratio) / ratio
}
