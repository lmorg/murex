package shell

import (
	"encoding/json"
	"github.com/lmorg/murex/lang/proc/streams"
	"time"
)

type history struct {
	filename string
	Last     string
	List     []histItem
	writer   streams.Io
}

type histItem struct {
	DateTime time.Time
	Block    string
}

func openHistFile(filename string) (h history, err error) {
	h.filename = filename
	h.writer, err = streams.NewFile(filename)
	return
}

func (h *history) Write(block []rune) {
	item := histItem{
		DateTime: time.Now(),
		Block:    string(block),
	}
	b, _ := json.Marshal(item)
	h.writer.Writeln(b)
}

func (h *history) Close() {
	if h.writer != nil {
		h.writer.Close()
	}
}
