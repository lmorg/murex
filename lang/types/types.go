package types

import "strings"

const (
	// system types
	Generic = "*"
	Null    = "null"
	Die     = "die"
	Boolean = "bool"

	// string types
	String    = "str"
	Binary    = "bin"
	Json      = "json"
	CodeBlock = "block"

	// number types
	Number  = "num"
	Integer = "int"
	Float   = "float"
)

func IsTrue(stdout []byte, exitNum int) bool {
	if exitNum != 0 {
		return false
	}

	s := strings.ToLower(strings.TrimSpace(string(stdout)))
	if len(s) == 0 || s == "null" || s == "0" || s == "false" || s == "no" || s == "off" || s == "fail" || s == "failed" {
		return false
	}

	return true
}

const TrueString = "true"
const FalseString = "false"

var TrueByte = []byte(TrueString)
var FalseByte = []byte(FalseString)
