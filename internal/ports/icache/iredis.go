package icache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	Client() *redis.Client
	Ping(ctx context.Context) error
	Close() error
}
