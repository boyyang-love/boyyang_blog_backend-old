/**
 * @Author: boyyang
 * @Date: 2022-02-18 14:33:00
 * @LastEditTime: 2022-06-09 16:43:01
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\utils\generateToken.go
 */

package utils

import (
	"blog/setting"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int    `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(username string, password string, id int) (string, error) {
	// nowTime := time.Now()
	// expireTime := nowTime.Add(5 * time.Hour)
	claims := Claims{
		Username: username,
		Password: password,
		Id:       id,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: expireTime.Unix(),
			Issuer: "YZH",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(setting.JwtSecret)
	return token, err
}
