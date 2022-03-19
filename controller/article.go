/**
 * @Author: boyyang
 * @Date: 2022-02-16 10:20:47
 * @LastEditTime: 2022-02-19 11:28:39
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\controller\article.go
 */

package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"websit/models"
	"websit/setupDatabase"
	"websit/utils"

	"github.com/gin-gonic/gin"
)

// 获取所有文章
func GetArticles(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	_page, _ := strconv.ParseInt(page, 10, 32)
	_limit, _ := strconv.ParseInt(limit, 10, 32)
	var articles []models.Article
	var count int
	setupDatabase.DB.Limit(_limit).Offset((_page - 1) * _limit).Preload("Author").Find(&articles).Offset(-1).Limit(-1).Count(&count)
	c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "获取成功", Count: count}, articles))
}

// 根据id查询
func GetArticleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "id为必填项"}, nil))
		return
	}
	var article models.Article
	setupDatabase.DB.Preload("Author").Where("id = ?", id).Find((&article))
	c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "获取成功"}, article))
}

// 添加文章
func AddArticle(c *gin.Context) {
	title := c.PostForm("title")
	subtile := c.PostForm("subtitle")
	content := c.PostForm("content")
	image := c.PostForm("image")
	token := c.Request.Header["Token"]
	var author_id int
	userInfo, err := utils.ParseToken(token[0])
	if err == nil {
		author_id = userInfo.Id
	}
	if len(title) == 0 {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "文章标题不能为空"}, nil))
	} else {
		article := models.Article{
			Title:    title,
			Subtitle: subtile,
			Image:    image,
			Content:  content,
			UserID:   author_id,
		}
		res := setupDatabase.DB.Create(&article)
		fmt.Println(res.Value)
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "文章添加成功"}, res.Value))
	}
}
