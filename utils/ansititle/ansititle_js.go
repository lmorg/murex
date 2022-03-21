//go:build js
// +build js

package ansititle

import (
	"errors"
	"os"

	"github.com/lmorg/murex/utils/readline"
)

func Write(title []byte) error { return nil }
func Icon(title []byte) error  { return nil }
func Tmux(title []byte) error  { return nil }
