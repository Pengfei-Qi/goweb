package redis

import (
	"fmt"
	"goweb30_app_framework2/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(redisConf *settings.RedisConfig) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password, // no password set
		DB:       redisConf.DB,       // use default DB
		PoolSize: redisConf.PoolSize,
	})

	_, err = client.Ping().Result()
	if err != nil {
		fmt.Printf("redis connect failed")
		return err
	}
	return
}
func Close() {
	_ = rdb.Close()
}
