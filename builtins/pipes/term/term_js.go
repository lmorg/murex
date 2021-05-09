// +build js

package term

import (
	"sync"
	"syscall/js"

	"github.com/lmorg/murex/utils/virtualterm"
)

var (
	vterm      = virtualterm.NewTerminal(120, 25)
	vtermMutex sync.Mutex
)

func init() {
	vterm.LfIncCr = true
}

func vtermWrite(r []rune) {
	vtermMutex.Lock()

	vterm.Write(r)

	html := vterm.ExportHtml()

	jsDoc := js.Global().Get("document")
	outElement := jsDoc.Call("getElementById", "term")
	outElement.Set("innerHTML", html)

	vtermMutex.Unlock()
}
