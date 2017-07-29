package shell

import (
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SigHandler() {
	defer func() {
		if r := recover(); r != nil {
			os.Stderr.WriteString(fmt.Sprintln("Exception caught: ", r))
			SigHandler()
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGQUIT)
	go func() {
		//killTimer := time.Now().Add(3 * time.Second)
		for {
			sig := <-c
			switch sig.String() {
			case syscall.SIGTERM.String():
				Instance.Terminal.ExitRawMode()
				os.Stderr.WriteString("Shell recieved SIGTERM! Exiting..." + utils.NewLineString)
				os.Exit(1)
			case os.Interrupt.String():
				//doublePress := time.Now().After(killTimer)
				//killTimer = time.Now().Add(3 * time.Second)
				//if !doublePress {
				if Instance == nil {
					//os.Stderr.WriteString("^C")
					go proc.KillForeground()
					time.Sleep(time.Second)
				} else {
					//os.Stderr.WriteString("^KILL")
					fList := proc.GlobalFIDs.ListAll()
					for i := range fList {
						if fList[i].Kill != nil {
							go fList[i].Kill()
						}
					}
				}
			case syscall.SIGQUIT.String():
				os.Stderr.WriteString("Shell recieved SIGQUIT! Exiting..." + utils.NewLineString)
				os.Exit(2)
			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}
