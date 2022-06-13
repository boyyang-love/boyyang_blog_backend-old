/**
 * @Author: boyyang
 * @Date: 2022-06-13 13:04:47
 * @LastEditTime: 2022-06-13 13:11:23
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\user\updateUserPassword.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 更新用户密码
func UpdateUserPassword(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := utils.ParseToken(token)
	old_password := utils.MD5(c.PostForm("old_password"))
	new_password := utils.MD5(c.PostForm("new_password"))
	res := global.
		DB.
		Where("id = ?", claims.Id).
		Find(&models.User{})
	if res.RowsAffected > 0 {
		user := models.User{}
		res.Scan(&user)
		if user.Password == old_password {
			user.Password = new_password
			err := global.
				DB.
				Model(models.User{}).
				Where("id = ?", claims.Id).
				Update(&user).
				Error
			if err == nil {
				c.JSON(
					http.StatusOK,
					utils.Msg(utils.Message{Code: 1, Msg: "更新成功"}),
				)
			} else {
				c.JSON(
					http.StatusBadRequest,
					utils.Msg(utils.Message{Code: 0, Msg: "更新失败"}),
				)
			}
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "密码错误"}),
			)
		}
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "用户不存在"}),
		)
	}
}
