package signals

import (
	"os"

	"github.com/lmorg/murex/lang/proc"
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

func sigtstp() {
	os.Stderr.WriteString("TODO: background")
}

func sigint(interactive bool) {
	os.Stderr.WriteString(PromptSIGINT)
	sigterm(interactive)
}

func sigterm(interactive bool) {
	if interactive {
		proc.ForegroundProc.Kill()

	} else {
		os.Exit(0)
	}
}

func sigquit(interactive bool) {
	if interactive {
		os.Stderr.WriteString(PromptSIGQUIT)

		kill := make([]func(), 0)
		p := proc.ForegroundProc
		for p.Id != 0 {
			parent := p.Parent
			if p.Kill != nil {
				kill = append(kill, p.Kill)
			}
			p = parent
		}
		for i := len(kill) - 1; i > -1; i-- {
			kill[i]()
		}

	} else {
		os.Stderr.WriteString("Murex received SIGQUIT!" + utils.NewLineString)
		os.Exit(2)
	}
}
