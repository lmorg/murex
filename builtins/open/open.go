package open

import (
	"compress/gzip"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/lmorg/murex/builtins/core/httpclient"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils"
)

var rxExt *regexp.Regexp = regexp.MustCompile(`(?i)\.([a-z]+)(\.gz|)$`)

//type OpenAgent struct {
//	Block         []rune
//	GoFunc        func(process *proc.Process) error `json:"-"`
//	PassPathOrURL bool
//}
//
//var OpenAgents map[string]OpenAgent

func init() {
	proc.GoFunctions["open"] = open

	//OpenAgents = make(map[string]OpenAgent)
}

func open(p *proc.Process) error {
	if p.IsMethod {
		dt := p.Stdin.GetDataType()
		p.Stdout.SetDataType(dt)

		return preview(p.Stdout, p.Stdin, dt)
	}

	filename, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if utils.IsURL(filename) {
		resp, err := httpclient.Request("GET", filename, nil, p.Config, true)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		dt := define.MimeToMurex(resp.Header.Get("Content-Type"))
		p.Stdout.SetDataType(dt)
		return preview(p.Stdout, resp.Body, dt)
	}

	var ext string
	match := rxExt.FindAllStringSubmatch(filename, -1)
	if len(match) > 0 && len(match[0]) > 1 {
		ext = strings.ToLower(match[0][1])
	}

	dt := define.GetExtType(ext)
	//if dt == "" {
	//	p.Stdout.SetDataType(types.Generic)
	//} else {
	p.Stdout.SetDataType(dt)
	//}

	//for _, filename := range p.Parameters.StringArray() {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	if len(filename) > 3 && strings.ToLower(filename[len(filename)-3:]) == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			file.Close()
			return err
		}
		err = preview(p.Stdout, gz, dt)
		file.Close()
		if err != nil {
			return err
		}

	} else {
		err = preview(p.Stdout, file, dt)
		file.Close()
		if err != nil {
			return err
		}

	}
	//}

	return nil
}

func preview(writer io.Writer, reader io.Reader, dt string) (err error) {
	switch dt {

	case "image":
		return pvImage(writer, reader)

	default:
		_, err = io.Copy(writer, reader)
		return err
	}
}
