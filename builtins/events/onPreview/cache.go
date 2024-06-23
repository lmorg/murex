package onpreview

import "github.com/lmorg/murex/utils/cache"

func cacheNamespace(key string) string {
	return cache.PREVIEW_EVENT + ":" + key
}

func cacheHashGet(key string, item string, cmdLine []rune, block []rune) string {
	var v bool

	ok := cache.Read(cache.PREVIEW_EVENT, key, &v)

	if ok && v { // cacheCmdLine == true
		return cache.CreateHash(string(cmdLine), block)
	}

	return cache.CreateHash(item, block)
}

func cacheHashSet(key string, cacheCmdLine bool) {
	cache.Write(cache.PREVIEW_EVENT, key, cacheCmdLine, cache.Days(365))
}
