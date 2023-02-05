package ansititle

import "github.com/lmorg/murex/lang/tty"

func write(ansi []byte) error {
	if ansi == nil {
		return nil
	}

	_, err := tty.Stdout.Write(ansi)
	return err
}
