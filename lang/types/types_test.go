package types_test

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

func testIsTrue(t *testing.T, val string, exp bool, exitNum int) {
	t.Helper()
	count.Tests(t, 3)
	var act bool

	act = types.IsTrue([]byte(strings.ToLower(val)), exitNum)
	if act != exp {
		t.Error("IsTrue output does not match expected:")
		t.Log("Value:   ", val)
		t.Log("Expected:", exp)
		t.Log("Actual:  ", act)
		t.Log("Exit num:", exitNum)
		t.Log("Case:      lower")
	}

	act = types.IsTrue([]byte(strings.Title(val)), exitNum)
	if act != exp {
		t.Error("IsTrue output does not match expected (tit):")
		t.Log("Value:   ", val)
		t.Log("Expected:", exp)
		t.Log("Actual:  ", act)
		t.Log("Exit num:", exitNum)
		t.Log("Case:      title")
	}

	act = types.IsTrue([]byte(strings.ToUpper(val)), exitNum)
	if act != exp {
		t.Error("IsTrue output does not match expected (tit):")
		t.Log("Value:   ", val)
		t.Log("Expected:", exp)
		t.Log("Actual:  ", act)
		t.Log("Exit num:", exitNum)
		t.Log("Case:      upper")
	}
}

func TestIsTrue(t *testing.T) {
	tests := map[string]bool{
		"":              false,
		"random string": true,
		"true":          true,
		"false":         false,
		"null":          false,
		"-1":            true,
		"-0":            true,
		"0":             false,
		"1":             true,
		"yes":           true,
		"no":            false,
		"on":            true,
		"of":            true,
		"off":           false,
		"success":       true,
		"pass":          true,
		"fail":          false,
		"failed":        false,
		"positive":      true,
		"negative":      true,
		"enabled":       true,
		"disabled":      false,
	}

	exitNums := map[int]interface{}{
		-42: true,
		-1:  true,
		0:   nil,
		1:   false,
		42:  false,
	}

	for exitNum, override := range exitNums {

		for val, exp := range tests {

			if override != nil {
				exp = override.(bool)
			}

			testIsTrue(t, val, exp, exitNum)

		}

	}

}

func TestIsBlock(t *testing.T) {
	tests := map[string]bool{
		"": false,

		"{":   false,
		" {":  false,
		"{ ":  false,
		" { ": false,

		"}":   false,
		" }":  false,
		"} ":  false,
		" } ": false,

		"}{":   false,
		" }{":  false,
		"}{ ":  false,
		" }{ ": false,

		"} {":   false,
		" } {":  false,
		"} { ":  false,
		" } { ": false,

		"{}":   true,
		" {}":  true,
		"{} ":  true,
		" {} ": true,

		"{ }":   true,
		" { }":  true,
		"{ } ":  true,
		" { } ": true,

		"{1}":   true,
		" {1}":  true,
		"{1} ":  true,
		" {1} ": true,

		"{ 1}":   true,
		" { 1}":  true,
		"{ 1} ":  true,
		" { 1} ": true,

		"{1 }":   true,
		" {1 }":  true,
		"{1 } ":  true,
		" {1 } ": true,

		"{ 1 }":   true,
		" { 1 }":  true,
		"{ 1 } ":  true,
		" { 1 } ": true,

		"{42}":   true,
		" {42}":  true,
		"{42} ":  true,
		" {42} ": true,

		"{ 42}":   true,
		" { 42}":  true,
		"{ 42} ":  true,
		" { 42} ": true,

		"{42 }":   true,
		" {42 }":  true,
		"{42 } ":  true,
		" {42 } ": true,

		"{ 42 }":   true,
		" { 42 }":  true,
		"{ 42 } ":  true,
		" { 42 } ": true,

		"{f}":   true,
		" {f}":  true,
		"{f} ":  true,
		" {f} ": true,

		"{ f}":   true,
		" { f}":  true,
		"{ f} ":  true,
		" { f} ": true,

		"{f }":   true,
		" {f }":  true,
		"{f } ":  true,
		" {f } ": true,

		"{ f }":   true,
		" { f }":  true,
		"{ f } ":  true,
		" { f } ": true,

		"{foobar}":   true,
		" {foobar}":  true,
		"{foobar} ":  true,
		" {foobar} ": true,

		"{ foobar}":   true,
		" { foobar}":  true,
		"{ foobar} ":  true,
		" { foobar} ": true,

		"{foobar }":   true,
		" {foobar }":  true,
		"{foobar } ":  true,
		" {foobar } ": true,

		"{ foobar }":   true,
		" { foobar }":  true,
		"{ foobar } ":  true,
		" { foobar } ": true,

		"{$}":   true,
		" {$}":  true,
		"{$} ":  true,
		" {$} ": true,

		"{ $}":   true,
		" { $}":  true,
		"{ $} ":  true,
		" { $} ": true,

		"{$ }":   true,
		" {$ }":  true,
		"{$ } ":  true,
		" {$ } ": true,

		"{ $ }":   true,
		" { $ }":  true,
		"{ $ } ":  true,
		" { $ } ": true,

		"{$foobar}":   true,
		" {$foobar}":  true,
		"{$foobar} ":  true,
		" {$foobar} ": true,

		"{ $foobar}":   true,
		" { $foobar}":  true,
		"{ $foobar} ":  true,
		" { $foobar} ": true,

		"{$foobar }":   true,
		" {$foobar }":  true,
		"{$foobar } ":  true,
		" {$foobar } ": true,

		"{ $foobar }":   true,
		" { $foobar }":  true,
		"{ $foobar } ":  true,
		" { $foobar } ": true,

		`{!"£$%^&*()}`:   true,
		` {!"£$%^&*()}`:  true,
		`{!"£$%^&*()} `:  true,
		` {!"£$%^&*()} `: true,

		`{!"£$%^&*() }`:   true,
		` {!"£$%^&*() }`:  true,
		`{!"£$%^&*() } `:  true,
		` {!"£$%^&*() } `: true,

		`{ !"£$%^&*()}`:   true,
		` { !"£$%^&*()}`:  true,
		`{ !"£$%^&*()} `:  true,
		` { !"£$%^&*()} `: true,

		`{ !"£$%^&*() }`:   true,
		` { !"£$%^&*() }`:  true,
		`{ !"£$%^&*() } `:  true,
		` { !"£$%^&*() } `: true,
	}

	count.Tests(t, len(tests))
	for val, exp := range tests {
		act := types.IsBlock([]byte(val))
		if act != exp {
			t.Error("IsBlock output does not match expected:")
			t.Log("Value:   ", val)
			t.Log("Expected:", exp)
			t.Log("Actual:  ", act)
		}
	}
}
