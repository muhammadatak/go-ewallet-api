package connection

import (
	"context"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/muhammadatak/go-e-wallet/domain"
)

func GetCacheConnection() domain.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("error connecting to cache %s", err.Error())
	}
	return cache
}
