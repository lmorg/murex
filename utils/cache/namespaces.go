package cache

const (
	PREVIEW_COMMAND      = "preview_command"
	PREVIEW_DYNAMIC      = "preview_dynamic"
	PREVIEW_EVENT        = "preview_event"
	MAN_SUMMARY          = "man_summary"
	MAN_PATHS            = "man_paths"
	MAN_FLAGS            = "man_flags"
	AUTOCOMPLETE_DYNAMIC = "autocomplete_dynamic"
	HINT_SUMMARY         = "hint_summary"
)

func InitCache() {
	initCache(PREVIEW_COMMAND)
	initCache(PREVIEW_DYNAMIC)
	initCache(PREVIEW_EVENT)
	initCache(MAN_SUMMARY)
	initCache(MAN_PATHS)
	initCache(MAN_FLAGS)
	initCache(AUTOCOMPLETE_DYNAMIC)
	initCache(HINT_SUMMARY)
}
