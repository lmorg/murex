package streams

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"io"
)

// This structure exists as a null interface for named pipes

type Null struct{}

func (t *Null) MakeParent()                                              {}
func (t *Null) UnmakeParent()                                            {}
func (t *Null) MakePipe()                                                {}
func (t *Null) Read([]byte) (int, error)                                 { return 0, io.EOF }
func (t *Null) ReadLine(func([]byte)) error                              { return nil }
func (t *Null) ReadArray(func([]byte)) error                             { return nil }
func (t *Null) ReadMap(*config.Config, func(string, string, bool)) error { return nil }
func (t *Null) ReadAll() ([]byte, error)                                 { return []byte{}, nil }
func (t *Null) WriteTo(io.Writer) (int64, error)                         { return 0, io.EOF }
func (t *Null) Write(b []byte) (int, error)                              { return len(b), nil }
func (t *Null) Writeln(b []byte) (int, error)                            { return len(b), nil }
func (t *Null) Stats() (uint64, uint64)                                  { return 0, 0 }
func (t *Null) GetDataType() string                                      { return types.Null }
func (t *Null) SetDataType(string)                                       {}
func (t *Null) DefaultDataType(bool)                                     {}
func (t *Null) IsTTY() bool                                              { return false }
func (t *Null) Close()                                                   {}
