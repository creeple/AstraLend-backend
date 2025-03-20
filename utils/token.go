/**
 * @Author: lixiang
 * @Date: 2025-03-19 15:44
 * @Description: token.go
 */

package utils

import (
	"AstraLend-backend/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateToken(username string) (string, error) {
	jwtConfig := config.Config.Jwt
	jwtKey := jwtConfig.SecretKey
	expirationTime := jwtConfig.ExpireTime
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime,
		"iat":      time.Now(),
		"iss":      "astra_lend",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "token create failed", err
	}
	return tokenString, nil
}
