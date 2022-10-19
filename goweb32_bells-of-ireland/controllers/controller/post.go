package controller

import (
	"goweb32_bells-of-ireland/logic"
	"goweb32_bells-of-ireland/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreateArticleHandler(c *gin.Context) {
	// 1. 获取参数
	post := new(models.Post)
	if err := c.ShouldBindJSON(&post); err != nil {
		ResponseError(c, CodeInvalidPram)
		return
	}

	// 2. 录入数据
	// 从系统头获取当前用户信息
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Warn("getCurrentUserID failed, err is ", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorID = userID
	if err := logic.CreateArticle(post); err != nil {
		zap.L().Warn("logic.CreateArticle failed, err is ", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}

	// 3. 返回
	ResponseSuccess(c, nil)
}

func QueryArticleDetailHandler(c *gin.Context) {
	//获取参数
	postID := c.Param("id")
	id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		zap.L().Info("received data failed , err is ", zap.Error(err))
		ResponseError(c, CodeInvalidPram)
		return
	}
	//查询数据
	data, err := logic.QueryArticleDetail(id)
	if err != nil {
		zap.L().Warn("QueryArticleDetail is failed, error is ", zap.Error(err))
		ResponseError(c, CodeInvalidPram)
		return
	}
	//响应
	ResponseSuccess(c, data)

}
