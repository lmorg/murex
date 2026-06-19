package mxtty

import (
	"encoding/json"
	"fmt"
	"os"
)

func IsTtyphoon() bool {
	return os.Getenv("MXTTY") == "true"
}

const (
	seqApc = "\x1b_"
	seqST  = "\x1b\\"
)

func WriteApcJson(op, key string, value any) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = fmt.Printf("%s%s;%s;%s%s", seqApc, op, key, string(b), seqST)
	return err
}
