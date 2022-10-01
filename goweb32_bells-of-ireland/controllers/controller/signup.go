package controller

import (
	"errors"
	"goweb32_bells-of-ireland/dao/mysql"
	"goweb32_bells-of-ireland/logic"
	"goweb32_bells-of-ireland/models"
	"strings"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数, 并进行参数校验
	p := new(models.PramsSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("signUp prams is invalid ,", zap.Error(err))

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
	//2. 逻辑处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorAccountExit) {
			ResponseError(c, CodeAccountExist)
			return
		}
		ResponseError(c, CodeSeverBusy)
		return
	}
	//3. 数据返回
	ResponseSuccess(c, nil)
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
