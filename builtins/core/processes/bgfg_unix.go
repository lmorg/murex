//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package processes

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
)

var rxJobId = regexp.MustCompile(`^%[0-9]+$`)

func getProcess(s string) (*lang.Process, error) {
	s = strings.TrimSpace(s)

	// empty command. Get latest
	if len(s) == 0 {
		return lang.Jobs.GetLatest()
	}

	// is a number, likely function ID
	i, err := strconv.Atoi(s)
	if err == nil {
		return lang.GlobalFIDs.Proc(uint32(i))
	}

	// has a % prefix, likely a job ID
	if rxJobId.MatchString(s) {
		i, err = strconv.Atoi(s[1:])
		if err != nil {
			return nil, fmt.Errorf("cannot get job ID from '%s': %s", s, err.Error())
		}

		return lang.Jobs.Get(i)
	}

	// lets just match it to the most recent command
	return lang.Jobs.GetFromCommandLine(s)
}

func mkbg(p *lang.Process) error {
	s := p.Parameters.StringAll()
	f, err := getProcess(s)
	if err != nil {
		return err
	}

	if f.State.Get() != state.Stopped {
		return errors.New("FID is not a stopped process. Run `jobs` or `fid-list` to see a list of stopped processes")
	}

	if f.SystemProcess.External() {
		err = f.SystemProcess.Signal(syscall.SIGCONT)
		if err != nil {
			return err
		}
	}
	go unstop(f)

	updateTree(f, true)

	f.State.Set(state.Executing)

	lang.ShowPrompt <- true
	return nil
}

func cmdForeground(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	s := p.Parameters.StringAll()
	f, err := getProcess(s)
	if err != nil {
		return err
	}

	lang.HidePrompt <- true
	go unstop(f)
	updateTree(f, false)

	lang.ForegroundProc.Set(f)
	f.State.Set(state.Executing)

	if f.SystemProcess.External() {
		lang.UnixPidToFg(f)

		err = f.SystemProcess.Signal(syscall.SIGCONT)
		if err != nil {
			// don't "return err" because we still want to wait for the process
			// to finish. So lets just print a debug message instead.
			debug.Logf("!!! failed syscall in cmdForeground()->(f: [%d] %s %s)->f.SystemProcess.Signal(syscall.SIGCONT):\n!!! error: %s",
				f.SystemProcess.Pid(), f.Name.String(), f.Parameters.StringAll(), err.Error())
		}
	}

	<-f.Context.Done()
	return nil
}

func unstop(p *lang.Process) {
	p.WaitForStopped <- true
}
