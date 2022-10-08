package controller

import (
	"errors"
	"goweb32_bells-of-ireland/logic"
	"goweb32_bells-of-ireland/models"
	"goweb32_bells-of-ireland/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func LoginHandler(c *gin.Context) {
	//获取ID 及 密码
	p := new(models.PramsLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login prams is invalid ,", zap.Error(err))

		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidPram)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidPram, removeTopStruct(errs.Translate(trans)))
		return
	}
	var user = &models.User{
		Email:    p.Email,
		Password: p.Password,
	}

	//校验用户信息
	if errs := logic.CheckLoginUserInfo(user); errs != nil {
		zap.L().Error("checkLoginUserInfo failed ", zap.String("email", p.Email), zap.Error(errs))
		if errors.Is(errs, logic.ErrorAccountNotExist) || errors.Is(errs, logic.ErrorInvalidPwd) {
			ResponseError(c, CodeMissAccountOrPassword)
			return
		}
		ResponseError(c, CodeSeverBusy)
		return
	}
	//生成TOKEN
	atoken, rtoken, err := jwt.GetToken(user.UserID, user.Email)
	if err != nil {
		ResponseError(c, CodeTokenException)
		return
	}
	//成功响应
	ResponseSuccess(c, gin.H{"Authorization": atoken, "refreshToken": rtoken})
}

func RefreshToken(c *gin.Context) {
	//1. 从系统头获取token
	rToken := c.Request.Header.Get("refreshToken")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseError(c, CodeNotLogin)
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseError(c, CodeInvalidToken)
		c.Abort()
		return
	}
	//2. 获取token
	//accessToken 登录token,
	//refreshToken 刷新token ,为空时需登录重新获取
	//accessToken, refreshToken, err := jwt.RefreshToken(parts[1], rToken)
	accessToken, _, err := jwt.RefreshToken(parts[1], rToken)
	if err != nil {
		ResponseError(c, CodeTokenException)
		return
	}
	//成功响应
	ResponseSuccess(c, gin.H{"Authorization": accessToken})

}
