// +build js

package shell

// SignalHandler is an internal function to capture and handle OS signals.
// However since no signals will be sent via a webpage, this is just an empty
// function when compiled against js/wasm
func SignalHandler(_ bool) {}
