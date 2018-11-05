package shell

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
)

const interruptPrompt = "^C"
const eofPrompt = "^D"

// SigHandler is an internal function to capture and handle OS signals (eg SIGTERM).
func SigHandler() {
	defer func() {
		if r := recover(); r != nil {
			os.Stderr.WriteString(fmt.Sprintln("Exception caught: ", r))
			SigHandler()
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for {
			sig := <-c
			switch sig.String() {
			case syscall.SIGTERM.String():
				os.Stderr.WriteString("Shell received SIGTERM!" + utils.NewLineString)
				os.Exit(1)

			case os.Interrupt.String():
				if Prompt == nil {
					go proc.ForegroundProc.Kill()
					//os.Stderr.WriteString(interruptPrompt)

				} else {
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
					//os.Stderr.WriteString(interruptPrompt)
				}

			case syscall.SIGQUIT.String():
				os.Stderr.WriteString("Shell received SIGQUIT!" + utils.NewLineString)
				os.Exit(2)

			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}
