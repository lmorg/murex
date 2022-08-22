package hintsummary

import "sync"

type summaryCacheT struct {
	mutex   sync.Mutex
	summary map[string]string
}

func NewSummaryCache() *summaryCacheT {
	sc := new(summaryCacheT)
	sc.summary = make(map[string]string)
	return sc
}

func (sc *summaryCacheT) Get(command string) string {
	sc.mutex.Lock()

	s := sc.summary[command]

	sc.mutex.Unlock()
	return s
}

func (sc *summaryCacheT) Set(command string, summary []rune) {
	if len(summary) == 0 {
		return
	}

	sc.mutex.Lock()
	sc.summary[command] = string(summary)
	sc.mutex.Unlock()
}

func (sc *summaryCacheT) Dump() interface{} {
	dump := make(map[string]string)

	sc.mutex.Lock()
	for k, v := range sc.summary {
		dump[k] = v
	}
	sc.mutex.Unlock()

	return dump
}

var Cache = NewSummaryCache()
