package logic

import (
	"goweb32_bells-of-ireland/dao/mysql"
	"goweb32_bells-of-ireland/models"
	"goweb32_bells-of-ireland/pkg/snowflake"
)

func SignUp(up *models.PramsSignUp) {
	//1. 判断用户是否存在
	mysql.QueryUserInfoByUsername()
	//2. 生成UUID
	snowflake.GenID()
	//3. 保存数据
	mysql.InsertUserInfo()
}
