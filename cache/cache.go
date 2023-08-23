package cache

import (
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	redisCache "github.com/gregjones/httpcache/redis"

	"github.com/gomodule/redigo/redis"
)

// Create a new disk-based cache
func NewDiskCache(basePath string) httpcache.Cache {
	return diskcache.New(basePath)
}

// Create a new memory-based cache
func NewMemoryCache() httpcache.Cache {
	return httpcache.NewMemoryCache()
}

// Create a new redis-based cache
func NewRedisCache(conn redis.Conn) httpcache.Cache {
	return redisCache.NewWithClient(conn)
}
