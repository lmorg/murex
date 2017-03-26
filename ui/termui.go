// +build ignore

/*
	this is only included for reference incase I do build
*/

package ui

import (
	//ui "github.com/gizak/termui"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/ui/render"
)

var block string = `text(test/100.txt) getfile(http://www.mirrorservice.org/sites/www.linuxmint.com/pub/linuxmint.com/stable/18.1/linuxmint-18.1-xfce-64bit.iso).regex(s/.*//)`

// This will have the interactive shell code
func Start() {

	render.TermuiEnabled = true

	//render.NewOutput()
	//render.WriteString("Press a or q\n")
	//par.Width = ui.TermWidth()
	//ui.Body.AddRows(
	//	ui.NewRow(ui.NewCol(12, 0, par)),
	//)

	//ui.Body.Align()
	//ui.Render(ui.Body)
	/*p := ui.NewPar(":PRESS q TO QUIT DEMO")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Text Box"
	p.BorderFg = ui.ColorCyan
	*/

	//	ui.Render(p, g) // feel free to call Render, it's async and non-block

	//ui.Body.Align()
	//ui.Render(ui.Body)

	//ui.Handle("/sys/kbd/a", func(ui.Event) {
	lang.ProcessNewBlock([]byte(block))
	//})

	/*ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})*/

	/*ui.Loop()*/
}
