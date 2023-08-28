package types

import "fmt"

func Table2Map(v [][]string, callback func(map[string]any) error) error {
	var (
		err error
		//m      = make(map[string]interface{}, len(v[0]))
		j      int
		recLen = len(v[0])
	)

	for i := 1; i < len(v); i++ {
		if len(v[i]) != recLen {
			return fmt.Errorf("row %d has a different number of records to row 0:\nrow 0: %d (headings)\nrow %d: %d (records)",
				i, recLen, i, len(v[i]))
		}

		m := make(map[string]any)
		for j = 0; j < recLen; j++ {
			m[v[0][j]] = v[i][j]
		}

		err = callback(m)
		if err != nil {
			return err
		}
	}

	return nil
}
