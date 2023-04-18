package pkg

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

var (
	redisClient         redis.UniversalClient
	initRedisClientOnce sync.Once
)

// GetRedisClient TODO
func GetRedisClient() redis.UniversalClient {
	initRedisClientOnce.Do(initRedisClient)
	return redisClient
}

func initRedisClient() {
	host := "localhost:6379"
	if value := viper.GetString("redis.host"); value != "" {
		host = value
	}
	address := strings.Split(host, ",")
	var opts = &redis.UniversalOptions{
		PoolSize:     8,
		MinIdleConns: 10,
		MaxConnAge:   60 * time.Second,
		PoolTimeout:  10 * time.Second,
		IdleTimeout:  3 * time.Second,
	}
	if value := viper.GetInt("redis.pool_size"); value > 0 {
		opts.PoolSize = value
	}
	if value := viper.GetInt("redis.min_idle_conns"); value > 0 {
		opts.MinIdleConns = value
	}
	if value := viper.GetInt("redis.max_conn_age"); value > 0 {
		opts.MaxConnAge = time.Duration(value) * time.Second
	}
	if value := viper.GetInt("redis.pool_timeout"); value > 0 {
		opts.PoolTimeout = time.Duration(value) * time.Second
	}
	if value := viper.GetInt("redis.idle_timeout"); value > 0 {
		opts.IdleTimeout = time.Duration(value) * time.Second
	}
	if value := viper.GetInt("redis.db"); value > 0 {
		opts.DB = value
	}
	zap.S().Infof("redis client config %+v", opts)
	if viper.GetBool("redis.cluster") {
		redisClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        address,
			PoolSize:     opts.PoolSize,
			MinIdleConns: opts.MinIdleConns,
			MaxConnAge:   opts.MaxConnAge,
			PoolTimeout:  opts.PoolTimeout,
			IdleTimeout:  opts.IdleTimeout,
		})
	} else {
		redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:        address,
			DB:           opts.DB,
			PoolSize:     opts.PoolSize,
			MinIdleConns: opts.MinIdleConns,
			MaxConnAge:   opts.MaxConnAge,
			PoolTimeout:  opts.PoolTimeout,
			IdleTimeout:  opts.IdleTimeout,
		})
	}
}
