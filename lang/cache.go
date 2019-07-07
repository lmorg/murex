package lang

import (
	"sync"
	"time"
)

// AstCache is an exported AST cache
var AstCache = newAstCache()

type cacheItem struct {
	created time.Time
	nodes   astNodes
	pErr    ParserError
}

type astCache struct {
	cache sync.Map
}

func newAstCache() *astCache {
	c := new(astCache)
	go astGarbageCollector(c)
	return c
}

func astGarbageCollector(c *astCache) {
	for {
		time.Sleep(5 * time.Minute)
		c.cache.Range(func(k interface{}, v interface{}) bool {
			if v.(*cacheItem).created.Add(5 * time.Minute).Before(time.Now()) {
				c.cache.Delete(k)
			}
			return true
		})
	}
}

// ParseCache scans through the AST cache looking for a match. If none found
// then it runs the parser itself.
func (c *astCache) ParseCache(block []rune) (astNodes, ParserError) {
	key := string(block)
	v, ok := c.cache.Load(key)
	if ok {
		cache := v.(*cacheItem)
		return cache.nodes, cache.pErr
	}

	nodes, pErr := parser(block)

	cache := &cacheItem{
		created: time.Now(),
		nodes:   nodes,
		pErr:    pErr,
	}
	c.cache.Store(key, cache)
	return nodes, pErr
}

// Dump returns the items in AST cache
func (c *astCache) Dump() (dump []string) {
	c.cache.Range(func(k interface{}, v interface{}) bool {
		dump = append(dump, v.(*cacheItem).created.String())
		return true
	})
	return dump
}
