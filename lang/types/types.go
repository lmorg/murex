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
	Number    = "num"
	Integer   = "int"
	Float     = "float"
)

// TrueString is `true` boolean value
const TrueString = "true"

// FalseString is `false` boolean value
const FalseString = "false"

// TrueByte is `true` as a []byte slice
var TrueByte = []byte(TrueString)

// FalseByte is `false` as a []byte slice
var FalseByte = []byte(FalseString)

// IsTrue checks if a process has returned a `true` state.
// This will check a few conditions as not every external process will return a non-zero exit number on a failure.
func IsTrue(stdout []byte, exitNum int) bool {
	switch {
	case exitNum > 0:
		return false

	case exitNum < 0:
		return true

	default:
		s := strings.ToLower(strings.TrimSpace(string(stdout)))
		if len(s) == 0 || s == "null" || s == "0" || s == "false" || s == "no" || s == "off" || s == "fail" || s == "failed" {
			return false
		}

		return true
	}
}

// IsBlock checks if the []byte slice is a code or JSON block
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
