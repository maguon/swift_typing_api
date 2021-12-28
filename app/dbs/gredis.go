package dbs

import (
	"context"
	"encoding/json"
	"fmt"
	"swift_typing_api/common"
	"swift_typing_api/conf"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	RedisExpiredTimes = 600
)

var ctx = context.Background()

type IRedis interface {
	IsConnected() bool
	Get(key string, data interface{}) error
	Expire(key string) error
	Set(key string, val []byte) error
	Remove(keys ...string) error
	Keys(pattern string) ([]string, error)
}

// GRedis struct
type GRedis struct {
	client     *redis.Client
	expiryTime int
}

func NewRedis() IRedis {

	redisConfig := conf.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		common.GetLogger().Error(pong, err)
		return nil
	}
	expiryTime := RedisExpiredTimes
	return &GRedis{client: rdb, expiryTime: expiryTime}
}

// IsConnected check redis is connected or not
func (g *GRedis) IsConnected() bool {
	if g.client == nil {
		return false
	}

	_, err := g.client.Ping(ctx).Result()
	if err != nil {
		common.GetLogger().Error(err)
		return false
	}
	return true
}

// Get get key from redis
func (g *GRedis) Get(key string, data interface{}) error {
	val, err := g.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil
	}
	if err != nil {

		common.GetLogger().Info("Cache fail to get: ", err)
		return nil
	}
	common.GetLogger().Debug("Get from redis %s - %s", key, val)

	err = json.Unmarshal(val, &data)
	if err != nil {
		return err
	}

	return nil
}
func (g *GRedis) Expire(key string) error {
	err := g.client.Expire(ctx, key, time.Duration(g.expiryTime)*time.Second).Err()
	if err != nil {
		common.GetLogger().Info("Cache fail to expire: ", err)
		return err
	}
	return nil
}

// Set data to redis
func (g *GRedis) Set(key string, val []byte) error {
	err := g.client.Set(ctx, key, val, time.Duration(g.expiryTime)*time.Second).Err()
	if err != nil {
		common.GetLogger().Info("Cache fail to set: ", err)
		return err
	}
	common.GetLogger().Debugf("Set to redis %s - %s", key, val)

	return nil
}

// Remove return list keys from redis
func (g *GRedis) Remove(keys ...string) error {
	err := g.client.Del(ctx, keys...).Err()
	if err != nil {
		common.GetLogger().Error("Cache fail to delete key %s: %s", keys, err)
		return err
	}
	common.GetLogger().Debug("Cache deleted key", keys)

	return nil
}

// Keys get redis keys by pattern regex
func (g *GRedis) Keys(pattern string) ([]string, error) {
	keys, err := g.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}
