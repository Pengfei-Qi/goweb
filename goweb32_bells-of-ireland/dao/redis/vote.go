package redis

import (
	"errors"
	"goweb32_bells-of-ireland/models"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

/**
实现帖子投票:
direction=1
  a. 之前未投过票, 现在投赞成票  +1
  b. 之前投反对票, 现在投赞成票  +2
direction=0
  a. 之前投反对票, 现在取消投票  +1
  b. 之前投赞成票, 现在取消投票  -1
direction=-1
  a. 之前未投过票, 现在投反对票  -1
  b. 之前投赞成票, 现在投反对票  -2
*/

var (
	ErrorsTimeExpired  = errors.New("帖子已过期")
	ErrorsPostRepeated = errors.New("帖子不能重新投票")
)

const (
	PostExpireTime = 7 * 24 * 3600
	PostScore      = 432
)

func VoteForPost(userID, postID string, direction float64) error {
	//获取投票时间
	postTime, _ := client.ZScore(GetRedisKey(KeyPostTimeZSet), postID).Result()
	//检查帖子过期时间
	if float64(time.Now().Unix()-PostExpireTime) > postTime {
		return ErrorsTimeExpired
	}
	//获取帖子分数
	pv, _ := client.ZScore(GetRedisKey(KeyPostVotedZSetPrefix+postID), userID).Result()
	//帖子不能重复投票
	if pv == direction {
		return ErrorsPostRepeated
	}
	// 判断用户行为
	var pt float64
	if direction > pv {
		pt = 1
	} else {
		pt = -1
	}
	//创建redis事务
	txPipeline := client.TxPipeline()
	//创建帖子分数
	score := math.Abs(direction - pv)
	txPipeline.ZIncrBy(GetRedisKey(KeyPostScoreZSet), pt*score*PostScore, postID)
	//创建帖子记录
	if direction == 0 {
		txPipeline.ZRem(GetRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		txPipeline.ZAdd(GetRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  direction,
			Member: userID,
		})
	}
	_, err := txPipeline.Exec()
	return err
}

func CreatePostTime(p *models.Post) error {
	//创建redis事务
	txPipeline := client.TxPipeline()
	postId := strconv.FormatInt(p.ID, 10)
	//帖子时间
	txPipeline.ZAdd(GetRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	//帖子分数
	txPipeline.ZAdd(GetRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  PostScore,
		Member: postId,
	})
	//添加社区和帖子关联关系
	txPipeline.SAdd(GetRedisKey(KeyCommunitySetPrefix + strconv.FormatInt(p.CommunityId, 10)))
	_, err := txPipeline.Exec()
	return err
}
