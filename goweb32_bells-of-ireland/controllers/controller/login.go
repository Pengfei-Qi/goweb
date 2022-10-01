package controller

import (
	"goweb32_bells-of-ireland/logic"
	"goweb32_bells-of-ireland/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	//获取ID 及 密码
	p := new(models.PramsLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login prams is invalid ,", zap.Error(err))

		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//校验用户信息
	if ok, errs := logic.CheckLoginUserInfo(p); !ok {
		zap.L().Error("CheckLoginUserInfo is failed , err is :" + errs.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Error(),
		})
		return
	}

	//成功响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "login success",
	})

}
