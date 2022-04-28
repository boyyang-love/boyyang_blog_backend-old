/**
 * @Author: boyyang
 * @Date: 2022-02-14 17:01:53
 * @LastEditTime: 2022-04-23 17:57:39
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\user\user.go
 */
package controller

import (
	"blog/models"
	"blog/setupDatabase"
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
	data := setupDatabase.
		DB.
		Where("Username = ? AND ID != ? ", form.Username, claims.Id).
		First(&user)
	// 如果用户存
	if data.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "该用户名已经存在"}, nil))
		return
	}
	err := setupDatabase.
		DB.
		Model(models.User{}).
		Where("id = ?", claims.Id).
		Update(&form).
		Error
	if err == nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "用户信息更新成功"}, nil))
	} else {
		c.JSON(http.StatusBadRequest, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "用户信息更新失败"}, err))
	}
}
