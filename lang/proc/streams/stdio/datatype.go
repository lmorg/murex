package stdio

import (
	"bytes"
	"github.com/lmorg/murex/lang/types"
	"regexp"
)

// StdoutDataType defines whether the data type header should be included in stdout
var StdoutDataType bool

var (
	prefixS    string         = string([]byte{1}) + "mxdt" + string([]byte{2})
	prefixB    []byte         = []byte(prefixS)
	rxImportDt *regexp.Regexp = regexp.MustCompile(prefixS + `([a-z]+)` + string([]byte{3}))
)

// DataTypeHeader returns the data type header if the shell requires to. Otherwise it returns an empty string.
func DataTypeHeader(dt string, force bool) string {
	if StdoutDataType || force {
		return prefixS + dt + string([]byte{3})
	}
	return ""
}

// DataTypeReader is an object for checking the data type header or caching the stdin if the header is incomplete.
type DataTypeReader struct {
	b         []byte
	skipCheck bool
}

// SetDataType reads the data type header and sets the stdio data type if the header is present.
func (dtr *DataTypeReader) SetDataType(stdio Io, stdin []byte, offset int) {
	if dtr.skipCheck {
		return
	}

	b := append(dtr.b, stdin...)

	match := rxImportDt.FindSubmatch(b)
	if len(match) == 2 {
		dtr.b = nil
		dtr.skipCheck = true
		//return string(match[1])
		stdio.SetDataType(string(match[1]))
		return
	}

	if len(b) > 27 {
		dtr.b = nil
		dtr.skipCheck = true
		//return types.Generic
		stdio.SetDataType(types.Generic)
		return
	}

	if len(b) <= 6 && bytes.Equal(b, prefixB[:len(b)]) {
		dtr.b = b
		return
	}

	if len(b) > 6 && bytes.Equal(b[:6], prefixB) {
		for _, char := range b[:6] {
			if char < 'a' || 'z' < char {
				dtr.b = nil
				dtr.skipCheck = true
				stdio.SetDataType(types.Generic)
				return
			}
		}
	}

	dtr.b = nil
	dtr.skipCheck = true
	//return types.Generic
	stdio.SetDataType(types.Generic)
}
