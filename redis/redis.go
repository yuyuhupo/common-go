package redis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

const (
	// Timeout is the default timeout for a Redis client.
	Timeout = 5 * time.Second
)

type redis struct {
	cmd goredis.Cmdable
}

// New creates a new Redis client.
func New(config Config) IRedis {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	rdb := goredis.NewClient(&goredis.Options{
		Network:  config.Network,
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %s, error: %v", pong, err)
		return nil
	}

	return &redis{cmd: rdb}
}

func (r *redis) IsConnected() bool {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	if r.cmd == nil {
		return false
	}
	_, err := r.cmd.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to ping Redis: %v", err)
		return false
	}

	return true
}

func (r *redis) Get(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	strValue, err := r.cmd.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(strValue), value)
	if err != nil {
		return err
	}
	return nil
}

func (r *redis) Set(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	strValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = r.cmd.Set(ctx, key, strValue, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redis) SetWithExpiration(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	strValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = r.cmd.Set(ctx, key, strValue, expiration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redis) Remove(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	_, err := r.cmd.Del(ctx, keys...).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redis) Keys(pattern string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	keys, err := r.cmd.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (r *redis) RemovePattern(pattern string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	keys, err := r.cmd.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}

	_, err = r.cmd.Del(ctx, keys...).Result()
	if err != nil {
		return err
	}
	return nil
}
