package types

import (
	"bytes"
	"strings"
)

// These are only a list of the base types. Others can be added via builtins or during runtime. However their
// behavior will default to string (str).
const (
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

// `true` boolean value
const TrueString = "true"

// `false` boolean value
const FalseString = "false"

// `true` as a []byte slice
var TrueByte = []byte(TrueString)

// `false` as a []byte slice
var FalseByte = []byte(FalseString)

// Checks if a process has returned a `true` state.
// This will check a few conditions as not every external process will return a non-zero exit number on a failure.
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

// Checks if the []byte slice is a code or JSON block
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
