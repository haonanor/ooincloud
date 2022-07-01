package initialize

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

func BigCache(ex time.Duration) *bigcache.BigCache {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(ex * time.Minute))
	return cache
}
