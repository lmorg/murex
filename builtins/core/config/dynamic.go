package cmdconfig

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
)

func getDynamic(block []rune, args []string, fileRef *ref.File) func() (interface{}, int, error) {
	return func() (interface{}, int, error) {
		block = block[1 : len(block)-1]

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		fork.Name.Set("config")
		fork.Parameters.DefineParsed(args)
		fork.FileRef = fileRef
		exitNum, err := fork.Execute(block)

		if err != nil {
			return nil, exitNum, fmt.Errorf("dynamic config code could not compile: %s", err.Error())
		}

		b, err := fork.Stdout.ReadAll()
		if err != nil {
			return nil, exitNum, err
		}

		return string(b), exitNum, nil
	}
}

func setDynamic(block []rune, args []string, fileRef *ref.File, dataType string) func(interface{}) (int, error) {
	return func(value interface{}) (int, error) {
		block = block[1 : len(block)-1]
		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_CREATE_STDIN)
		fork.Name.Set("config")
		fork.Parameters.DefineParsed(args)
		fork.FileRef = fileRef
		s, err := types.ConvertGoType(value, types.String)
		if err != nil {
			return 1, err
		}

		fork.Stdin.SetDataType(dataType)
		_, err = fork.Stdin.Write([]byte(s.(string)))
		if err != nil {
			return 1, err
		}

		exitNum, err := fork.Execute(block)
		if err != nil {
			return exitNum, fmt.Errorf("dynamic config code could not compile: %s", err.Error())
		}

		return exitNum, nil
	}
}
