package middleware

import (
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"time"
)

var mySecret = []byte("atian-2019")

type CustomClaims struct {
	UserID   int
	Username string
	jwt2.StandardClaims
}

var JWT = jwt.New(jwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	},
	Expiration:    true,
	SigningMethod: jwt.SigningMethodHS256,
	ErrorHandler: func(context context.Context, err error) {
		if err != nil {
			//context.StatusCode(http.StatusUnauthorized)
			_, _ = context.JSON(iris.Map{
				"success": false,
				"err_msg": "token 验证失败:" + err.Error(),
				"data":    []int{},
			})
			return
		}
		context.Next()
	},
}).Serve

func GenerateToken(userId int, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(24) * time.Hour)
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt2.MapClaims{
		"userId":   userId,
		"username": username,
		"exp":      expireTime.Unix(),
		"iss":      "iris",
	})
	//token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, CustomClaims{
	//	userId,
	//	username,
	//	jwt2.StandardClaims{
	//		ExpiresAt: expireTime.Unix(),
	//		Issuer:    "iris",
	//	},
	//})
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(mySecret)
}
