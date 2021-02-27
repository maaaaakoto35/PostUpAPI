package controllers

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtCustomClaims this struct is for jwt.
type JwtCustomClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// jwtUserID this middleware is taking user_id and user_name.
func jwtUserID(c Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims.UserID
}

// jwtUserID this middleware is taking user_id and user_name.
func jwtUserName(c Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims.UserName
}

// setJwt this func is setting token.
func setJwt(userID string, userName string) (t string, err error) {
	// 秘密鍵を読み込み
	keyPath := os.Getenv("SECRET_KEY_PATH")
	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return
	}

	claims := &JwtCustomClaims{
		userID,
		userName,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, err = token.SignedString(key)
	return
}
