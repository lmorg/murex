package cache

import "github.com/lmorg/murex/utils/cache/cachelib"

const (
	PREVIEW_COMMAND      = "preview_command"
	PREVIEW_DYNAMIC      = "preview_dynamic"
	MAN_SUMMARY          = "man_summary"
	MAN_PATHS            = "man_paths"
	MAN_FLAGS            = "man_flags"
	AUTOCOMPLETE_DYNAMIC = "autocomplete_dynamic"
)

func init() {
	cachelib.CreateTable(PREVIEW_COMMAND)
	cachelib.CreateTable(PREVIEW_DYNAMIC)
	cachelib.CreateTable(MAN_SUMMARY)
	cachelib.CreateTable(MAN_PATHS)
	cachelib.CreateTable(MAN_FLAGS)
	cachelib.CreateTable(AUTOCOMPLETE_DYNAMIC)
}
