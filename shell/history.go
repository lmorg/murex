package shell

import (
	"bufio"
	"encoding/json"
	"github.com/lmorg/murex/lang/proc/streams"
	"os"
	"strings"
	"time"
)

type history struct {
	filename string
	Last     string
	List     []HistItem
	writer   streams.Io
}

type HistItem struct {
	DateTime time.Time
	Block    string
}

func newHist(filename string) (h history, err error) {
	h.filename = filename

	h.List, _ = openHist(filename)
	/*// On this one rare occasion we don't care about errors.
	if err != nil {
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
	}*/

	h.writer, err = streams.NewFile(filename)
	return h, err
}

func openHist(filename string) (list []HistItem, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return list, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var item HistItem
		err := json.Unmarshal(scanner.Bytes(), &item)
		if err != nil || len(item.Block) == 0 {
			continue
		}
		list = append(list, item)
		if Instance != nil {
			Instance.SaveHistory(strings.Replace(item.Block, "\n", " ", -1))
		}
	}
	return list, nil
}

func (h *history) Write(block []rune) {
	item := HistItem{
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
