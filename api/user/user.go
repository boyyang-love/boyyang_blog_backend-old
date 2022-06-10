/**
 * @Author: boyyang
 * @Date: 2022-02-14 17:01:53
 * @LastEditTime: 2022-06-10 15:40:10
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\api\user\user.go
 */
package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := utils.ParseToken(token)
	var form models.User
	c.Bind(&form)
	var user models.User
	data := global.
		DB.
		Where("Username = ? AND ID != ? ", form.Username, claims.Id).
		First(&user)
	// 如果用户存
	if data.RowsAffected > 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "用户名已存在"}),
		)
		return
	}
	err := global.
		DB.
		Model(models.User{}).
		Where("id = ?", claims.Id).
		Update(&form).
		Error
	if err == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "更新成功", Data: form.ID}),
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "更新失败", Data: err}),
		)
	}
}
