// +build ignore

/*
	this is only included for reference incase I do build
*/

package render

import (
	ui "github.com/gizak/termui"
	"github.com/lmorg/murex/utils"
	"os"
)

var TermuiEnabled bool

/*func WriteString(s string) {
	divOutput.Text += s
	ui.Render(divOutput)
}*/

func NewGaugeBar(label string) (gauge *ui.Gauge) {
	if !TermuiEnabled {
		return
	}
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	gauge = ui.NewGauge()
	//gauge.Percent = percent
	gauge.Width = ui.TermWidth()
	gauge.Height = 3
	gauge.Float = ui.AlignBottom
	gauge.BorderLabel = label
	gauge.Label = ""
	//gauge.BarColor = ui.ColorRed
	//gauge.BorderFg = ui.ColorWhite
	//gauge.BorderLabelFg = ui.ColorCyan

	/*termui.Body.Rows[0].Height = 10

	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(12, 0, gauge)))

	ui.Body.Align()
	ui.Render(termui.Body)*/
	ui.Render(gauge)

	return
}

func UpdateGaugeBar(gauge *ui.Gauge, percent int, label string) {
	if !TermuiEnabled {
		os.Stderr.WriteString(label + utils.NewLineString)
		return
	}

	gauge.Width = ui.TermWidth()
	gauge.Percent = percent
	gauge.Label = label

	//termui.Body.Align()
	//termui.Render(termui.Body)
	ui.Render(gauge)
}
