/**
 * @Author: boyyang
 * @Date: 2022-02-16 10:19:41
 * @LastEditTime: 2022-04-05 15:38:33
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\article.go
 */

package routers

import (
	controller "blog/controller/article"

	"github.com/gin-gonic/gin"
)

func AritcleRouterInit(r *gin.Engine) {
	articleRouters := r.Group("api")
	{
		articleRouters.GET("articles", controller.GetArticles)
		articleRouters.GET("articlesDetail", controller.GetArticleDetail)
		articleRouters.POST("addArticle", controller.AddArticle)
		articleRouters.POST("delArticle", controller.DelArticle)
		articleRouters.GET("thumbsUp", controller.ThumbsUp)
	}
}
