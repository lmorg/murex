package ansititle

import "os"

func write(ansi []byte) error {
	if ansi == nil {
		return nil
	}

	_, err := os.Stdout.Write(ansi)
	return err
}
