package lang

import (
	"sync"
	"time"
)

// AstCache is an exported AST cache
var AstCache *astCache = newAstCache()

type cacheItem struct {
	lastUsed time.Time
	nodes    astNodes
	pErr     ParserError
}

type astCache struct {
	mutex sync.Mutex
	cache map[string]*cacheItem
}

func newAstCache() *astCache {
	c := new(astCache)
	c.cache = make(map[string]*cacheItem)
	go astGarbageCollector(c)
	return c
}

func astGarbageCollector(c *astCache) {
	for {
		time.Sleep(60 * time.Second)
		c.mutex.Lock()
		for key := range c.cache {
			if c.cache[key].lastUsed.Add(60 * time.Second).Before(time.Now()) {
				delete(c.cache, key)
			}
		}
		c.mutex.Unlock()
	}
}

// ParseCache scans through the AST cache looking for a match. If none found
// then it runs the parser itself.
func (c *astCache) ParseCache(block []rune) (astNodes, ParserError) {
	key := string(block)

	c.mutex.Lock()
	cache := c.cache[key]
	c.mutex.Unlock()

	if cache != nil {
		c.mutex.Lock()
		cache.lastUsed = time.Now()
		c.mutex.Unlock()

		return cache.nodes, cache.pErr
	}

	nodes, pErr := parser(block)
	//if pErr.Code > 0 {
	//	return nil, pErr
	//}

	c.mutex.Lock()
	c.cache[key] = &cacheItem{
		lastUsed: time.Now(),
		nodes:    nodes,
		pErr:     pErr,
	}
	c.mutex.Unlock()
	return nodes, pErr
}

// Dump returns the items in AST cache
func (c *astCache) Dump() (dump []string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key := range c.cache {
		dump = append(dump, c.cache[key].lastUsed.String())
	}
	return dump
}
