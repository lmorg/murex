package cache

const (
	PREVIEW_COMMAND      = "preview_command"
	PREVIEW_DYNAMIC      = "preview_dynamic"
	MAN_SUMMARY          = "man_summary"
	MAN_PATHS            = "man_paths"
	MAN_FLAGS            = "man_flags"
	AUTOCOMPLETE_DYNAMIC = "autocomplete_dynamic"
	HINT_SUMMARY         = "hint_summary"
)

func InitCache() {
	initCache(PREVIEW_COMMAND)
	initCache(PREVIEW_DYNAMIC)
	initCache(MAN_SUMMARY)
	initCache(MAN_PATHS)
	initCache(MAN_FLAGS)
	initCache(AUTOCOMPLETE_DYNAMIC)
	initCache(HINT_SUMMARY)
}
