package lang

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestParserErrorsBlock(t *testing.T) {
	tests := []struct {
		Block string
		Code  int
	}{
		{
			Block: `no errors here`,
			Code:  NoParsingErrors,
		},

		{
			Block: `| foobar`,
			Code:  ErrUnexpectedPipeTokenPipe,
		},
		{
			Block: `=> foobar`,
			Code:  ErrUnexpectedPipeTokenEqGt,
		},
		{
			Block: `? foobar`,
			Code:  ErrUnexpectedPipeTokenQm,
		},

		{
			Block: `{`,
			Code:  ErrUnexpectedOpenBraceFunc,
		},
		{
			Block: `}`,
			Code:  ErrUnexpectedCloseBrace,
		},
		{
			Block: `echo }`,
			Code:  ErrClosingBraceBlockNoOpen,
		},

		{
			Block: `)`,
			Code:  ErrClosingBraceQuoteNoOpen,
		},
		{
			Block: `echo )`,
			Code:  ErrClosingBraceQuoteNoOpen,
		},

		{
			Block: `echo \`,
			Code:  ErrUnterminatedEscape,
		},
		{
			Block: `echo '`,
			Code:  ErrUnterminatedQuotesSingle,
		},
		{
			Block: `echo "`,
			Code:  ErrUnterminatedQuotesDouble,
		},
		{
			Block: `echo "\"`,
			Code:  ErrUnterminatedQuotesDouble,
		},
		{
			Block: `echo {`,
			Code:  ErrUnterminatedBraceBlock,
		},
		{
			Block: `echo (`,
			Code:  ErrUnterminatedBraceQuote,
		},
		{
			Block: `echo $bob[`,
			Code:  ErrUnclosedIndex,
		},

		/*{
			Block: `echo |`,
			Code:  ErrPipingToNothing,
		},
		{
			Block: `echo ->`,
			Code:  ErrPipingToNothing,
		},
		{
			Block: `echo =>`,
			Code:  ErrPipingToNothing,
		},
		{
			Block: `echo ?`,
			Code:  ErrPipingToNothing,
		},*/

		/*{
			Block: `echo &&`,
			Code:  ErrPipingToNothing,
		},
		{
			Block: `echo ||`,
			Code:  ErrPipingToNothing,
		},*/

		/*{
			Block: `runmode $bob &&`,
			Code:  ErrUnableToParseParametersInRunmode,
		},
		{
			Block: `runmode bob`,
			Code:  ErrPipingToNothing,
		},*/
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		_, pErr := parser([]rune(test.Block))
		if pErr.Code != test.Code {
			t.Errorf("Error code doesn't match expected in test %d", i)
			t.Logf("  Block:   '%s'", test.Block)
			t.Logf("  Expected: %d (%s)", test.Code, errMessages[test.Code])
			t.Logf("  Actual:   %d (%s)", pErr.Code, errMessages[pErr.Code])
		}

	}
}
