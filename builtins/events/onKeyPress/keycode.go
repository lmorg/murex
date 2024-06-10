package onkeypress

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/readline"
)

func init() {
	lang.DefineFunction("key-code", cmdKeyCodes, types.String)
}

const (
	msgSeq     = "ANSI Constants:   "
	msgRaw     = "Byte Sequence:    "
	msgUnicode = "Contains Unicode: "
)

func cmdKeyCodes(p *lang.Process) error {
	if p.Background.Get() {
		return fmt.Errorf("this builtin does not support running in the background")
	}

	if p.IsMethod {
		return fmt.Errorf("this builtin does not support running as a method")
	}

	os.Stdout.WriteString("Press any key to print its escape constants...\n")

	state, err := readline.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer readline.Restore(int(os.Stdin.Fd()), state)

	b := make([]byte, 1024)
	i, err := os.Stdin.Read(b)
	if err != nil {
		return err
	}

	escaped := ansi.GetConsts(b[:i])

	if p.Stdout.IsTTY() {
		p.Stdout.Write([]byte(msgSeq))
	}

	p.Stdout.Write([]byte(escaped))

	if p.Stdout.IsTTY() {
		p.Stdout.Writeln([]byte(fmt.Sprintf("\n%s%%%v", msgRaw, b[:i])))
		p.Stdout.Writeln([]byte(fmt.Sprintf("%s%v", msgUnicode, utf8.RuneCount(b[:i]) != i)))
	}

	return nil
}
