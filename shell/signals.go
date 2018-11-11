package shell

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
)

const promptInterrupt = "^C"
const promptSIGQUIT = "^\\"
const promptEOF = "^D"

// SigHandler is an internal function to capture and handle OS signals (eg SIGTERM).
func SigHandler() {
	defer func() {
		if r := recover(); r != nil {
			os.Stderr.WriteString(fmt.Sprintln("Exception caught: ", r))
			SigHandler()
		}
	}()

	//tty.MakeRaw(int(os.Stdout.Fd()))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for {
			sig := <-c
			switch sig.String() {
			//case syscall.SIGTSTP.String():
			//	os.Stderr.WriteString("wubalubub")
			//	os.Exit(2)

			case syscall.SIGINT.String():
				os.Stderr.WriteString(promptInterrupt)
				fallthrough

			case syscall.SIGTERM.String():
				if Prompt == nil {
					os.Exit(0)

				} else {
					proc.ForegroundProc.Kill()

					/*kill := make([]func(), 0)
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
					}*/
				}

			case syscall.SIGQUIT.String():
				if Prompt == nil {
					os.Stderr.WriteString("Murex received SIGQUIT!" + utils.NewLineString)
					os.Exit(2)

				} else {
					os.Stderr.WriteString(promptSIGQUIT)

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
				}

			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}
