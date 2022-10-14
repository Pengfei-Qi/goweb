package controller

import (
	"goweb32_bells-of-ireland/logic"

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
