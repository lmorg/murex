package types

import (
	"bytes"
	"strings"
)

const (
	// These are only a list of the base types. Others can be added via builtins or during runtime. However their
	// behavior will default to string (str).
	Generic   = "*"
	Null      = "null"
	Die       = "die"
	Boolean   = "bool"
	String    = "str"
	Binary    = "bin"
	CodeBlock = "block"
	Json      = "json"
	Csv       = "csv"
	Number    = "num"
	Integer   = "int"
	Float     = "float"
)

const TrueString = "true"
const FalseString = "false"

var TrueByte = []byte(TrueString)
var FalseByte = []byte(FalseString)

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

func IsBlock(b []byte) bool {
	b = bytes.TrimSpace(b)
	if len(b) < 2 {
		return false
	}

	if b[0] == '{' && b[len(b)-1] == '}' {
		return true
	}

	return false
}
