package apachelogs

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Slice indexes after the log has been `strings.Split`.
// This method is faster than using regexp matches.
const (
	sliceStrAccIp        = 0
	sliceStrAccUserId    = 2
	sliceStrAccDateTime  = 3
	sliceStrAccMethod    = 5
	sliceStrAccUri       = 6
	sliceStrAccProtocol  = 7
	sliceStrAccStatus    = 8
	sliceStrAccSize      = 9
	sliceStrAccReferrer  = 10
	sliceStrAccUserAgent = 11
)

// regexp match string and slice indexes.
// This method is slower but more accurate, so this is only used as a fallback if the string splitting fails.
const (
	accessLogFormat  = `^(.*?) (.*?) (.*?) \[(.*?) \+[0-9]{4}\] "(.*?)" ([\-0-9]+) ([\-0-9]+) "(.*?)" "(.*?)"`
	sliceRxIp        = 1
	sliceRxUserId    = 3
	sliceRxDateTime  = 4
	sliceRxRequest   = 5
	sliceRxStatus    = 6
	sliceRxSize      = 7
	sliceRxReferrer  = 8
	sliceRxUserAgent = 9
)

var rxAccessFormat *regexp.Regexp = regexpCompile(accessLogFormat)

// Parse an access log line and return it as the `AccessLine` structure.
// `matched` refers to the pattern matcher. This is also where the operator structures from this package become relevant.
// See Firesword for a working example of this: https://github.com/lmorg/firesword
func ParseAccessLine(line string) (accLog *AccessLine, err error, matched bool) {
	// Quick strings.Split parser
	accLog, err = parserStringSplit(strings.Split(line, " "))

	// Quick parse failed, falling back to slower regexp parser
	if err != nil || accLog.Status.I == 0 {
		accLog, err = parserRegexpSplit(rxAccessFormat.FindStringSubmatch(line))
	}

	if err == nil {
		matched, err = PatternMatch(accLog)
	}

	return
}

// Internal function: `strings.Split` the access log. It's faster than regexp.
func parserStringSplit(split []string) (accLog *AccessLine, err error) {
	accLog = new(AccessLine)
	defer func() (line *AccessLine, err error) {
		// Catch any unforeseen errors that might cause a panic. Mostly just to catch any out-of-bounds errors when
		// working with slices, all of which should already be accounted for but this should cover any human error on
		// my part.
		if r := recover(); r != nil {
			err = errors.New("panic caught in string split parser")
		}
		return
	}()

	if len(split) < 3 {
		err = errors.New(fmt.Sprint("len(split) < 3", split))
		return
	}

	if len(split) >= 9 {
		accLog.DateTime, err = time.Parse(DateTimeAccessFormat, split[sliceStrAccDateTime][1:])

		if err != nil {
			return
		}

	} else {
		err = errors.New(fmt.Sprint("len(split) < 9", split))
		return
	}

	if len(split) < sliceStrAccStatus {
		err = errors.New(fmt.Sprint("len(split) < sliceStrAccStatus", split))
		return
	}
	if len(split) < sliceStrAccReferrer {
		err = errors.New(fmt.Sprint("len(split) < sliceStrAccReferrer", split))
		return
	}
	if split[sliceStrAccStatus] == `"-"` {
		err = errors.New("empty request (typically 408)")
		return

	} else {
		accLog.IP = split[sliceStrAccIp]
		accLog.Method = split[sliceStrAccMethod][1:]
		uri := strings.SplitN(split[sliceStrAccUri], "?", 2)
		accLog.URI = uri[0]
		if len(uri) == 2 {
			accLog.QueryString = "?" + uri[1]
		}
		accLog.Protocol = split[sliceStrAccProtocol][:len(split[sliceStrAccProtocol])-1]
		accLog.Size, _ = strconv.Atoi(split[sliceStrAccSize])
		accLog.Referrer = split[sliceStrAccReferrer][1 : len(split[sliceStrAccReferrer])-1]
		accLog.UserID = split[sliceStrAccUserId]

		pos := len(split) - 1
		for ; pos >= sliceStrAccUserAgent; pos-- {
			if split[pos][len(split[pos])-1:] == `"` {
				break
			}
		}
		if sliceStrAccUserAgent != pos {
			accLog.UserAgent = strings.Join(split[sliceStrAccUserAgent:pos+1], " ")
		} else {
			accLog.UserAgent = split[sliceStrAccUserAgent]
		}
		accLog.UserAgent = accLog.UserAgent[1 : len(accLog.UserAgent)-1]

		// Get processing time. This isn't part of the standard Apache combined log format but it is used in by a
		// couple of other solutions I've written: Level 10 Fireball and Bronze Dagger: https://github.com/lmorg
		if pos+1 < len(split) {
			accLog.ProcTime, _ = strconv.Atoi(split[pos+1])
		}

		accLog.Status = NewStatus(split[sliceStrAccStatus])
	}

	return
}

// Internal function: `regexp.FindStringSubmatch`. Accurate but much _much_ slower.
func parserRegexpSplit(split []string) (accLog *AccessLine, err error) {
	accLog = new(AccessLine)
	defer func() (line *AccessLine, err error) {
		if r := recover(); r != nil {
			err = errors.New("panic caught in regexp parser")
		}
		return
	}()

	if len(split) < sliceRxDateTime {
		err = errors.New(fmt.Sprintf("len(split){%d} < sliceRxDateTime: %s", len(split), split))
		return
	}

	accLog.DateTime, err = time.Parse(DateTimeAccessFormat, split[sliceRxDateTime])
	if err != nil {
		return
	}

	if len(split) < sliceRxUserAgent {
		err = errors.New(fmt.Sprintf("len(split){%d} < sliceRxUserAgent: %s", len(split), split))
		return
	}

	accLog.IP = split[sliceRxIp]
	accLog.Status = NewStatus(split[sliceRxStatus])
	accLog.Size, _ = strconv.Atoi(split[sliceRxSize])
	accLog.Referrer = split[sliceRxReferrer]
	accLog.UserAgent = split[sliceRxUserAgent]
	accLog.UserID = split[sliceRxUserId]

	request := strings.Split(split[sliceRxRequest], " ")

	switch len(request) {
	case 0:
		accLog.Method = "???"
		accLog.URI = "???"
		accLog.Protocol = "???"
	case 1:
		accLog.Method = "???"
		accLog.Protocol = "???"

		// function unrolled for optimisation reasons
		uri := strings.SplitN(request[0], "?", 2)
		accLog.URI = uri[0]
		if len(uri) == 2 {
			accLog.QueryString = "?" + uri[1]
		}
	case 2:
		accLog.Method = request[0]
		accLog.Protocol = "???"

		// function unrolled for optimisation reasons
		uri := strings.SplitN(request[1], "?", 2)
		accLog.URI = uri[0]
		if len(uri) == 2 {
			accLog.QueryString = "?" + uri[1]
		}
	case 3:
		accLog.Method = request[0]
		accLog.Protocol = request[2]

		// function unrolled for optimisation reasons
		uri := strings.SplitN(request[1], "?", 2)
		accLog.URI = uri[0]
		if len(uri) == 2 {
			accLog.QueryString = "?" + uri[1]
		}
	default:
		accLog.Method = request[0]
		accLog.Protocol = request[len(request)-1]

		s := strings.Join(request[1:len(request)-1], " ")

		// function unrolled for optimisation reasons
		uri := strings.SplitN(s, "?", 2)
		accLog.URI = uri[0]
		if len(uri) == 2 {
			accLog.QueryString = "?" + uri[1]
		}
	}

	return
}
