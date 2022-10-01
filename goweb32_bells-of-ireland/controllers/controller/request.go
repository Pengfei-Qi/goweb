package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUser(c *gin.Context) error {
	//1. 获取ID
	//id, exists := c.Get(middlewares.CtxUserId)
	//if !exists{
	//	return ErrorUserNotLogin
	//}
	//2. 获取用户信息
	//userId:= id.(int64)
	return nil
}
