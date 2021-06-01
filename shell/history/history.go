package history

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/lmorg/murex/builtins/pipes/file"
	"github.com/lmorg/murex/lang/stdio"
)

// History exports common functions needed for shell history
type History struct {
	filename string
	list     []Item
	writer   stdio.Io
}

// Item is the structure of an individual item in the History.list slice
type Item struct {
	Index    int
	DateTime time.Time
	Block    string
}

// New creates a History object
func New(filename string) (h *History, err error) {
	h = new(History)
	h.filename = filename
	h.list, _ = openHist(filename)
	h.writer, err = file.NewFile(filename)

	return h, err
}

func openHist(filename string) (list []Item, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return list, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var item Item
		err := json.Unmarshal(scanner.Bytes(), &item)
		if err != nil || len(item.Block) == 0 {
			continue
		}
		item.Index = len(list)
		list = append(list, item)
	}
	return list, nil
}

// Write item to history file. eg ~/.murex_history
func (h *History) Write(s string) (int, error) {
	block := strings.TrimSpace(s)

	type jsonline struct {
		DateTime time.Time `json:"datetime"`
		Block    string    `json:"block"`
	}

	item := Item{
		DateTime: time.Now(),
		Block:    block,
		Index:    len(h.list),
	}

	if len(h.list) == 0 || h.list[len(h.list)-1].Block != block {
		h.list = append(h.list, item)
	}

	line := jsonline{
		Block:    block,
		DateTime: item.DateTime,
	}

	b, err := json.Marshal(line)
	if err != nil {
		return h.Len(), err
	}

	_, err = h.writer.Writeln(b)
	return h.Len(), err
}

/*// Close history file
func (h *History) Close() {
	if h.Writer != nil {
		h.Writer.Close()
	}
}*/

// GetLine returns a specific line from the history file
func (h *History) GetLine(i int) (string, error) {
	if i < 0 {
		return "", errors.New("Cannot use a negative index when requesting historic commands")
	}
	if i < len(h.list) {
		return h.list[i].Block, nil
	}
	return "", errors.New("Index requested greater than number of items in history")
}

// Len returns the number of items in the history file
func (h *History) Len() int {
	return len(h.list)
}

// Dump returns the entire history file
func (h *History) Dump() interface{} {
	return h.list
}
