package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")
var CtxUserId = "userId"

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	//1. 获取ID
	id, ok := c.Get(CtxUserId)
	if !ok {
		err = ErrorUserNotLogin
		return

	}
	//2. 获取用户信息
	userID, ok = id.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
