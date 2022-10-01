package models

//定义请求时的结构体

// PramsSignUp 注册时参数
type PramsSignUp struct {
	Email      string `json:"email" binding:"required,email"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// PramsLogin 登陆时参数
type PramsLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
