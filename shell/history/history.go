package history

import (
	"errors"

	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/stdio"

	"bufio"
	"encoding/json"
	"os"
	"strings"
	"time"
)

// History exports common functions needed for shell history
type History struct {
	filename string
	//Last     string
	list   []histItem
	writer stdio.Io
}

type histItem struct {
	Index    int
	DateTime time.Time
	Block    string
}

// New creates a History object
func New(filename string) (h *History, err error) {
	h = new(History)
	h.filename = filename
	h.list, _ = openHist(filename)
	h.writer, err = streams.NewFile(filename)
	return h, err
}

func openHist(filename string) (list []histItem, err error) {
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

		/*if shell != nil {
			shell.SaveHistory(strings.Replace(item.Block, "\n", " ", -1))
		}*/
	}
	return list, nil
}

// Write item to history file. eg ~/.murex_history
func (h *History) Write(s string) (int, error) {
	block := strings.TrimSpace(s)

	item := histItem{
		DateTime: time.Now(),
		Block:    block,
		Index:    len(h.list),
	}

	if len(h.list) > 0 && h.list[len(h.list)-1].Block != block {
		h.list = append(h.list, item)
	}

	b, err := json.Marshal(item)
	if err != nil {
		return len(h.list), err
	}

	type ws struct {
		DateTime time.Time
		Block    string
	}
	var w ws
	w.Block = item.Block
	w.DateTime = item.DateTime
	b, err = json.Marshal(w)
	if err != nil {
		return len(h.list), err
	}

	_, err = h.writer.Writeln(b)
	return len(h.list), err
}

/*// Close history file
func (h *History) Close() {
	if h.Writer != nil {
		h.Writer.Close()
	}
}*/

func (h *History) GetLine(i int) (string, error) {
	if i < 0 {
		return "", errors.New("Cannot use a negative index when requsting historic commands")
	}
	if i < len(h.list) {
		return h.list[i].Block, nil
	}
	return "", errors.New("Index requested greater than number of items in history")
}

func (h *History) Len() int {
	return len(h.list)
}

func (h *History) Dump() interface{} {
	return h.list
}
