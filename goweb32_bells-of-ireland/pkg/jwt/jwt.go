package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	mySigningKey        = []byte("AllYourBase")
	TokenExpireDuration = time.Minute * 90
)

type MyCustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GetToken 参考链接: https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible#example-NewWithClaims-CustomClaimsType
func GetToken(userID int64, email string) (string, error) {
	claims := &MyCustomClaims{
		userID,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "my_loves",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	//zap.L().Info(fmt.Sprintf("generate key is %v %v", ss, err))
	return ss, err
}

// ParseToken token解析 参考链接: https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible#example-NewWithClaims-CustomClaimsType
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//fmt.Printf("userID: %v \t email: %v \t ExpiresAt:%v", claims.UserID, claims.Email, claims.StandardClaims.ExpiresAt)
		return claims, nil
	}
	return nil, errors.New("token is invalid")
}
