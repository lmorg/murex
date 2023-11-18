package open

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

var rxExt = regexp.MustCompile(`(?i)\.([a-z0-9]+)(\.gz)?$`)

func init() {
	lang.DefineFunction("open", open, types.Any)
}

func open(p *lang.Process) (err error) {
	var dataType string

	if p.IsMethod {
		return OpenPipe(p, p.Stdin)
	}

	path, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	closers, err := OpenFile(p, &path, &dataType)
	if err != nil {
		return err
	}

	err = preview(p, path, dataType)
	if err != nil {
		return err
	}

	return CloseFiles(closers)
}

func OpenPipe(p *lang.Process, pipe stdio.Io) error {
	dataType := pipe.GetDataType()
	p.Stdout.SetDataType(dataType)

	ext := GetExt("", dataType)
	tmp, err := utils.NewTempFile(p.Stdin, ext)
	if err != nil {
		return err
	}
	defer tmp.Close()

	return preview(p, tmp.FileName, dataType)
}

func OpenFile(p *lang.Process, path *string, dataType *string) ([]io.Closer, error) {
	var (
		closers []io.Closer
		err     error
		ext     string
	)

	switch {
	case utils.IsURL(*path):
		var body io.ReadCloser
		body, *dataType, err = http(p, *path)
		if err != nil {
			return closers, err
		}

		ext = GetExt("", *dataType)
		tmp, err := utils.NewTempFile(body, ext)
		if err != nil {
			return closers, err
		}

		*path = tmp.FileName

	default:
		ext = GetExt(*path, "")
		*dataType = lang.GetExtType(ext)
	}

	if *dataType == "gz" || (len(*path) > 3 && strings.ToLower((*path)[len(*path)-3:]) == ".gz") {
		file, err := os.Open(*path)
		if err != nil {
			return closers, err
		}
		//defer file.Close()
		closers = append(closers, file)

		gz, err := gzip.NewReader(file)
		if err != nil {
			return closers, err
		}
		//defer gz.Close()
		closers = append(closers, gz)

		ext = GetExt(*path, "")
		*dataType = lang.GetExtType(ext)
		tmp, err := utils.NewTempFile(gz, ext)
		//defer tmp.Close()
		closers = append(closers, tmp)

		if err != nil {
			return closers, err
		}

		*path = tmp.FileName
	}

	return closers, err
}

func CloseFiles(closers []io.Closer) error {
	var s string

	for i := len(closers) - 1; i > -1; i-- {
		err := closers[i].Close()
		if err != nil {
			s = fmt.Sprintf("%s: %s", err.Error(), s)
		}
	}

	if len(s) > 0 {
		return fmt.Errorf("unable to close files: %s", s[:len(s)-2])
	}

	return nil
}

func GetExt(path, dataType string) string {
	if path != "" {
		match := rxExt.FindAllStringSubmatch(path, -1)
		if len(match) > 0 && len(match[0]) > 1 {
			return strings.ToLower(match[0][1])
		}
	}

	m := lang.GetFileExts()
	for ext := range m {
		if m[ext] == dataType {
			return ext
		}
	}

	return ""
}

func preview(p *lang.Process, path, dataType string) error {
	if dataType == "" {
		dataType = types.Generic
	}

	p.Stdout.SetDataType(dataType)
	agent, err := OpenAgents.Get(dataType)

	// we check if std(in|err) is a TTY because stdout might be a fake TTY in preview pane
	if (p.Stdout.IsTTY() || (p.Stdin.IsTTY() && p.Stderr.IsTTY())) && dataType == types.Generic {
		return openSystemCommand(p, path)
	}

	if !p.Stdout.IsTTY() || err != nil {
		// Not a TTY or no open agent exists so fallback to passing []bytes along
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		_, err = io.Copy(p.Stdout, file)
		return err
	}

	fork := p.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN)
	fork.Name.Set("open")
	fork.Parameters.DefineParsed([]string{path})
	fork.FileRef = agent.FileRef
	_, err = fork.Execute(agent.Block)

	if err != nil {
		p.Stderr.Writeln([]byte("`open` code could not compile: " + err.Error()))
	}

	return err
}
