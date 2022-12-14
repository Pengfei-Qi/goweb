package mysql

import (
	"goweb32_bells-of-ireland/models"

	"golang.org/x/crypto/bcrypt"
)

// CheckUserExist 判断用户是否存在
func CheckUserExist(email string) (err error) {
	sqlStr := "select count(1) from user where email = ? "
	var num int
	if err = db.Get(&num, sqlStr, email); err != nil {
		return err
	}
	if num > 0 {
		return ErrorAccountExit
	}
	return
}

func InsertUserInfo(user *models.User) (err error) {
	insertStr := "insert into user (user_id,username,password,email) values(?,?,?,?)"
	user.Password, _ = encryptPassword(user.Password)
	if _, err = db.Exec(insertStr, user.UserID, user.Username, user.Password, user.Email); err != nil {
		return err
	}
	return
}

// encryptPassword 密码加密
func encryptPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHashAndPwd 密码校验
func CompareHashAndPwd(pwd1, pwd2 string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2)); err != nil {
		return false
	} else {
		return true
	}
}

func GetUserByEmail(user *models.User) error {
	queryStr := "select user_id,username,password,email from user where email = ?"
	err := db.Get(user, queryStr, user.Email)
	return err
}

func GetUserByID(id int64) (user *models.User, err error) {
	user = new(models.User)
	queryStr := "select user_id,username,password,email from user where user_id = ?"
	err = db.Get(user, queryStr, id)
	return
}
