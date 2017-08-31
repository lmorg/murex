package preview

import (
	"compress/gzip"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/data"
	"io"
	"os"
	"regexp"
	"strings"
)

var rxExt *regexp.Regexp = regexp.MustCompile(`\.([a-zA-Z]+)(\.gz|)$`)

func init() {
	proc.GoFunctions["open"] = open
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

	var ext string
	match := rxExt.FindAllStringSubmatch(filename, -1)
	if len(match) > 0 && len(match[0]) > 1 {
		ext = strings.ToLower(match[0][1])
	}

	dt := data.GetExtType(ext)
	if dt == "" {
		p.Stdout.SetDataType(types.Generic)
	} else {
		p.Stdout.SetDataType(dt)
	}

	//for _, filename := range p.Parameters.StringArray() {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
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
