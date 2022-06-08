/**
 * @Author: boyyang
 * @Date: 2022-02-18 14:40:33
 * @LastEditTime: 2022-06-08 18:25:48
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\utils\parseToken.go
 */

package utils

import (
	"blog/setting"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return setting.JwtSecret, nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok {
			if tokenClaims.Valid {
				return claims, nil
			}
		}
	}
	return nil, err
}
