package log_stash

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwyPayLoad struct {
	NickName string `json:"nick_name"`
	RoleID   uint   `json:"role"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"username"`
}

type CustomClaims struct {
	JwyPayLoad
	jwt.StandardClaims
}

func parseToken(token string) (jwtPayload *JwyPayLoad) {
	Token, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if Token == nil || Token.Claims == nil {
		return nil
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		return nil
	}
	return &claims.JwyPayLoad
}
