package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

//生成Token
func Make(id, expire int64, iss string, signKey []byte) string {
	// Create the Claims
	claims := &jwt.MapClaims{
		"user_id": id,
		"nbf":     time.Now().Unix(),
		"exp":     time.Now().Unix() + expire,
		"iss":     iss,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString(signKey)
	if err == nil {
		return str
	}
	return ""
}

// 校验token是否有效
func Check(tokenStr string, signKey []byte) (bool, int64) {
	if tokenStr == "" {
		return false, 0
	}
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return false, 0
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return true, int64(claims["user_id"].(float64))
	}
	return false, 0
}
