//go:build js
// +build js

package home

// MyDir is the $USER directory.
var MyDir = "/"

// UserDir is the home directory of a `username`.
func UserDir(_ string) string {
	return "/"
}
