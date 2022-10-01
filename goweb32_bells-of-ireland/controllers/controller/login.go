package controller

import (
	"errors"
	"goweb32_bells-of-ireland/logic"
	"goweb32_bells-of-ireland/models"

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
	//校验用户信息
	if errs := logic.CheckLoginUserInfo(p); errs != nil {
		zap.L().Error("checkLoginUserInfo failed ", zap.String("email", p.Email), zap.Error(errs))
		if errors.Is(errs, logic.ErrorAccountNotExist) || errors.Is(errs, logic.ErrorInvalidPwd) {
			ResponseError(c, CodeMissAccountOrPassword)
			return
		}
		ResponseError(c, CodeSeverBusy)
		return
	}

	//成功响应
	ResponseSuccess(c, nil)

}
