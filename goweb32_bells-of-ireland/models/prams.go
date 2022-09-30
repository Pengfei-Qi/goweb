package models

type PramsSignUp struct {
	Age        uint8  `json:"age" binding:"required,gte=1,lte=120"`
	Email      string `json:"email" binding:"required,email"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
