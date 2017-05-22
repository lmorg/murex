// +build windows

package proc

import (
	"github.com/kr/pty"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Disable PTY support in Windows.
func ExternalPty(p *Process) error {
	return External(p)
}
