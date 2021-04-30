package shell

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
)

const (
	// PromptSIGINT defines the string to write when ctrl+c is pressed
	PromptSIGINT = "^C"

	// PromptSIGQUIT defines the string to write when ctrl+\ is pressed
	PromptSIGQUIT = "^\\"

	// PromptEOF defines the string to write when ctrl+d is pressed
	PromptEOF = "^D"
)

func sigint(interactive bool) {
	//os.Stderr.WriteString(PromptSIGINT)
	sigterm(interactive)
}

/*func sigtstp() {
	// see signals_unix.go
}*/

func sigterm(interactive bool) {
	if interactive {
		p := lang.ForegroundProc.Get()
		//p.Json("p =", p)

		switch {
		case p == nil:
			lang.ShellProcess.Stderr.Writeln([]byte("!!! Unable to identify forground process !!!"))
		case p.Kill == nil:
			lang.ShellProcess.Stderr.Writeln([]byte("!!! Unable to identify forground kill function !!!"))
		default:
			p.Kill()
		}

	} else {
		os.Exit(0)
	}
}

func sigquit(interactive bool) {
	if interactive {
		//os.Stderr.WriteString(PromptSIGQUIT)
		os.Stderr.WriteString("Murex received SIGQUIT!" + utils.NewLineString)

		fids := lang.GlobalFIDs.ListAll()
		for _, p := range fids {
			if p.Kill != nil /*&& !p.HasTerminated()*/ {
				procName := p.Name
				procParam, _ := p.Parameters.String(0)
				if p.Name == "exec" {
					procName = procParam
					procParam, _ = p.Parameters.String(1)
				}
				if len(procParam) > 10 {
					procParam = procParam[:10]
				}
				lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf("!!! Sending kill signal to fid %d: %s %s !!!", p.Id, procName, procParam)))
				p.Kill()
			}
		}

		lang.ShellProcess.Stderr.Writeln([]byte("!!! Starting new prompt !!!"))
		go ShowPrompt()

	} else {
		os.Stderr.WriteString("Murex received SIGQUIT!" + utils.NewLineString)
		os.Exit(2)
	}
}
