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
	Columns   = "columns"
	Binary    = "bin"
	CodeBlock = "block"
	Json      = "json"
	JsonLines = "jsonl"
	Number    = "num"
	Integer   = "int"
	Float     = "float"
	Path      = "path"
	Paths     = "paths"
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
	return IsTrueString(string(stdout), exitNum)
}

func IsTrueString(stdout string, exitNum int) bool {
	switch {
	case exitNum > 0:
		return false

	case exitNum < 0:
		return true

	default:
		s := strings.ToLower(strings.TrimSpace(stdout))
		if len(s) == 0 || s == "null" || s == "0" || s == "false" || s == "no" || s == "off" || s == "fail" || s == "failed" || s == "disabled" {
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

// IsBlockRune checks if the []rune slice is a code or JSON block
func IsBlockRune(r []rune) bool {
	r = trimSpaceRune(r)
	if len(r) < 2 {
		return false
	}

	if r[0] == '{' && r[len(r)-1] == '}' {
		return true
	}

	return false
}

func trimSpaceRune(r []rune) []rune {
	if len(r) == 0 {
		return []rune{}
	}

	for {
		l := len(r) - 1

		if r[l] == ' ' || r[l] == '\t' || r[l] == '\r' || r[l] == '\n' {
			if l == 0 {
				return []rune{}
			}
			r = r[:l]
		} else {
			break
		}
	}

	for {
		if r[0] == ' ' || r[0] == '\t' || r[0] == '\r' || r[0] == '\n' {
			if len(r) == 1 {
				return []rune{}
			}
			r = r[1:]
		} else {
			break
		}
	}

	return r
}
