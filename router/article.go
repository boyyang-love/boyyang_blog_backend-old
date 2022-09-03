/**
 * @Author: boyyang
 * @Date: 2022-02-16 10:19:41
 * @LastEditTime: 2022-09-03 13:50:02
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog_server\router\article.go
 */

package routers

import (
	api "blog/api/article"

	"github.com/gin-gonic/gin"
)

func AritcleRouterInit(r *gin.Engine) {
	articleRouters := r.Group("api")
	{
		articleRouters.GET("articles", api.Articles)
		articleRouters.GET("articlesDetail", api.ArticleDetail)
		articleRouters.POST("addArticle", api.AddArticle)
		articleRouters.POST("delArticle", api.DelArticle)
		articleRouters.POST("editArticle", api.EditArticle)
		articleRouters.GET("thumbsUpArticle", api.ThumbsUpArticle)
	}
}
