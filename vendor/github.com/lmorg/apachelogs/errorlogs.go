package apachelogs

import (
	"sort"
	"time"
)

// timestamp formatting in Apache error logs.
// You shouldn't need to alter this but it's provided as a variable just in case.
var DateTimeErrorFormat string = "Mon Jan 02 15:04:05 2006"

// A broken down record of each field in an error log.
type ErrorLine struct {
	DateTime     time.Time
	HasTimestamp bool // Sometimes log file entries don't have a timestamp
	Scope        []string
	Message      string
	FileName     string
}

// This type is provided to offer easy sorting of slices of `ErrorLine`
type ErrorLog []ErrorLine

func (e ErrorLog) Remove(index int)   { e = append(e[:index], e[index+1:]...) }
func (e ErrorLog) SortByDateTime()    { sort.Sort(ErrorLog(e)) }
func (e ErrorLog) Len() int           { return len(e) }
func (e ErrorLog) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e ErrorLog) Less(i, j int) bool { return e[i].DateTime.Before(e[j].DateTime) }
