package cache

import (
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/trace"
	"github.com/xinliangnote/go-gin-api/pkg/time_parse"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

type Option func(*option)

type Trace = trace.T

type option struct {
	Trace *trace.Trace
	Redis *trace.Redis
}

func newOption() *option {
	return &option{}
}

var _ Repo = (*cacheRepo)(nil)

type Repo interface {
	i()
	Set(key, value string, ttl time.Duration, options ...Option) error
	Get(key string, options ...Option) (string, error)
	TTL(key string) (time.Duration, error)
	Expire(key string, ttl time.Duration) bool
	ExpireAt(key string, ttl time.Time) bool
	Del(keys ...string) bool
	Incr(key string, options ...Option) int64
	Close()
}

type cacheRepo struct {
	client *redis.Client
}

func New() (Repo, error) {
	client, err := redisConnect()
	if err != nil {
		return nil, err
	}

	return &cacheRepo{
		client: client,
	}, nil
}

func (c *cacheRepo) i() {}

func redisConnect() (*redis.Client, error) {
	cfg := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Pass,
		DB:           cfg.Db,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}

	return client, nil
}

// Set set some <key,value> into redis
func (c *cacheRepo) Set(key, value string, ttl time.Duration, options ...Option) error {
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = time_parse.CSTLayoutString()
			opt.Redis.Handle = "set"
			opt.Redis.Key = key
			opt.Redis.Value = value
			opt.Redis.TTL = ttl
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	if err := c.client.Set(key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}

	return nil
}

// Get get some key from redis
func (c *cacheRepo) Get(key string, options ...Option) (string, error) {
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = time_parse.CSTLayoutString()
			opt.Redis.Handle = "get"
			opt.Redis.Key = key
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	value, err := c.client.Get(key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}

	return value, nil
}

// TTL get some key from redis
func (c *cacheRepo) TTL(key string) (time.Duration, error) {
	ttl, err := c.client.TTL(key).Result()
	if err != nil {
		return -1, errors.Wrapf(err, "redis get key: %s err", key)
	}

	return ttl, nil
}

// Expire expire some key
func (c *cacheRepo) Expire(key string, ttl time.Duration) bool {
	ok, _ := c.client.Expire(key, ttl).Result()
	return ok
}

// ExpireAt expire some key at some time
func (c *cacheRepo) ExpireAt(key string, ttl time.Time) bool {
	ok, _ := c.client.ExpireAt(key, ttl).Result()
	return ok
}

// Del del some key from redis
func (c *cacheRepo) Del(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}

	value, _ := c.client.Del(keys...).Result()
	return value > 0
}

func (c *cacheRepo) Incr(key string, options ...Option) int64 {
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = time_parse.CSTLayoutString()
			opt.Redis.Handle = "incr"
			opt.Redis.Key = key
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}
	value, _ := c.client.Incr(key).Result()
	return value
}

// Close close redis client
func (c *cacheRepo) Close() {
	c.client.Close()
}

// WithTrace 设置trace信息
func WithTrace(t Trace) Option {
	return func(opt *option) {
		if t != nil {
			opt.Trace = t.(*trace.Trace)
			opt.Redis = new(trace.Redis)
		}
	}
}
