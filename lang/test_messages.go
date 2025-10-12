package lang

/*
	This test library relates to the testing framework within the murex
	language itself rather than Go's test framework within the murex project.

	The naming convention here is basically the inverse of Go's test naming
	convention. ie Go source files will be named "test_unit.go" (because
	calling it unit_test.go would mean it's a Go test rather than murex test)
	and the code is named UnitTestPlans (etc) rather than TestUnitPlans (etc)
	because the latter might suggest they would be used by `go test`. This
	naming convention is a little counterintuitive but it at least avoids
	naming conflicts with `go test`.
*/

import "fmt"

// func tMsgPassed() string { return "All test conditions were met" }
func tMsgPassed() string { return "-" }
func tMsgStdout(property string, stdout []byte) string {
	return fmt.Sprintf("%s output: %s", property, stdout)
}

func tMsgStderr(property string, stdout []byte) string {
	return fmt.Sprintf("%s returned an error: %s", property, stdout)
}

func tMsgReadErr(stdType string, property string, err error) string {
	return fmt.Sprintf("Error reading %s from %s: %s", stdType, property, err)
}
func tMsgWriteErr(property string, err error) string {
	return fmt.Sprintf("Error writing to stdin for %s: %s", property, err)
}

func tMsgUnmarshalErr(property string, dt string, err error) string {
	return fmt.Sprintf("Error unmarshalling `%s` for %s: %s", dt, property, err)
}
func tMsgDataFormatValid(property string, dt string, v any) string {
	return fmt.Sprintf("%s data format valid. Data-type `%s` unmarshalled as `%T`", property, dt, v)
}
func tMsgDataFormatInvalid(property string, dt string, v any) string {
	return fmt.Sprintf("%s data format invalid. Data-type `%s` unmarshalled as `%T`", property, dt, v)
}

func tMsgCompileErr(property string, err error) string {
	return fmt.Sprintf("%s failed to compile: %s", property, err)
}
func tMsgNoneZeroExit(property string, exitnum int) string {
	return fmt.Sprintf("%s exit num non-zero: %d", property, exitnum)
}

func tMsgExitNumMismatch(exp int, act int) string {
	return fmt.Sprintf("ExitNum mismatch. Exp: %d, act: %d", exp, act)
}
func tMsgExitNumMatch() string {
	return "ExitNum matches expected"
}

func tMsgExitNumNotZero(property string, exitnum int) string {
	return fmt.Sprintf("%s failed the test with an exit num of %d", property, exitnum)
}
func tMsgExitNumZero(property string) string {
	return fmt.Sprintf("%s passed the test. Returned true", property)
}

func tMsgDataTypeMismatch(stdType string, act string) string {
	return fmt.Sprintf("Data-type mismatch on %s. act: '%s'", stdType, act)
}
func tMsgDataTypeMatch(stdType string) string {
	return fmt.Sprintf("Expected data-type matched on %s", stdType)
}

func tMsgStringMismatch(property string, std []byte) string {
	return fmt.Sprintf("%s string mismatch. act: '%s'", property, std)
}
func tMsgStringMatch(property string) string {
	return fmt.Sprintf("%s matches expected string", property)
}

func tMsgRegexCompileErr(property string, err error) string {
	return fmt.Sprintf("%s could not compile: %s", property, err)
}
func tMsgRegexMismatch(property string, std []byte) string {
	return fmt.Sprintf("%s expression did not match. act: '%s'", property, std)
}
func tMsgRegexMatch(property string) string {
	return fmt.Sprintf("%s matches expected regex expression", property)
}

func tMsgGtEqFail(property string, length, comparison int) string {
	return fmt.Sprintf("%s length (%d) is less than %d", property, length, comparison)
}
func tMsgGtEqMatch(property string, length, comparison int) string {
	return fmt.Sprintf("%s length (%d) is greater than or equal to %d", property, length, comparison)
}
