package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"gin-basic/internal/cfg"
	"gin-basic/internal/ports/icache"

	"github.com/redis/go-redis/v9"
)

const (
	defaultPoolSize     = 10
	defaultMinIdleConns = 2
	defaultDialTimeout  = 5 * time.Second
)

type redisClient struct {
	client *redis.Client
}

func NewRedis(c cfg.Redis) (icache.IRedis, error) {
	if c.Host == "" {
		return nil, nil
	}

	port, err := strconv.Atoi(c.Port)
	if err != nil {
		port = 6379
	}

	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", c.Host, port),
		Password:     c.Password,
		PoolSize:     defaultPoolSize,
		MinIdleConns: defaultMinIdleConns,
		DialTimeout:  defaultDialTimeout,
	})

	ctx, cancel := context.WithTimeout(context.Background(), defaultDialTimeout)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	return &redisClient{client: client}, nil
}

func (r *redisClient) Client() *redis.Client {
	return r.client
}

func (r *redisClient) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *redisClient) Close() error {
	return r.client.Close()
}
