package datatools

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	proc.GoFunctions["swivel-table"] = cmdSwivelTable
	proc.GoFunctions["swivel-datatype"] = cmdSwivelDataType
}

func cmdSwivelDataType(p *proc.Process) error {
	//dt := p.Stdin.GetDataType()
	dt, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(dt)

	table := make(map[string][]string)
	var row int

	err = p.Stdin.ReadMap(p.Config, func(key string, val string, eol bool) {
		if len(table[key]) == 0 {
			table[key] = make([]string, 0)
		}

		if len(table[key]) < row+1 {
			table[key] = append(table[key], make([]string, row+1-len(table[key]))...)
		}

		table[key][row] = val

		if eol {
			row++
		}
	})

	b, err := define.MarshalData(p, dt, table)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

// This function is ripe for optimisation. However given it's infrequent nature and small datasets, I'm not in any great
// rush.
func cmdSwivelTable(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	table := make([][]string, 2)
	row := 1

	err := p.Stdin.ReadMap(p.Config, func(key string, val string, eol bool) {
		table[row] = append(table[row], val)
		if len(table[row]) > len(table[0]) {
			table[0] = append(table[0], key)
		}

		if eol {
			row++
			table = append(table, []string{})
		}
	})

	if err != nil {
		return err
	}

	if len(table) > 0 {
		table = table[:len(table)-1]
	}

	b, err := define.MarshalData(p, dt, rotate(table))
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func rotate(t [][]string) [][]string {
	var maxColumns, maxRows int
	for _, row := range t {
		if len(row) > maxColumns {
			maxColumns = len(row)
		}
	}

	maxRows = len(t)
	table := make([][]string, maxColumns)
	for r := range table {
		table[r] = make([]string, maxRows)
	}

	for r, rows := range t {
		for column := range rows {
			table[column][r] = t[r][column]
		}
	}

	return table
}
