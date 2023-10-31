//go:build js
// +build js

package term

import (
	"os"
	"sync"
	"syscall/js"

	"github.com/lmorg/murex/utils/readline"
)

var divMutex sync.Mutex

func vtermWrite(r []rune) {
	readline.VTerm.Write(r)

	html := readline.VTerm.ExportHtml()

	divMutex.Lock()

	jsDoc := js.Global().Get("document")
	outElement := jsDoc.Call("getElementById", "term")
	outElement.Set("innerHTML", html)

	divMutex.Unlock()
}

func (t *term) File() *os.File {
	return nil
}
