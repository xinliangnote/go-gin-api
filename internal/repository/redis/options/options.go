package options

import "github.com/go-redis/redis/v7"

// RedisConfig 用于构建 Redis 连接的配置参数
type RedisConfig struct {
	Addr         string
	Addrs        []string
	MasterName   string
	Pass         string
	Db           int
	MaxRetries   int
	PoolSize     int
	MinIdleConns int
}

// BuildUniversalOptions 根据配置构建 UniversalOptions，自动适配单机/集群/哨兵模式
//
// 适配规则:
//   - 配置 Addrs（多地址）时，优先使用 Addrs → 集群模式
//   - 仅配置 Addr（单地址）时，回退为单机模式（向后兼容）
//   - 设置 MasterName 时，启用哨兵模式
func BuildUniversalOptions(cfg RedisConfig) *redis.UniversalOptions {
	addrs := cfg.Addrs
	if len(addrs) == 0 && cfg.Addr != "" {
		addrs = []string{cfg.Addr}
	}

	return &redis.UniversalOptions{
		Addrs:        addrs,
		MasterName:   cfg.MasterName,
		Password:     cfg.Pass,
		DB:           cfg.Db,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	}
}
