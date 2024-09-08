package sigfns

import (
	"fmt"
	"regexp"

	"github.com/lmorg/murex/lang"
)

func Sigint(interactive bool) {
	Sigterm(interactive)
}

func Sigterm(interactive bool) {
	if !interactive {
		lang.Exit(0)
	}

	p := lang.ForegroundProc.Get()

	switch {
	case p == nil:
		//lang.ShellProcess.Stderr.Writeln([]byte("!!! Unable to identify foreground process"))
	case p.Kill == nil:
		//lang.ShellProcess.Stderr.Writeln([]byte("!!! Unable to identify foreground kill function"))
	default:
		p.Kill()
	}
}

var rxWhiteSpace = regexp.MustCompilePOSIX(`[\r\n\t ]+`)

func Sigquit(interactive bool) {
	if !interactive {
		lang.Exit(2)
	}

	fids := lang.GlobalFIDs.ListAll()
	for _, p := range fids {
		if p.Kill != nil /*&& !p.HasTerminated()*/ {
			procName := p.Name.String()
			procParam, _ := p.Parameters.String(0)
			if procName == "exec" {
				procName = procParam
				procParam, _ = p.Parameters.String(1)
			}
			if len(procParam) > 60 {
				procParam = procParam[:60] + "..."
			}
			procParam = rxWhiteSpace.ReplaceAllString(procParam, " ")

			lang.ShellProcess.Stderr.Writeln([]byte(
				fmt.Sprintf(
					"!!! Force closing FID %d: %s %s",
					p.Id, procName, procParam)))
			p.Kill()

			if p.SystemProcess.External() {
				err := p.SystemProcess.Kill()
				if err != nil {
					lang.ShellProcess.Stderr.Writeln([]byte(
						fmt.Sprintf(
							"!!! Error terminating FID %d (%d), `%s`: %s",
							p.Id, p.SystemProcess.Pid(), procName, err.Error())))
				}
			}
		}
	}

	lang.ShowPrompt <- true
}
