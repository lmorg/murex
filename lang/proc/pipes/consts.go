package pipes

//go:generate stringer -type=PipeTypes

type PipeTypes int

const (
	pipeUndefined PipeTypes = iota
	pipeNull
	pipeStream
	pipeFileWriter
	pipeNetDialer
	pipeNetListener
)
