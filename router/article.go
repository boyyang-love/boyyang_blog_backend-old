/**
 * @Author: boyyang
 * @Date: 2022-02-16 10:19:41
 * @LastEditTime: 2022-02-18 14:09:21
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\router\article.go
 */

package routers

import (
	"websit/controller"

	"github.com/gin-gonic/gin"
)

func AritcleRouterInit(r *gin.Engine) {
	articleRouters := r.Group("api")
	{
		articleRouters.GET("articles", controller.GetArticles)
		articleRouters.GET("articlesDetail", controller.GetArticleDetail)
		articleRouters.POST("addArticle", controller.AddArticle)
		articleRouters.POST("delArticle", controller.DelArticle)
	}
}
