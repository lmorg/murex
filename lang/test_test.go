package lang

/*
	TestMurexTestingFramework tests murex's testing framework using Go's testing
	framework (confused? Essentially murex shell scripts have a testing framework
	leveredged via the `test` builtin. This can be used for testing and debugging
	murex shell script. The concept behind them is that you place all of the test
	code within the normal shell script and they sit there idle while murex is
	running. However the moment you enable the testing flag (via `config`) those
	test builtins start writing their results to the test report for your review)

	This Go source file tests that murex's test builtins and report functions
	work by testing the Go code that resides behind them.
*/

//func TestMurexTestingFramework(t *testing.T) {
//
//}
