package management

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"os"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("source", cmdSource, types.Null)
	lang.DefineFunction(".", cmdSource, types.Null)
}

func cmdSource(p *lang.Process) error {
	var (
		block []rune
		name  string
		err   error
		b     []byte
	)

	if p.IsMethod {
		b, err = p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))
		name = "<stdin>"

	} else {
		block, err = p.Parameters.Block(0)
		if err == nil {
			b = []byte(string(block))
			name = "N/A"

		} else {
			// get block from file
			name, err = p.Parameters.String(0)
			if err != nil {
				return err
			}

			file, err := os.Open(name)
			if err != nil {
				return err
			}

			b, err = io.ReadAll(file)
			if err != nil {
				return err
			}
			block = []rune(string(b))
		}
	}

	hash := ":" + quickHash(name+time.Now().String())
	fileRef := &ref.File{Source: ref.History.AddSource(name, p.FileRef.Source.Module+hash, b)}

	p.RunMode = runmode.Normal
	fork := p.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN)

	fork.Name.Set(p.Name.String())
	fork.FileRef = fileRef
	p.ExitNum, err = fork.Execute(block)
	return err
}

func quickHash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}
