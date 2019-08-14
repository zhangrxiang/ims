package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(24) * time.Hour)

	claims := Claims{
		Md5Str(username),
		Md5Str(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "iris",
			Subject:   username,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte("atian"))

	return token, err
}
