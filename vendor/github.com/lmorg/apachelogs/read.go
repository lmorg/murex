package apachelogs

import (
	"bufio"
	"compress/gzip"
	"os"
	"time"
)

// This function isn't required to use this package. It's just a helper function for the lazy.
// "callback" parameter is a function which gets called with each access log entry, so you can handle the entries as you so wish
// "errHandler" parameter is a function which gets called on error, so you can optionally display / hide / whatever errors
func ReadAccessLog(filename string, callback func(accessLine *AccessLine), errHandler func(err error)) {
	var (
		reader *bufio.Reader
		err    error
	)

	fi, err := os.Open(filename)
	if err != nil {
		errHandler(err)
		return
	}
	defer fi.Close()

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		fz, err := gzip.NewReader(fi)
		if err != nil {
			errHandler(err)
			return
		}
		reader = bufio.NewReader(fz)
	} else {
		reader = bufio.NewReader(fi)
	}

	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				errHandler(err)
			}
			break
		}

		line, err, matched := ParseAccessLine(string(b))

		if err != nil {
			errHandler(err)
			continue
		}

		if !matched {
			continue
		}

		line.FileName = filename
		callback(line)
	}

	return
}

// This function isn't required to use this package. It's just a helper function for the lazy.
// "callback" parameter is a function which gets called with each access log entry, so you can handle the entries as you so wish
// "errHandler" parameter is a function which gets called on error, so you can optionally display / hide / whatever errors
func ReadErrorLog(filename string, callback func(errorLine *ErrorLine), errHandler func(err error)) {
	var (
		reader *bufio.Reader
		err    error
		last   time.Time
	)

	fi, err := os.Open(filename)
	if err != nil {
		errHandler(err)
		return
	}
	defer fi.Close()

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		fz, err := gzip.NewReader(fi)
		if err != nil {
			errHandler(err)
			return
		}
		reader = bufio.NewReader(fz)
	} else {
		reader = bufio.NewReader(fi)
	}

	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				errHandler(err)
			}
			break
		}

		line, err := ParseErrorLine(b, last)

		if err != nil {
			errHandler(err)
			continue
		}

		last = line.DateTime
		line.FileName = filename
		callback(line)
	}

	return
}
