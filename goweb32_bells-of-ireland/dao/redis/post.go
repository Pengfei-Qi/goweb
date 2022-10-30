package redis

import (
	"fmt"
	"goweb32_bells-of-ireland/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func GetPostIDsByPage(params *models.ParamPostData) ([]string, error) {

	//排序参数
	key := GetRedisKey(KeyPostTimeZSet)
	if params.Order == models.OrderByScore {
		key = GetRedisKey(KeyPostScoreZSet)
	}
	return pageData(params.Page, params.Size, key)
}

func pageData(page, size int64, key string) ([]string, error) {
	start := (page - 1) * size
	stop := start + size - 1

	//获取分页参数
	return client.ZRevRange(key, start, stop).Result()
}

func GetPostsVoteNum(ids []string) (count []int64, err error) {

	pipeline := client.Pipeline()
	//根据ids 获取投票数据
	for _, id := range ids {
		key := GetRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	exec, err := pipeline.Exec()
	if err != nil {
		return
	}
	count = make([]int64, 0, len(exec))
	for _, c := range exec {
		//没看懂是怎么操作的
		cmd := c.(*redis.IntCmd)
		val := cmd.Val()
		count = append(count, val)
	}
	return
}

// GetCommunityPostIDsByOrder 按照社区分类及排序方式, 查询帖子id
/** 参考命令理解执行过程
> sadd cummunity "a1"
1
> sadd cummunity "a2"
1
> zadd time 20 "a1" 30 "a2"
2
> zinterstore mypoint 2  cummunity time AGGREGATE MAX #用来获取结果的
2
> zrange mypoint 0 -1
a1
a2
*/
func GetCommunityPostIDsByOrder(param *models.ParamPostData) ([]string, error) {

	//获取 orderKey
	orderKey := GetRedisKey(KeyPostTimeZSet)
	if param.Order == models.OrderByScore {
		orderKey = GetRedisKey(KeyPostScoreZSet)
	}
	//获取communityKey
	cKey := GetRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(param.CommunityID)))

	//新生成的key    例如: post:time:abc123
	key := fmt.Sprintf("%s:%s", orderKey, strconv.Itoa(int(param.CommunityID)))
	if count := client.Exists(key).Val(); count < 1 {
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, orderKey, cKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return pageData(param.Page, param.Size, key)
}
