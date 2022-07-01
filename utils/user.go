package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

func GenerateJwt(uid uint) (string, time.Time, error) {
	expireAt := time.Now().Add(time.Hour * 200)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expireAt.Unix(),
		Subject:   strconv.Itoa(int(uid)),
	})
	tokenString, err := token.SignedString([]byte(viper.GetStringMapString("jwt")["key"]))
	return tokenString, expireAt, err
}

func GenerateCookie(uid uint) (*http.Cookie, error) {
	tokenString, expireAt, err := GenerateJwt(uid)
	if err != nil {
		return &http.Cookie{}, err
	}
	return &http.Cookie{Name: "token", Value: tokenString, Expires: expireAt, Path: "/"}, nil
}
