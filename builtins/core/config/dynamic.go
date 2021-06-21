package cmdconfig

import (
	"errors"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func getDynamic(block []rune, args []string, fileRef *ref.File) func() (interface{}, error) {
	return func() (interface{}, error) {
		block = block[1 : len(block)-1]

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		fork.Name.Set("config")
		fork.Parameters.DefineParsed(args)
		fork.FileRef = fileRef
		exitNum, err := fork.Execute(block)

		if err != nil {
			return nil, errors.New("Dynamic config code could not compile: " + err.Error())
		}
		if exitNum != 0 && debug.Enabled {
			lang.ShellProcess.Stderr.Writeln([]byte("Dynamic config returned a none zero exit number." + utils.NewLineString))
		}

		b, err := fork.Stdout.ReadAll()
		if err != nil {
			return nil, err
		}

		return string(b), nil
	}
}

func setDynamic(block []rune, args []string, fileRef *ref.File, dataType string) func(interface{}) error {
	return func(value interface{}) error {
		//if !types.IsBlock([]byte(stringblock)) {
		//	return nil, errors.New("Dynamic config reader is not a code block")
		//}
		block = block[1 : len(block)-1]
		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_CREATE_STDIN)
		fork.Name.Set("config")
		fork.Parameters.DefineParsed(args)
		fork.FileRef = fileRef
		s, err := types.ConvertGoType(value, types.String)
		if err != nil {
			return err
		}

		fork.Stdin.SetDataType(dataType)
		_, err = fork.Stdin.Write([]byte(s.(string)))
		if err != nil {
			return err
		}

		exitNum, err := fork.Execute(block)

		if err != nil {
			return errors.New("Dynamic config code could not compile: " + err.Error())
		}
		if exitNum != 0 && debug.Enabled {
			lang.ShellProcess.Stderr.Writeln([]byte("Dynamic config returned a none zero exit number." + utils.NewLineString))
		}

		return nil
	}
}
