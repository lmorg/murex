package parser_test

import (
	"encoding/json"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/parser"
)

func quickJson(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "    ")
	return string(b)
}

func testParser(t *testing.T, block string, pos int,
	Escaped bool,
	Comment bool,
	QuoteSingle bool,
	QuoteDouble bool,
	QuoteBrace int,
	NestedBlock int,
	SquareBracket bool,
	ExpectFunc bool,
	LastFuncName string,
	FuncName string,
	Parameters []string,
	Variable string,
	PipeToken parser.PipeToken) {

	count.Tests(t, 1)
	t.Helper()

	pt, _ := parser.Parse([]rune(block), pos)
	var e bool

	if Escaped != pt.Escaped {
		t.Errorf("Escaped mismatch: Escaped (%v) != pt.Escaped (%v)", Escaped, pt.Escaped)
		e = true
	}
	if Comment != pt.Comment {
		t.Errorf("Comment mismatch: Comment (%v) != pt.Comment (%v)", Comment, pt.Comment)
		e = true
	}
	if QuoteSingle != pt.QuoteSingle {
		t.Errorf("QuoteSingle mismatch: QuoteSingle (%v) != pt.QuoteSingle (%v)", QuoteSingle, pt.QuoteSingle)
		e = true
	}
	if QuoteDouble != pt.QuoteDouble {
		t.Errorf("QuoteDouble mismatch: QuoteDouble (%v) != pt.QuoteDouble (%v)", QuoteDouble, pt.QuoteDouble)
		e = true
	}
	if QuoteBrace != pt.QuoteBrace {
		t.Errorf("QuoteBrace mismatch: QuoteBrace (%d) != pt.QuoteBrace (%d)", QuoteBrace, pt.QuoteBrace)
		e = true
	}
	if NestedBlock != pt.NestedBlock {
		t.Errorf("NestedBlock mismatch: NestedBlock (%d) != pt.NestedBlock (%d)", NestedBlock, pt.NestedBlock)
		e = true
	}
	if SquareBracket != pt.SquareBracket {
		t.Errorf("SquareBracket mismatch: SquareBracket (%v) != pt.SquareBracket (%v)", SquareBracket, pt.SquareBracket)
		e = true
	}
	if ExpectFunc != pt.ExpectFunc {
		t.Errorf("ExpectFunc mismatch: ExpectFunc (%v) != pt.ExpectFunc (%v)", ExpectFunc, pt.ExpectFunc)
		e = true
	}
	if LastFuncName != pt.LastFuncName {
		t.Errorf("LastFuncName mismatch: LastFuncName (%s) != pt.LastFuncName (%s)", LastFuncName, pt.LastFuncName)
		e = true
	}
	if FuncName != pt.FuncName {
		t.Errorf("FuncName mismatch: FuncName (%s) != pt.FuncName (%s)", FuncName, pt.FuncName)
		e = true
	}
	if Variable != pt.Variable {
		t.Errorf("Variable mismatch: Variable (%s) != pt.Variable (%s)", Variable, pt.Variable)
		e = true
	}
	if PipeToken != pt.PipeToken {
		t.Errorf("PipeToken mismatch: PipeToken (%s) != pt.PipeToken (%s)", PipeToken, pt.PipeToken)
		e = true
	}
	if len(Parameters) != len(pt.Parameters) {
		t.Errorf("Parameters mismatch: Parameters (%s) != pt.Parameters (%s)", quickJson(Parameters), quickJson(pt.Parameters))
		e = true
	}
	for i := range Parameters {
		if Parameters[i] != pt.Parameters[i] {
			t.Errorf("Parameters mismatch: Parameters (%s) != pt.Parameters (%s)", quickJson(Parameters), quickJson(pt.Parameters))
			e = true
			break
		}
	}

	if e {
		t.Logf("  Block: %s", block)
	}
}

func TestParser(t *testing.T) {
	testParser(t, "out ", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{}, "", parser.PipeTokenNone)

	testParser(t, " out Hello ", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello"}, "", parser.PipeTokenNone)
	testParser(t, "  out Hello ", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello"}, "", parser.PipeTokenNone)
	testParser(t, "out Hello ", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello"}, "", parser.PipeTokenNone)
	testParser(t, "out  Hello ", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello"}, "", parser.PipeTokenNone)
	testParser(t, "out Hello  ", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello"}, "", parser.PipeTokenNone)
	testParser(t, "out:Hello ", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello"}, "", parser.PipeTokenNone)
	testParser(t, "out: Hello ", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello"}, "", parser.PipeTokenNone)
}

func TestParserEscaped(t *testing.T) {
	// escaped chars
	testParser(t, `out \`, 0,
		true, false, false, false, 0, 0, false, false, "", "out", []string{``}, "", parser.PipeTokenNone)
	testParser(t, `out \\`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`\`}, "", parser.PipeTokenNone)
	testParser(t, `out \#`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`#`}, "", parser.PipeTokenNone)
	testParser(t, `out \'`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`'`}, "", parser.PipeTokenNone)
	testParser(t, `out \"`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`"`}, "", parser.PipeTokenNone)
	testParser(t, `out \(`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`(`}, "", parser.PipeTokenNone)
	testParser(t, `out \)`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`)`}, "", parser.PipeTokenNone)
	testParser(t, `out \-`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`-`}, "", parser.PipeTokenNone)
	testParser(t, `out \ `, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{` `}, "", parser.PipeTokenNone)
	testParser(t, `out \:`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`:`}, "", parser.PipeTokenNone)
	testParser(t, `out \>`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`>`}, "", parser.PipeTokenNone)
	testParser(t, `out \|`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`|`}, "", parser.PipeTokenNone)
	testParser(t, `out \;`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`;`}, "", parser.PipeTokenNone)
	testParser(t, `out \r`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"\r"}, "", parser.PipeTokenNone)
	testParser(t, `out \n`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"\n"}, "", parser.PipeTokenNone)
	testParser(t, `out \s`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{" "}, "", parser.PipeTokenNone)
	testParser(t, `out \t`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"\t"}, "", parser.PipeTokenNone)
	testParser(t, `out \?`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`?`}, "", parser.PipeTokenNone)
	testParser(t, `out \{`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`{`}, "", parser.PipeTokenNone)
	testParser(t, `out \}`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`}`}, "", parser.PipeTokenNone)
	testParser(t, `out \[`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`[`}, "", parser.PipeTokenNone)
	testParser(t, `out \]`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`]`}, "", parser.PipeTokenNone)
	testParser(t, `out \$`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`$`}, "", parser.PipeTokenNone)
	testParser(t, `out \@`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`@`}, "", parser.PipeTokenNone)
	testParser(t, `out \<`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`<`}, "", parser.PipeTokenNone)

	// escape inside quotes
	testParser(t, `out '\\'`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`\\`}, "", parser.PipeTokenNone)
	testParser(t, `out "\\"`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`\`}, "", parser.PipeTokenNone)
	testParser(t, `out (\\)`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`\\`}, "", parser.PipeTokenNone)

	// escaping a random char
	testParser(t, `out \q`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`q`}, "", parser.PipeTokenNone)
}

func TestParserComment(t *testing.T) {
	testParser(t, "#out Hello world", 0,
		false, true, false, false, 0, 0, false, true, "", "", []string{}, "", parser.PipeTokenNone)
	testParser(t, "out foo #out bar", 0,
		false, true, false, false, 0, 0, false, false, "", "out", []string{"foo"}, "", parser.PipeTokenNone)
	testParser(t, `out '#'`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`#`}, "", parser.PipeTokenNone)
	testParser(t, `out "#"`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`#`}, "", parser.PipeTokenNone)
	testParser(t, `out (#)`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`#`}, "", parser.PipeTokenNone)
}

func TestParserQuoteSingle(t *testing.T) {
	// As a parameter
	testParser(t, "out 'Hello world", 0,
		false, false, true, false, 0, 0, false, false, "", "out", []string{"Hello world"}, "", parser.PipeTokenNone)
	testParser(t, "out 'Hello world'", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello world"}, "", parser.PipeTokenNone)
	testParser(t, `out 'Hello world"`, 0,
		false, false, true, false, 0, 0, false, false, "", "out", []string{`Hello world"`}, "", parser.PipeTokenNone)

	// As a function
	testParser(t, `'out' foo bar`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"foo", "bar"}, "", parser.PipeTokenNone)
	testParser(t, `'out foo bar`, 0,
		false, false, true, false, 0, 0, false, true, "", "out foo bar", []string{}, "", parser.PipeTokenNone)
	testParser(t, `'out foo bar'`, 0,
		false, false, false, false, 0, 0, false, true, "", "out foo bar", []string{}, "", parser.PipeTokenNone)

	// Other quotes
	testParser(t, `out '"'`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`"`}, "", parser.PipeTokenNone)
	testParser(t, `out '('`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`(`}, "", parser.PipeTokenNone)
	testParser(t, `out ')'`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`)`}, "", parser.PipeTokenNone)
}

func TestParserQuoteDouble(t *testing.T) {
	// As a parameter
	testParser(t, `out "Hello world`, 0,
		false, false, false, true, 0, 0, false, false, "", "out", []string{"Hello world"}, "", parser.PipeTokenNone)
	testParser(t, `out "Hello world"`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello world"}, "", parser.PipeTokenNone)
	testParser(t, `out "Hello world'`, 0,
		false, false, false, true, 0, 0, false, false, "", "out", []string{"Hello world'"}, "", parser.PipeTokenNone)

	// As a function
	testParser(t, `"out" foo bar`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"foo", "bar"}, "", parser.PipeTokenNone)
	testParser(t, `"out foo bar`, 0,
		false, false, false, true, 0, 0, false, true, "", "out foo bar", []string{}, "", parser.PipeTokenNone)
	testParser(t, `"out foo bar"`, 0,
		false, false, false, false, 0, 0, false, true, "", "out foo bar", []string{}, "", parser.PipeTokenNone)

	// Other quotes
	testParser(t, `out "'"`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`'`}, "", parser.PipeTokenNone)
	testParser(t, `out "("`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`(`}, "", parser.PipeTokenNone)
	testParser(t, `out ")"`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`)`}, "", parser.PipeTokenNone)
}

func TestParserQuoteBrace(t *testing.T) {
	// As a parameter
	testParser(t, "out (Hello world", 0,
		false, false, false, false, 1, 0, false, false, "", "out", []string{"Hello world"}, "", parser.PipeTokenNone)
	testParser(t, "out (Hello world)", 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{"Hello world"}, "", parser.PipeTokenNone)
	testParser(t, "out ((Hello world", 0,
		false, false, false, false, 2, 0, false, false, "", "out", []string{"(Hello world"}, "", parser.PipeTokenNone)
	testParser(t, "out ((Hello (world", 0,
		false, false, false, false, 3, 0, false, false, "", "out", []string{"(Hello (world"}, "", parser.PipeTokenNone)
	testParser(t, "out ((Hello world)", 0,
		false, false, false, false, 1, 0, false, false, "", "out", []string{"(Hello world)"}, "", parser.PipeTokenNone)
	testParser(t, "out ((Hello (world)", 0,
		false, false, false, false, 2, 0, false, false, "", "out", []string{"(Hello (world)"}, "", parser.PipeTokenNone)

	// As a function
	testParser(t, `(out) foo bar`, 0,
		false, false, false, false, 0, 0, false, false, "", "(", []string{"out", "foo", "bar"}, "", parser.PipeTokenNone)
	testParser(t, `(out foo bar`, 0,
		false, false, false, false, 1, 0, false, false, "", "(", []string{"out foo bar"}, "", parser.PipeTokenNone)
	testParser(t, `(out foo bar)`, 0,
		false, false, false, false, 0, 0, false, false, "", "(", []string{"out foo bar"}, "", parser.PipeTokenNone)

	// Other quotes
	testParser(t, `out (')`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`'`}, "", parser.PipeTokenNone)
	testParser(t, `out (")`, 0,
		false, false, false, false, 0, 0, false, false, "", "out", []string{`"`}, "", parser.PipeTokenNone)
}

func TestParserPipeBasic(t *testing.T) {
	testParser(t, "foo | bar", 0,
		false, false, false, false, 0, 0, false, true, "foo", "bar", []string{}, "", parser.PipeTokenPosix)
	testParser(t, "foo -> bar", 0,
		false, false, false, false, 0, 0, false, true, "foo", "bar", []string{}, "", parser.PipeTokenArrow)
	testParser(t, "foo => bar", 0,
		false, false, false, false, 0, 0, false, true, "foo", "bar", []string{}, "", parser.PipeTokenGeneric)
	testParser(t, "foo ? bar", 0,
		false, false, false, false, 0, 0, false, true, "foo", "bar", []string{}, "", parser.PipeTokenRedirect)
}
