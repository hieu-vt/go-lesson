package appredis

import (
	"context"
	"flag"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/go-redis/redis/v8"
)

var (
	defaultRedisName      = "DefaultRedis"
	defaultRedisMaxActive = 0 // 0 is unlimited max active connection
	defaultRedisMaxIdle   = 10
)

type RedisDBOpt struct {
	Prefix    string
	RedisUri  string
	MaxActive int
	MaxIde    int
	password  string
}

type appRedis struct {
	name   string
	client *redis.Client
	logger logger.Logger
	*RedisDBOpt
}

func NewAppRedis(name string, prefix string) *appRedis {
	return &appRedis{
		name: name,
		RedisDBOpt: &RedisDBOpt{
			Prefix: prefix,
		},
	}
}

func (ar *appRedis) Name() string {
	return ar.name
}

func (ar *appRedis) GetPrefix() string {
	return ar.Prefix
}

func (ar *appRedis) InitFlags() {
	prefix := ar.Prefix
	if ar.Prefix != "" {
		prefix += "-"
	}

	flag.StringVar(&ar.RedisUri, prefix+"-uri", "redis://localhost:6379", "(For go-redis) Redis connection-string. Ex: redis://localhost/0")
	flag.IntVar(&ar.MaxActive, prefix+"-pool-max-active", defaultRedisMaxActive, "(For go-redis) Override redis pool MaxActive")
	flag.IntVar(&ar.MaxIde, prefix+"-pool-max-idle", defaultRedisMaxIdle, "(For go-redis) Override redis pool MaxIdle")
}

func (ar *appRedis) isDisabled() bool {
	return ar.RedisUri == ""
}

func (ar *appRedis) Configure() error {
	if ar.isDisabled() {
		return nil
	}

	ar.logger = logger.GetCurrent().GetLogger(ar.name)
	ar.logger.Info("Connecting to Redis at ", ar.RedisUri, "...")

	opt, err := redis.ParseURL(ar.RedisUri)

	if err != nil {
		ar.logger.Error("Cannot parse Redis ", err.Error())
		return err
	}

	opt.PoolSize = ar.MaxActive
	opt.MinIdleConns = ar.MaxIde

	client := redis.NewClient(opt)

	// Ping to test Redis connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		ar.logger.Error("Cannot connect Redis. ", err.Error())
		return err
	}

	// Connect successfully, assign client to goRedisDB
	ar.client = client
	return nil
}

func (ar *appRedis) Run() error {
	return ar.Configure()
}

func (ar *appRedis) Get() interface{} {
	return ar.client
}

func (ar *appRedis) Stop() <-chan bool {
	if ar.client != nil {
		if err := ar.client.Close(); err != nil {
			ar.logger.Info("cannot close ", ar.name)
		}
	}

	c := make(chan bool)
	go func() { c <- true }()
	return c
}
