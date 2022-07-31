package io

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
	"github.com/phayes/permbits"
)

func init() {
	lang.DefineFunction("g", cmdLsG, types.Json)
	lang.DefineFunction("!g", cmdLsNotG, types.Json)
	lang.DefineFunction("rx", cmdLsRx, types.Json)
	lang.DefineFunction("!rx", cmdLsRx, types.Json)
	lang.DefineMethod("f", cmdLsF, types.ReadArray, types.Json)
	//lang.DefineMethod("!f", cmdLsF, types.ReadArray, types.Json)
}

func cmdLsG(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	glob := p.Parameters.StringAll()

	files, err := filepath.Glob(glob)
	if err != nil {
		return
	}

	j, err := json.Marshal(files, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsNotG(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)

	glob, err := filepath.Glob(p.Parameters.StringAll())
	if err != nil {
		return
	}

	all, err := filepath.Glob("*")
	if err != nil {
		return
	}

	var files []string
	for i := range all {
		if !lists.Match(glob, all[i]) {
			files = append(files, all[i])
		}
	}

	j, err := json.Marshal(files, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsRx(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	rx, err := regexp.Compile(p.Parameters.StringAll())
	if err != nil {
		return
	}

	files, err := filepath.Glob("*")
	if err != nil {
		return
	}

	var matched []string
	for i := range files {
		if rx.MatchString(files[i]) != p.IsNot {
			matched = append(matched, files[i])
		}
	}

	j, err := json.Marshal(matched, p.Stdout.IsTTY())
	if err != nil {
		return
	}

	_, err = p.Stdout.Writeln(j)
	return
}

func cmdLsF(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	var (
		add, remove fFlagsT
	)

	params := p.Parameters.RuneArray()
	if len(params) == 0 {
		return errors.New("missing parameters")
	}

	for _, r := range params {
		add, remove, err = fFlagsParser(r, add, remove)
		if err != nil {
			return err
		}
		/*case "-h":
		p.ExitNum = 2
		usage := []byte("Usage:\n  +f   include files\n  +d   include directories\n  -s   exclude symlinks")
		p.Stderr.Writeln(usage)
		return nil*/
	}

	var files, matched []string

	if p.IsMethod {
		p.Stdin.ReadArray(func(b []byte) {
			files = append(files, string(b))
		})

	} else {
		files, err = filepath.Glob("*")
		if err != nil {
			return
		}
	}

	for i := range files {
		if p.HasCancelled() {
			break
		}

		info, err := os.Stat(files[i])
		if err != nil {
			continue
		}
		mode := info.Mode()
		perm := permbits.FileMode(mode)

		if ((add.File() && mode.IsRegular()) ||
			(add.Dir() && mode.IsDir()) ||
			(add.Symlink() && mode&os.ModeSymlink != 0) ||
			(add.DevBlock() && mode&os.ModeDevice != 0) ||
			(add.DevChar() && mode&os.ModeCharDevice != 0) ||
			(add.Socket() && mode&os.ModeSocket != 0) ||
			(add.NamedPipe() && mode&os.ModeNamedPipe != 0) ||

			(add.UserRead() && perm.UserRead()) ||
			(add.GroupRead() && perm.GroupRead()) ||
			(add.OtherRead() && perm.OtherRead()) ||

			(add.UserWrite() && perm.UserWrite()) ||
			(add.GroupWrite() && perm.GroupWrite()) ||
			(add.OtherWrite() && perm.OtherWrite()) ||

			(add.UserExecute() && perm.UserExecute()) ||
			(add.GroupExecute() && perm.GroupExecute()) ||
			(add.OtherExecute() && perm.OtherExecute()) ||

			(add.SetUid() && mode&os.ModeSetuid != 0) ||
			(add.SetGid() && mode&os.ModeSetgid != 0) ||
			(add.Sticky() && mode&os.ModeSticky != 0) ||

			(add.Irregular() && mode&os.ModeIrregular != 0)) &&

			!((remove.File() && mode.IsRegular()) ||
				(remove.Dir() && mode.IsDir()) ||
				(remove.Symlink() && mode&os.ModeSymlink != 0) ||
				(remove.DevBlock() && mode&os.ModeDevice != 0) ||
				(remove.DevChar() && mode&os.ModeCharDevice != 0) ||
				(remove.Socket() && mode&os.ModeSocket != 0) ||
				(remove.NamedPipe() && mode&os.ModeNamedPipe != 0) ||

				(remove.UserRead() && perm.UserRead()) ||
				(remove.GroupRead() && perm.GroupRead()) ||
				(remove.OtherRead() && perm.OtherRead()) ||

				(remove.UserWrite() && perm.UserWrite()) ||
				(remove.GroupWrite() && perm.GroupWrite()) ||
				(remove.OtherWrite() && perm.OtherWrite()) ||

				(remove.UserExecute() && perm.UserExecute()) ||
				(remove.GroupExecute() && perm.GroupExecute()) ||
				(remove.OtherExecute() && perm.OtherExecute()) ||

				(remove.SetUid() && mode&os.ModeSetuid != 0) ||
				(remove.SetGid() && mode&os.ModeSetgid != 0) ||
				(remove.Sticky() && mode&os.ModeSticky != 0) ||

				(remove.Irregular() && mode&os.ModeIrregular != 0)) {

			matched = append(matched, files[i])
		}
	}

	var b []byte
	b, err = json.Marshal(matched, p.Stdout.IsTTY())
	if err == nil {
		_, err = p.Stdout.Writeln(b)
	}

	return
}
