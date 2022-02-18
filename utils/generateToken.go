/*
 * @Author: boyyang
 * @Date: 2022-02-18 14:33:00
 * @LastEditTime: 2022-02-18 16:44:30
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\utils\generateToken.go
 */

package utils

import (
	"fmt"
	"websit/setting"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int    `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(username string, password string, id int) (string, error) {
	fmt.Println(setting.JwtSecret, username, password, id)
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
	fmt.Println("tokenClaims", tokenClaims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(setting.JwtSecret)

	fmt.Println(token)

	return token, err
}
