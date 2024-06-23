package cmdcache

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/cache"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("murex-cache", cmdCache, types.Json)
}

func cmdCache(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	mode, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	switch mode {
	case "trim":
		return cacheTrim(p)
	case "clear":
		return cacheClear(p)
	case "get":
		return cacheGet(p)
	case "namespaces":
		return cacheNamespaces(p)
	case "create-key":
		return cacheCreateKey(p)
	case "db-path":
		return cacheDbPath(p)
	case "db-enabled":
		return cacheDbEnabled(p)
	}

	return nil
}

func cacheGet(p *lang.Process) error {
	namespace, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	key, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	var v any
	_ = cache.Read(namespace, key, &v)
	return _toJson(p, v)
}

func cacheNamespaces(p *lang.Process) error {
	v := cache.ListNamespaces()
	return _toJson(p, v)
}

func cacheCreateKey(p *lang.Process) error {
	cmdLine, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	block, err := p.Parameters.String(2)
	if err != nil {
		return err
	}
	s := cache.CreateHash(cmdLine, []rune(block))
	return _toJson(p, s)
}

func cacheTrim(p *lang.Process) error {
	v, err := cache.Trim(p.Context)
	if err != nil {
		return err
	}

	return _toJson(p, v)
}

func cacheClear(p *lang.Process) error {
	v, err := cache.Clear(p.Context)
	if err != nil {
		return err
	}

	return _toJson(p, v)
}

func cacheDbPath(p *lang.Process) error {
	s := cache.DbPath()
	return _toJson(p, s)
}

func cacheDbEnabled(p *lang.Process) error {
	v := cache.DbEnabled()
	return _toJson(p, v)
}

func _toJson(p *lang.Process, v any) error {
	b, err := json.Marshal(v, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
