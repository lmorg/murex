package columns

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.RegisterMarshaller(types.Columns, marshal)
}
