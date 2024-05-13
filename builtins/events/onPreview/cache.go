package onpreview

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
)

var cacheTTL int

func init() {
	config.InitConf.Define("cache", "default-onPreview-TTL", config.Properties{
		Description: "The delimiter for records in a CSV file.",
		Default:     60 * 60 * 24 * 30, // 30 days
		DataType:    types.Integer,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  cacheTtlRead,
			Write: cacheTtlWrite,
		},
	})
}

func cacheTtlRead() (any, error) {
	return cacheTTL, nil
}

func cacheTtlWrite(v any) error {
	i, err := types.ConvertGoType(v, types.Integer)
	if err != nil {
		return err
	}

	cacheTTL = i.(int)
	return nil
}
