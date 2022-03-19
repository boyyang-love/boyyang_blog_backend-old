/**
 * @Author: boyyang
 * @Date: 2022-02-18 16:49:14
 * @LastEditTime: 2022-02-18 17:00:32
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\utils\md5.go
 */
package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	r_str := fmt.Sprintf("%x", h.Sum(nil))
	return r_str
}
