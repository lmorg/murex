package open

import (
	"compress/gzip"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

var rxExt = regexp.MustCompile(`(?i)\.([a-z0-9]+)(\.gz)?$`)

func init() {
	//lang.GoFunctions["open"] = open
	lang.DefineFunction("open", open, types.Any)
}

func open(p *lang.Process) (err error) {
	var (
		ext      string
		dataType string
	)

	if p.IsMethod {
		dataType = p.Stdin.GetDataType()
		p.Stdout.SetDataType(dataType)

		ext = getExt("", dataType)
		tmp, err := utils.NewTempFile(p.Stdin, ext)
		defer tmp.Close()

		if err != nil {
			return err
		}

		return preview(p, tmp.FileName, dataType)
	}

	path, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	switch {
	case utils.IsURL(path):
		var body io.ReadCloser
		body, dataType, err = http(p, path)
		if err != nil {
			return err
		}

		ext = getExt("", dataType)
		tmp, err := utils.NewTempFile(body, ext)
		if err != nil {
			return err
		}

		path = tmp.FileName

	default:
		ext = getExt(path, "")
		dataType = lang.GetExtType(ext)
	}

	if dataType == "gz" || (len(path) > 3 && strings.ToLower(path[len(path)-3:]) == ".gz") {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		gz, err := gzip.NewReader(file)
		if err != nil {
			return err
		}
		defer gz.Close()

		ext = getExt(path, "")
		dataType = lang.GetExtType(ext)
		tmp, err := utils.NewTempFile(gz, ext)
		defer tmp.Close()

		if err != nil {
			return err
		}

		path = tmp.FileName
	}

	return preview(p, path, dataType)
}

func getExt(path, dataType string) string {
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
