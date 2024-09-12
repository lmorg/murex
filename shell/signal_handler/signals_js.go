//go:build js
// +build js

package signalhandler

// EventLoop is an internal function to capture and handle OS signals.
// However since no signals will be sent via a webpage, this is just an empty
// function when compiled against js/wasm
func EventLoop(_ bool) {}

func Register(_ bool) {}

var Handlers *SignalFunctionsT

type SignalFunctionsT struct{}
