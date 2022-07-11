/**
 * @Author: boyyang
 * @Date: 2022-02-16 10:19:41
 * @LastEditTime: 2022-07-11 13:04:05
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\router\article.go
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
		articleRouters.GET("thumbsUpArticle", api.ThumbsUpArticle)
	}
}
