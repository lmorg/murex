package io

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func init() {
	proc.GoFunctions["lockfile"] = cmdLockfile
}

func cmdLockfile(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	method, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	name, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	lockfile := lockFilePath(name)

	switch method {
	case "lock":
		if fileExists(lockfile) {
			return errors.New("Lockfile already exists.")
		}

		file, err := os.Create(lockfile)
		if err != nil {
			return err
		}

		file.WriteString(fmt.Sprintf("%d:%d", os.Getpid(), p.Scope.Id))
		file.Close()

	case "unlock":
		if !fileExists(lockfile) {
			return errors.New("Lockfile does not exist.")
		}
		return os.Remove(lockfile)

	case "wait":
		for {
			if !fileExists(lockfile) {
				return nil
			}
			time.Sleep(100 * time.Millisecond)
		}

	case "path":
		_, err = p.Stdout.Write([]byte(lockfile))
		return err

	default:
		return errors.New("That isn't a valid method: " + method)
	}

	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func lockFilePath(key string) string {
	return consts.TempDir + key + ".lockfile"
}
