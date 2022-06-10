/**
 * @Author: boyyang
 * @Date: 2022-02-16 10:19:41
 * @LastEditTime: 2022-06-10 15:40:46
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\article.go
 */

package routers

import (
	api "blog/api/article"

	"github.com/gin-gonic/gin"
)

func AritcleRouterInit(r *gin.Engine) {
	articleRouters := r.Group("api")
	{
		articleRouters.GET("articles", api.GetArticles)
		articleRouters.GET("articlesDetail", api.GetArticleDetail)
		articleRouters.POST("addArticle", api.AddArticle)
		articleRouters.POST("delArticle", api.DelArticle)
		articleRouters.GET("thumbsUp", api.ThumbsUp)
	}
}
