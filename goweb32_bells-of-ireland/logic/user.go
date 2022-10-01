package logic

import (
	"errors"
	"goweb32_bells-of-ireland/dao/mysql"
	"goweb32_bells-of-ireland/models"
	"goweb32_bells-of-ireland/pkg/snowflake"
)

func SignUp(up *models.PramsSignUp) (err error) {
	//1. 判断用户是否存在
	if err = mysql.CheckUserExist(up.Email); err != nil {
		return err
	}
	//2. 生成UUID
	userId := snowflake.GenID()

	//3. 保存数据
	var user = &models.User{
		UserID:   userId,
		Email:    up.Email,
		Username: up.Username,
		Password: up.Password,
	}
	//密码加密
	return mysql.InsertUserInfo(user)
}

// CheckLoginUserInfo 用户登陆
func CheckLoginUserInfo(loginInfo *models.PramsLogin) (flag bool, err error) {
	var u1 models.User
	//获取用户信息
	if err, u1 = mysql.GetUserByEmail(loginInfo.Email); u1.IsEmpty() {
		return false, errors.New("找不到该用户")
	}
	//校验密码
	if !mysql.CompareHashAndPwd(u1.Password, loginInfo.Password) {
		return false, errors.New("密码不一致,请重新输入")
	}
	return true, nil
}
