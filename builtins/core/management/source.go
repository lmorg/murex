package management

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/lmorg/murex/config/profile"
	profilepaths "github.com/lmorg/murex/config/profile/paths"
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
	// source from STDIN
	if p.IsMethod {
		return sourceExecReader(p, p.Stdin, "<stdin>", sourceQuickHash(p, "<stdin>"))
	}

	// source from ARGV
	block, err := p.Parameters.Block(0)
	if err == nil {
		return sourceExecBlock(p, block, []byte(string(block)), "$ARGV", sourceQuickHash(p, "$ARGV"))
	}

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	// source from module
	if strings.HasPrefix(name, "module:") {
		return sourceModuleLocate(p, name[7:])
	}

	// source from file
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return sourceExecReader(p, f, name, sourceQuickHash(p, name))
}

func sourceModuleLocate(p *lang.Process, name string) error {
	split := strings.SplitN(name, "/", 2)
	if len(split) != 2 {
		return fmt.Errorf("invalid package/module name '%s'", name)
	}

	path := profilepaths.ModulePath()
	mod, err := profile.LoadPackage(path+"/"+split[0], false, false)
	if err != nil {
		return err
	}

	for i := range mod {
		if mod[i].Name == split[1] {
			return sourceModuleExecute(p, &mod[i])
		}
	}

	return fmt.Errorf("cannot find module '%s'", split[1])
}

func sourceModuleExecute(p *lang.Process, mod *profile.Module) error {
	if mod.Preload != "" {
		f, err := os.Open(mod.PreloadPath())
		if err != nil {
			return fmt.Errorf("cannot open preload script: %v", err)
		}
		defer f.Close()
		err = sourceExecReader(p, f, mod.PreloadPath(), fmt.Sprintf("%s/%s", mod.Package, mod.Name))
		if err != nil {
			return fmt.Errorf("error executing preload script: %v", err)
		}
	}

	f, err := os.Open(mod.Path())
	if err != nil {
		return fmt.Errorf("cannot open module script: %v", err)
	}
	defer f.Close()
	err = sourceExecReader(p, f, mod.Path(), fmt.Sprintf("%s/%s", mod.Package, mod.Name))
	if err != nil {
		return fmt.Errorf("error executing module script: %v", err)
	}
	return nil
}

func sourceExecReader(p *lang.Process, r io.Reader, filename, packageModule string) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	return sourceExecBlock(p, []rune(string(b)), b, filename, packageModule)
}

func sourceExecBlock(p *lang.Process, block []rune, b []byte, filename, packageModule string) error {
	var err error

	fileRef := &ref.File{Source: ref.History.AddSource(filename, packageModule, b)}

	p.RunMode = runmode.Normal // TODO: is this a bug or correct?
	fork := p.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN)

	fork.Name.Set(p.Name.String())
	fork.FileRef = fileRef
	p.ExitNum, err = fork.Execute(block)
	return err
}

func sourceQuickHash(p *lang.Process, name string) string {
	hasher := md5.New()
	hasher.Write([]byte(name + time.Now().String()))
	return p.FileRef.Source.Module + ":" + base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}
