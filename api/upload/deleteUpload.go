/**
 * @Author: boyyang
 * @Date: 2022-06-09 17:01:22
 * @LastEditTime: 2022-06-10 15:39:51
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\api\upload\deleteUpload.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeleteUpload(c *gin.Context) {
	id := c.Query("id")
	var upload models.Upload
	if strings.Trim(id, "") == "" {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "id不能为空"}),
		)
	} else {
		global.
			DB.
			Where("id = ?", id).
			First(&upload).
			Delete(&upload)
		path := upload.Url
		_, err := global.Client.Object.Delete(context.Background(), path)
		if err == nil {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "删除成功"}),
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "删除失败", Error: err}),
			)
		}
	}
}
