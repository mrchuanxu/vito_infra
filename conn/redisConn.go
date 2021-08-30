package conn

import (
	"context"
	"time"

	"github.com/go-redis/redis/v7"
)

const (
	// DefaultMaxIdleCount 默认最大闲置连接数
	DefaultMaxIdleCount = 10
	// DefaultMaxActiveCount 默认最大活动连接数
	DefaultMaxActiveCount = 100
	// DefaultDialConnectTimeout 默认拨号超时
	DefaultDialConnectTimeout = 5
	// DefaultDialReadTimeout 默认读超时
	DefaultDialReadTimeout = 1
	// DefaultDialWriteTimeout 默认写超时
	DefaultDialWriteTimeout = 1
	// DefaultIdelTimeout 默认闲置超时
	DefaultIdelTimeout = 300
	// DefaultConnCheckTime 默认连接检查时间间隔
	DefaultConnCheckTime = 120
	// DefaultWaitFlag           = true
)

// RedisConfig redis config func
type RedisConfig func(*redisConfigs)

// RedisConfig Redis连接配置
type redisConfigs struct {
	Addr string
	// DbID     int
	Password string
}

// PrepareRedis 准备redis连接
func PrepareRedis(ctx context.Context, redisConfs ...RedisConfig) (redisClient *redis.Client, err error) {
	c := &redisConfigs{}
	for _, conf := range redisConfs {
		conf(c)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:               c.Addr,
		Password:           c.Password,
		PoolSize:           DefaultMaxActiveCount,
		MinIdleConns:       DefaultMaxIdleCount,
		IdleTimeout:        DefaultIdelTimeout * time.Second,
		ReadTimeout:        DefaultDialReadTimeout * time.Second,
		WriteTimeout:       DefaultDialWriteTimeout * time.Second,
		DialTimeout:        DefaultDialConnectTimeout * time.Second,
		IdleCheckFrequency: DefaultConnCheckTime * time.Second,
	})
	return redisClient, nil
}

// RedisAddr redis address
func RedisAddr(addr string) RedisConfig {
	return func(arg *redisConfigs) {
		arg.Addr = addr
	}
}

// RedisPassword redis pass
func RedisPassword(pass string) RedisConfig {
	return func(arg *redisConfigs) {
		arg.Password = pass
	}
}
