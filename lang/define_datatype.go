package lang

import (
	"github.com/lmorg/murex/lang/types"
)

type DataTypeLayoutT int

const (
	DataTypeIsUnknown  DataTypeLayoutT = iota
	DataTypeIsTable                    // eg csv
	DataTypeIsMarkup                   // eg xml
	DataTypeIsKeyValue                 // eg toml
	DataTypeIsObject                   // eg json
	DataTypeIsList                     // eg str
	DataTypeIsNumeric                  // eg float, int, etc
	DataTypeIsBoolean                  // eg bool
)

var (
	_dataTypeLayout = make(map[string]DataTypeLayoutT)
)

func init() {

	_dataTypeLayout[types.String] = DataTypeIsList

}

func RegisterDataType(dt string, layout DataTypeLayoutT) {
	_dataTypeLayout[dt] = layout
}

func GetDataTypeLayout(dt string) DataTypeLayoutT {
	return _dataTypeLayout[dt]
}
