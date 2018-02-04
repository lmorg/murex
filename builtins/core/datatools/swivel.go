package datatools

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	proc.GoFunctions["swivel-table"] = cmdSwivelTable
}

func cmdSwivelTable(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	table := make([][]string, 2)
	row := 1

	err := p.Stdin.ReadMap(&proc.GlobalConf, func(key string, val string, eol bool) {
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
