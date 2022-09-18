package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.235.233:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	return
}

// doCommand go-redis基本使用示例
func doCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	//赋值
	result, err := rdb.Set(ctx, "company", "上海", time.Hour).Result()
	if err != nil {
		fmt.Printf("redis exec result , %v\n", err)
		return
	}
	fmt.Printf("redis exec set value result is %v \n", result)

	//取值
	v1 := rdb.Get(ctx, "company").Val()
	fmt.Printf("[1]redis exec get value result is %v \n", v1)

	v2, err := rdb.Get(ctx, "company").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case v2 == "":
		fmt.Println("value is empty")
	default:
		fmt.Printf("[2]redis exec get value result is %v \n", v2)
	}
}

// doDemo rdb.Do 方法使用示例
func doDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 直接执行命令获取错误
	err := rdb.Do(ctx, "set", "treeCount", 10, "EX", 3600).Err()
	fmt.Println(err)

	// 执行命令获取结果
	val, err := rdb.Do(ctx, "get", "treeCount").Result()
	fmt.Println(val, err)
}

// zsetDemo 操作zset示例
func zsetDemo() {
	// key
	zsetKey := "language_rank"
	// value
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	if err := initRedis(); err != nil {
		fmt.Printf("connect redis db failed , error info :%v \n", err)
		panic(err)
	}
	defer rdb.Close()
	fmt.Printf("connect redis db success \n")

	//doCommand()
	//doDemo()

	zsetDemo()
}
