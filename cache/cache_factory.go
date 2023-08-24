package cache

import (
	"todoapp/cache/lru"
	"todoapp/config"
)

func NewCache() Cache {
	switch config.Config.CacheConfig.Type {
	case "lru":
		return lru.NewLRU()
	default:
		return lru.NewLRU()
	}
}
