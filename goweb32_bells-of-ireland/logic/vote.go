package logic

import (
	"goweb32_bells-of-ireland/dao/redis"
	"goweb32_bells-of-ireland/models"
	"strconv"

	"go.uber.org/zap"
)

//帖子投票逻辑

func VoteForPost(userID int64, voteData *models.PramsVoteData) error {

	zap.L().Info("logic.voteData logs is ",
		zap.Int64("userId", userID),
		zap.String("postId", voteData.PostId),
		zap.Int8("direction", voteData.Direction))
	return redis.VoteForPost(strconv.FormatInt(userID, 10), voteData.PostId, float64(voteData.Direction))

}
