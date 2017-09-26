package csv

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
)

func readMap(read stdio.Io, config *config.Config, callback func(key, value string, last bool)) error {
	csvParser, err := NewParser(read, config)
	if err != nil {
		return err
	}

	err = csvParser.ReadLine(func(records []string, headings []string) {
		for i := range records {
			callback(headings[i], records[i], i == len(records)-1)
		}
	})

	return err
}
