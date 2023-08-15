package common

import (
	"github.com/Earl-Power/Gin.Vue/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("Earl_Power_Gin")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func ReleaseToken(user models.User) (string, error) {
	ExpirationTime := time.Now().Add(3 * 24 * time.Hour)
	ClaimsInfo := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "learnlib.com",
			Subject:   "User Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ClaimsInfo)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	ClaimsInit := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, ClaimsInit, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, ClaimsInit, err
}
