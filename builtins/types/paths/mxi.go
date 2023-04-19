package paths

import (
	"fmt"
	gopath "path"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/path"
)

func init() {
	lang.MxInterfaces[types.Path] = new(mxiPath)
}

type mxiPath struct {
	path string
}

func (mxi mxiPath) New(path string) (lang.MxInterface, error) {
	return &mxiPath{path}, nil
}

func (mxi *mxiPath) GetValue() interface{} {
	v, _ := path.Unmarshal([]byte(mxi.path))
	return v
}

func (mxi *mxiPath) GetString() string {
	return gopath.Clean(mxi.path)
}

func (mxi *mxiPath) Set(v interface{}, changePath []string) error {
	switch t := v.(type) {
	case string:
		mxi.path = t
	case []byte:
		mxi.path = string(t)
	default:
		if len(changePath) == 2 &&
			(changePath[1] == path.EXISTS || changePath[1] == path.IS_DIR || changePath[1] == path.IS_RELATIVE) {
			return fmt.Errorf("'%s' is a read only property", changePath[1])
		}

		b, err := path.Marshal(v)
		if err != nil {
			return err
		}
		mxi.path = string(b)
	}
	return nil
}
