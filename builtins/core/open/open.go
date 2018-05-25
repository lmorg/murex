package open

import (
	"compress/gzip"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
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
		dt         string
		readCloser io.ReadCloser
		stdin      stdio.Io
	)

	if p.IsMethod {
		dt = p.Stdin.GetDataType()
		p.Stdout.SetDataType(dt)

		return preview(p, p.Stdin)
	}

	path, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	switch {
	case utils.IsURL(path):
		readCloser, dt, err = http(p, path)

	default:
		dt = define.GetExtType(getExt(path))
		readCloser, err = os.Open(path)
	}

	if err != nil {
		return err
	}

	if dt == "gz" || (len(path) > 3 && strings.ToLower(path[len(path)-3:]) == ".gz") {
		gz, err := gzip.NewReader(readCloser)
		if err != nil {
			return err
		}

		// would this break things?
		defer readCloser.Close()

		dt = define.GetExtType(getExt(path))
		stdin := streams.NewReadCloser(gz)
		stdin.SetDataType(dt)

	} else {
		stdin = streams.NewReadCloser(readCloser)
		stdin.SetDataType(dt)
	}

	return preview(p, stdin)
}

func getExt(path string) string {
	match := rxExt.FindAllStringSubmatch(path, -1)
	if len(match) > 0 && len(match[0]) > 1 {
		return strings.ToLower(match[0][1])
	}
	return ""
}

func preview(p *proc.Process, stdin stdio.Io) error {
	dataType := stdin.GetDataType()

	block, _ := OpenAgents.Get(dataType)

	if !p.Stdout.IsTTY() || len(block) == 0 {
		// Not a TTY or no open agent exists so fallback to passing []bytes along
		_, err := io.Copy(p.Stdout, stdin)
		return err
	}

	branch := proc.ShellProcess.BranchFID()
	defer branch.Close()
	branch.Process.Scope = branch.Process
	branch.Process.Parent = branch.Process
	_, err := lang.RunBlockNewConfigSpace(block, stdin, p.Stdout, p.Stderr, branch.Process)

	if err != nil {
		ansi.Stderrln(p, ansi.FgRed, "`open` code could not compile: "+err.Error())
	}

	return err
}
