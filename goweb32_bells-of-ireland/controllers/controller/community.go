package controller

import (
	"goweb32_bells-of-ireland/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {

	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Warn("communityHandler get getCommunityList failed , error is ", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, data)

}

func CommunityDetailHandler(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		zap.L().Warn("id is invalid , error is ", zap.Error(err))
		ResponseError(c, CodeInvalidPram)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Warn("communityHandler get GetCommunityDetail failed , error is ", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, data)
}
