package redis

import (
	"fmt"
	"goweb32_bells-of-ireland/settings"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init(redisConf *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password, // no password set
		DB:       redisConf.DB,       // use default DB
		PoolSize: redisConf.PoolSize,
	})

	_, err = client.Ping().Result()
	if err != nil {
		zap.L().Error("redis connect failed")
		return err
	}
	return
}
func Close() {
	_ = client.Close()
}
