package cache

import "github.com/lmorg/murex/utils/cache/cachelib"

const (
	PREVIEW_DYNAMIC      = "preview_dynamic"
	MAN_SUMMARY          = "man_summary"
	AUTOCOMPLETE_DYNAMIC = "autocomplete_dynamic"
)

func init() {
	cachelib.CreateTable(PREVIEW_DYNAMIC)
	cachelib.CreateTable(MAN_SUMMARY)
	cachelib.CreateTable(AUTOCOMPLETE_DYNAMIC)
}
