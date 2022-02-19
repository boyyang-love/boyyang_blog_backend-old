/*
 * @Author: boyyang
 * @Date: 2022-02-16 10:20:47
 * @LastEditTime: 2022-02-19 11:28:39
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\controller\article.go
 */

package controller

import (
	"net/http"
	"strconv"
	"websit/models"
	"websit/setupDatabase"
	"websit/utils"

	"github.com/gin-gonic/gin"
)

// 获取所有文章
func GetArticles(c *gin.Context) {
	var articles []models.Article
	setupDatabase.DB.Preload("Author").Find(&articles)
	utils.ReturnData(200, articles, c)
}

// 根据id查询
func GetArticleDetail(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "id为必填项",
		})
		return
	}

	var article models.Article
	setupDatabase.DB.Preload("Author").Where("id = ?", id).Find((&article))
	utils.ReturnData(200, article, c)
}

// 添加文章
func AddArticle(c *gin.Context) {
	title := c.PostForm("title")
	subtile := c.PostForm("subtitle")
	content := c.PostForm("content")
	author_id := c.PostForm("author")

	if len(title) == 0 || len(author_id) == 0 {
		utils.ReturnData(0, "标题以及作者id不能为空", c)
	} else {
		id, err := strconv.Atoi(author_id)

		if err != nil {
			utils.ReturnData(0, "错误", c)
		}

		article := models.Article{
			Title:    title,
			Subtitle: subtile,
			Content:  content,
			UserID:   id,
		}

		res := setupDatabase.DB.Create(&article)
		utils.ReturnData(200, res, c)
	}
}
