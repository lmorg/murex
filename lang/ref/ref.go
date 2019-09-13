package ref

import (
	"sync"
	"time"
)

// File is used to group the source cursor
type File struct {
	Source *Source
	Line   int
	Column int
}

// Source is a cache of the source file
type Source struct {
	Filename string
	Module   string
	DateTime time.Time
	source   []byte
}

// Equal checks two FileRef sources are the same
func (s Source) Equal(source *Source) bool {
	return s.Module == source.Module &&
		s.Filename == source.Filename &&
		s.DateTime.Equal(source.DateTime)
}

type history struct {
	hist  []*Source
	mutex sync.Mutex
}

// AddSource creates a new ref.Source object and appends it to the source history
func (h *history) AddSource(filename, module string, source []byte) *Source {
	src := &Source{
		Filename: filename,
		Module:   module,
		DateTime: time.Now(),
		source:   source,
	}

	h.mutex.Lock()
	h.hist = append(h.hist, src)
	h.mutex.Unlock()

	return src
}

type dumpVals struct {
	Filename string
	Module   string
	DateTime time.Time
	Source   string
}

func (h *history) Dump() []dumpVals {
	dump := make([]dumpVals, len(h.hist))

	for i, src := range h.hist {
		dump[i] = dumpVals{
			Filename: src.Filename,
			Module:   src.Module,
			DateTime: src.DateTime,
			Source:   string(src.source),
		}
	}

	return dump
}

// History is an array of all the murex source code loaded
var History = new(history)
