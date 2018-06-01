package open

import (
	"compress/gzip"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
)

var rxExt *regexp.Regexp = regexp.MustCompile(`(?i)\.([a-z]+)(\.gz|)$`)

func init() {
	proc.GoFunctions["open"] = open
}

func open(p *proc.Process) (err error) {
	var (
		ext        string
		dataType   string
		readCloser io.ReadCloser
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
		readCloser, dataType, err = http(p, path)
		ext = getExt("", dataType)

	default:
		ext = getExt(path, "")
		dataType = define.GetExtType(ext)
		readCloser, err = os.Open(path)
	}

	defer readCloser.Close()

	if err != nil {
		return err
	}

	if dataType == "gz" || (len(path) > 3 && strings.ToLower(path[len(path)-3:]) == ".gz") {
		gz, err := gzip.NewReader(readCloser)
		if err != nil {
			return err
		}

		defer gz.Close()

		ext = getExt(path, "")
		dataType = define.GetExtType(ext)
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
	match := rxExt.FindAllStringSubmatch(path, -1)
	if len(match) > 0 && len(match[0]) > 1 {
		return strings.ToLower(match[0][1])
	}

	m := define.GetFileExts()
	for ext := range m {
		if m[ext] == dataType {
			return ext
		}
	}

	return ""
}

func preview(p *proc.Process, path, dataType string) error {
	if dataType == "" {
		dataType = types.Generic
	}

	p.Stdout.SetDataType(dataType)
	block, _ := OpenAgents.Get(dataType)

	if !p.Stdout.IsTTY() || len(block) == 0 {
		// Not a TTY or no open agent exists so fallback to passing []bytes along
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		_, err = io.Copy(p.Stdout, file)
		return err
	}

	branch := proc.ShellProcess.BranchFID()
	defer branch.Close()
	branch.Process.Scope = branch.Process
	branch.Process.Parent = branch.Process
	branch.Process.Name = "open"
	branch.Process.Parameters.Params = []string{path}
	_, err := lang.RunBlockNewConfigSpace(block, nil, p.Stdout, p.Stderr, branch.Process)

	if err != nil {
		ansi.Stderrln(p, ansi.FgRed, "`open` code could not compile: "+err.Error())
	}

	return err
}
