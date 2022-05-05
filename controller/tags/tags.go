/**
 * @Author: boyyang
 * @Date: 2022-04-29 11:18:05
 * @LastEditTime: 2022-04-29 17:42:49
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\tags\tags.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package controllers

import (
	"blog/models"
	"blog/setupDatabase"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 添加tags
func AddTags(c *gin.Context) {
	tag_name := c.PostForm("tag_name")
	tags := models.ImagesTag{
		TagName: tag_name,
	}
	var err error
	if tag_name != "" {
		res := setupDatabase.DB.Where("tag_name = ?", tag_name).Find(&tags)
		if res.RowsAffected == 0 {
			setupDatabase.DB.Create(&tags)
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "标签添加成功"}, err))
		} else {
			c.JSON(http.StatusBadRequest, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "标签名称已经存在"}, err))
		}
	} else {
		c.JSON(http.StatusBadRequest, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "请输入标签名称"}, err))
	}
}
