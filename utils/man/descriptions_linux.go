//go:build linux
// +build linux

package man

import (
	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/lang/stdio"
)

var manBlock = []rune(`man $command`)
