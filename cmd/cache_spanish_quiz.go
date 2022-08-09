package cmd

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

var (
	defaultTime = (30 * time.Minute)
	cache       = ttlcache.New(
		ttlcache.WithTTL[string, int](defaultTime),
	)
)

func CacheSession(finished chan bool, duration time.Duration) {
	time.Sleep(duration)
	finished <- true
}
