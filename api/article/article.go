/**
 * @Author: boyyang
 * @Date: 2022-02-16 10:20:47
 * @LastEditTime: 2022-07-05 17:45:54
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\article\article.go
 * [如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 获取所有文章
func GetArticles(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	_page, _ := strconv.ParseInt(page, 10, 32)
	_limit, _ := strconv.ParseInt(limit, 10, 32)
	var articles []models.Article
	var count int
	if _page == 0 && _limit == 0 {
		global.
			DB.
			Preload("Author").
			Find(&articles).
			Count(&count)
	} else {
		global.
			DB.
			Limit(_limit).
			Offset((_page - 1) * _limit).
			Preload("Author").
			Find(&articles).
			Offset(-1).
			Limit(-1).
			Count(&count)
	}
	c.JSON(
		http.StatusOK,
		utils.Msg(utils.Message{Code: 1, Msg: "获取成功", Data: articles, Count: count}),
	)
}

// 根据id查询
func GetArticleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 0, Msg: "文章id不能为空"}),
		)
		return
	}
	var article models.Article
	global.
		DB.
		Preload("Author").
		Where("id = ?", id).
		Find((&article))
	c.JSON(
		http.StatusOK,
		utils.Msg(utils.Message{Code: 1, Msg: "获取成功", Data: article}),
	)
}

// 添加文章
func AddArticle(c *gin.Context) {
	title := c.PostForm("title")
	subtitle := c.PostForm("subtitle")
	content := c.PostForm("content")
	image := c.PostForm("image")
	token := c.Request.Header["Token"]
	var author_id int
	userInfo, err := utils.ParseToken(token[0])
	if err == nil {
		author_id = userInfo.Id
	}
	if len(title) == 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "文章标题不能为空"}),
		)
	} else {
		article := models.Article{
			Title:    title,
			Subtitle: subtitle,
			Image:    image,
			Content:  content,
			UserID:   author_id,
		}
		err := global.
			DB.
			Create(&article).
			Error
		if err == nil {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "文章添加成功", Data: article.ID}),
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "文章添加失败"}),
			)
		}
	}
}

// 删除文章
func DelArticle(c *gin.Context) {
	id := c.PostForm("id")
	var articel models.Article
	err := global.
		DB.Where("id = ?", id).
		Delete(&articel).
		Error
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

// 文章点赞
func ThumbsUp(c *gin.Context) {
	id := c.Query("id")
	err := global.
		DB.
		Model(&models.Article{}).
		Where("id = ?", id).
		Update("thumbs_up", gorm.Expr("thumbs_up + 1")).
		Error
	if err == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "点赞成功"}),
		)
	}
}
