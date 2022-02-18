/*
 * @Author: boyyang
 * @Date: 2022-02-16 15:14:21
 * @LastEditTime: 2022-02-16 15:26:32
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\utils\returnData.go
 */

package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnData(code int, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}
