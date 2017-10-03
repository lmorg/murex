package apachelogs

import (
	"strings"
	"time"
)

// Parse error log entry. Input is a byte slice rather than string (as used in `ParseApacheLine`) because we need to
// inspect each character and Go's Reader interface returns byte slices anyway.
// `last` is the previous lines timestamp and is used only if the current line doesn't have a timestamp (sometimes that
// happens - it's annoying when it does!)
//
// This code is very new so there's scope for a great deal of optimisation still. However I expect the function
// parameters and returns to remain as is because it follows the same design as `ParseApacheLine` which has been
// stable for a long time now.
func ParseErrorLine(line []byte, last time.Time) (errLog *ErrorLine, err error) {
	errLog = new(ErrorLine)

	var (
		matchBrace bool
		start      int
	)

	for i := range line {

		switch line[i] {
		case ' ':
			continue

		case '[':
			matchBrace = true
			start = i

		case ']':
			if matchBrace == false {
				continue
			}
			matchBrace = false

			if start == 0 {
				errLog.DateTime, err = time.Parse(DateTimeErrorFormat, string(line[1:i]))
				if err == nil {
					errLog.HasTimestamp = true
				} else {
					errLog.Scope = append(errLog.Scope, string(line[1:i]))
					errLog.DateTime = last
				}

				start = i + 1

			} else {
				errLog.Scope = append(errLog.Scope, string(line[start+1:i]))
				start = i + 1

			}

		default:
			if matchBrace == false {
				errLog.Message = strings.TrimSpace(string(line[start:]))
				return
			}
		}
	}

	return
}
