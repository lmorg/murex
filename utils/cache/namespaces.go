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
	initNamespace(PREVIEW_COMMAND)
	initNamespace(PREVIEW_DYNAMIC)
	//initCache(PREVIEW_EVENT)
	initNamespace(MAN_SUMMARY)
	initNamespace(MAN_PATHS)
	initNamespace(MAN_FLAGS)
	initNamespace(AUTOCOMPLETE_DYNAMIC)
	initNamespace(HINT_SUMMARY)
}

func initNamespace(namespace string) {
	if configCacheDisabled {
		return
	}

	cache[namespace] = new(internalCacheT)
	cache[namespace].cache = make(map[string]*cacheItemT)
	createDb(namespace)
	disabled = false
}
