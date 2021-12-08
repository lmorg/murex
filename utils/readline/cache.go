package readline

import "sync"

//acSuggestions  map[string][]string
//acDescriptions map[string]map[string]string

// string

type cacheString struct {
	mutex   sync.Mutex
	indexes []string
	size    int
	values  map[string]string
}

func (c *cacheString) Init(rl *Instance) {
	c.mutex.Lock()

	c.indexes = make([]string, rl.MaxCacheSize)
	c.size = 0
	c.values = make(map[string]string)

	c.mutex.Unlock()
}

func (c *cacheString) Append(line []rune, v string) {
	sLine := string(line)
	c.mutex.Lock()

	if c.size == len(c.indexes)-1 {
		delete(c.values, c.indexes[0])
		c.indexes = append(c.indexes[1:], "")
	} else {
		c.size++
	}

	c.indexes[c.size] = sLine
	c.values[sLine] = v

	c.mutex.Unlock()
}

func (c *cacheString) Get(line []rune) string {
	sLine := string(line)
	c.mutex.Lock()

	v := c.values[sLine]

	c.mutex.Unlock()
	return v
}

// []rune

type cacheSliceRune struct {
	mutex   sync.Mutex
	indexes []string
	size    int
	values  map[string][]rune
}

func (c *cacheSliceRune) Init(rl *Instance) {
	c.mutex.Lock()

	c.indexes = make([]string, rl.MaxCacheSize)
	c.size = 0
	c.values = make(map[string][]rune)

	c.mutex.Unlock()
}

func (c *cacheSliceRune) Append(line, v []rune) {
	sLine := string(line)
	c.mutex.Lock()

	if c.size == len(c.indexes)-1 {
		delete(c.values, c.indexes[0])
		c.indexes = append(c.indexes[1:], "")
	} else {
		c.size++
	}

	c.indexes[c.size] = sLine
	c.values[sLine] = v

	c.mutex.Unlock()
}

func (c *cacheSliceRune) Get(line []rune) []rune {
	sLine := string(line)
	c.mutex.Lock()

	v := c.values[sLine]

	c.mutex.Unlock()
	return v
}

// Slice

// Map
