package redisclient

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// Config represents the configuration for the Redis client.
type Config struct {
	Addrs        []string
	Password     string
	DB           int
	ClusterMode  bool
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	MinIdleConns int
}

type redisClient struct {
	Client redis.UniversalClient
}

// NewRedisClient creates a new Redis client using provided configuration.
func NewRedisClient(cfg *Config) *redisClient {
	db := cfg.DB
	if cfg.ClusterMode {
		db = 0 // Redis Cluster does not support DB selection
	}

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        cfg.Addrs,
		Password:     cfg.Password,
		DB:           db,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	return &redisClient{Client: client}
}
