/*
 * @Author: boyyang
 * @Date: 2022-02-16 17:27:10
 * @LastEditTime: 2022-02-18 14:07:10
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\controller\upload.go
 */

package controller

import (
	"fmt"
	"path"
	"websit/utils"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")

	dst := path.Join("./assets", file.Filename)

	c.SaveUploadedFile(file, dst)

	fmt.Println(file.Filename)

	msg := map[string]interface{}{
		"name": file.Filename,
		"path": "/assets/" + file.Filename,
	}

	utils.ReturnData(200, msg, c)
}
