package jwt

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

/**
参考链接: https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html
*/

var (
	mySigningKey               = []byte("AllYourBase")
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

type MyCustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GetToken 参考链接: https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible#example-NewWithClaims-CustomClaimsType
func GetToken(userID int64, email string) (aToken, rToken string, err error) {
	claims := &MyCustomClaims{
		userID,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("token_expired")) * time.Minute).Unix(),
			Issuer:    "my_loves",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	aToken, err = token.SignedString(mySigningKey)
	//zap.L().Info(fmt.Sprintf("generate key is %v %v", ss, err))
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), //有效期
		Issuer:    "my_loves",                                        //签发人
	})
	rToken, err = newToken.SignedString(mySigningKey)

	return aToken, rToken, err
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

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	//解析rToken
	_, err = jwt.ParseWithClaims(rToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return
	}
	// 从旧access token中解析出claims数据
	var myClaims = &MyCustomClaims{}
	_, err = jwt.ParseWithClaims(aToken, myClaims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	v, _ := err.(*jwt.ValidationError)
	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GetToken(myClaims.UserID, myClaims.Email)
	}
	return
}
