/**
 * @Author: boyyang
 * @Date: 2022-02-18 16:49:14
 * @LastEditTime: 2022-06-23 18:41:35
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\utils\md5.go
 */
package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

// MD5加密
func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	r_str := fmt.Sprintf("%x", h.Sum(nil))
	return r_str
}
