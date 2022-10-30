package controller

import (
	"goweb32_bells-of-ireland/dao/redis"
	"goweb32_bells-of-ireland/logic"
	"goweb32_bells-of-ireland/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostVoteController(c *gin.Context) {
	//1. 校验参数
	pram := new(models.PramsVoteData)

	if err := c.ShouldBindJSON(pram); err != nil {
		zap.L().Error("PostVoteController prams failed, error is ", zap.Error(err))
		//判断是不是validator的异常
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidPram)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidPram, removeTopStruct(errs.Translate(trans)))
		return
	}
	//获取当前用户
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}

	//帖子投票记录
	err = logic.VoteForPost(userID, pram)
	if err != nil {
		zap.L().Error("logic.VoteForPost failed, error is", zap.Error(err))
		if err == redis.ErrorsPostRepeated {
			ResponseErrorWithMsg(c, CodeSeverBusy, err.Error())
			return
		}
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, nil)
}
