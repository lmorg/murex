package history

import (
	"bufio"
	"encoding/json"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/readline"
	"os"
	"strings"
	"time"
)

// History exports common functions needed for shell history
type History struct {
	filename string
	Last     string
	List     []histItem
	Writer   stdio.Io
	shell    *readline.Instance
}

type histItem struct {
	Index    int
	DateTime time.Time
	Block    string
}

// New creates a History object
func New(filename string, shell *readline.Instance) (h *History, err error) {
	h = new(History)
	h.filename = filename
	h.List, _ = openHist(filename, shell)
	h.Writer, err = streams.NewFile(filename)
	return h, err
}

func openHist(filename string, shell *readline.Instance) (list []histItem, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return list, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var item histItem
		err := json.Unmarshal(scanner.Bytes(), &item)
		if err != nil || len(item.Block) == 0 {
			continue
		}
		item.Index = len(list)
		list = append(list, item)

		if shell != nil {
			shell.SaveHistory(strings.Replace(item.Block, "\n", " ", -1))
		}
	}
	return list, nil
}

// Write item to history file. eg ~/.murex_history
func (h *History) Write(block string) {
	item := histItem{
		DateTime: time.Now(),
		Block:    block,
		Index:    len(h.List),
	}
	b, _ := json.Marshal(item)
	h.List = append(h.List, item)

	type ws struct {
		DateTime time.Time
		Block    string
	}
	var w ws
	w.Block = item.Block
	w.DateTime = item.DateTime
	b, _ = json.Marshal(w)
	h.Writer.Writeln(b)
}

// Close history file
func (h *History) Close() {
	if h.Writer != nil {
		h.Writer.Close()
	}
}
