package csvbad

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readMap(read stdio.Io, config *config.Config, callback func(*stdio.Map)) error {
	csvParser, err := NewParser(read, config)
	if err != nil {
		return err
	}

	err = csvParser.ReadLine(func(records []string, headings []string) {
		for i := range records {
			callback(&stdio.Map{
				Key:      headings[i],
				Value:    records[i],
				DataType: types.String,
				Last:     i == len(records)-1,
			})
		}
	})

	return err
}
