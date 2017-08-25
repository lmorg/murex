package pipes

//go:generate stringer -type=PipeTypes

// PipeTypes is an ID for each stream.Io interface.
// However this might get rewritten to support adding interfaces via builtins
type PipeTypes int

const (
	pipeUndefined PipeTypes = iota
	pipeNull
	pipeStream
	pipeFileWriter
	pipeNetDialer
	pipeNetListener
)
