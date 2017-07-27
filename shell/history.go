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
	List     []histItem
	writer   streams.Io
}

type histItem struct {
	Index    int
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
		if Instance != nil {
			Instance.SaveHistory(strings.Replace(item.Block, "\n", " ", -1))
		}
	}
	return list, nil
}

func (h *history) Write(block []rune) {
	item := histItem{
		DateTime: time.Now(),
		Block:    string(block),
		Index:    len(h.List),
	}
	b, _ := json.Marshal(item)
	h.writer.Writeln(b)
	h.List = append(h.List, item)
}

func (h *history) Close() {
	if h.writer != nil {
		h.writer.Close()
	}
}
