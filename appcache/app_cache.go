package appcache

import (
	"context"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"lesson-5-goland/common"
	"time"
)

type AppCache interface {
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Once(item *cache.Item) error
}

type appCache struct {
	store *cache.Cache
}

func NewAppCache(sc goservice.ServiceContext) *appCache {
	rdClient := sc.MustGet(common.PluginAppRedis).(*redis.Client)

	c := cache.New(&cache.Options{
		Redis:      rdClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &appCache{store: c}
}

func (ac *appCache) Get(ctx context.Context, key string, value interface{}) error {
	if err := ac.store.Get(ctx, key, &value); err != nil {
		return err
	}

	return nil
}

func (ac *appCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return ac.store.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   ttl,
	})
}

func (ac *appCache) Delete(ctx context.Context, key string) error {
	return ac.store.Delete(ctx, key)
}

func (ac *appCache) Once(item *cache.Item) error {
	err := ac.store.Once(item)

	return err
}
