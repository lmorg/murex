package csv

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readMap(read stdio.Io, config *config.Config, callback func(*stdio.Map)) error {
	reader := csv.NewReader(read)

	v, err := config.Get("csv", "separator", types.String)
	if err != nil {
		return err
	}
	if len(v.(string)) > 0 {
		reader.Comma = rune(v.(string)[0])
	}

	v, err = config.Get("csv", "comment", types.String)
	if err != nil {
		return err
	}
	if len(v.(string)) > 0 {
		reader.Comment = rune(v.(string)[0])
	}

	for {
		recs, err := reader.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		for i := range recs {
			callback(&stdio.Map{
				Key:      strconv.Itoa(i),
				Value:    recs[i],
				DataType: types.String,
				Last:     i == len(recs)-1,
			})
		}
	}
}
