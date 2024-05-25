package routes

import (
	"github.com/NSDN/nya-server/controllers"
	"github.com/gin-gonic/gin"
)

// 注册帖子相关路由。
// 应当在初始化路由的函数中调用它，并传入一个路由引擎实例的指针。
func registerTopicRoutes(router *gin.Engine) {
	regitsereQueryArticlesRoute(router)
	registerCreateArticleRoute(router)
}

/** 注册请求帖文列表的路由 */
func regitsereQueryArticlesRoute(router *gin.Engine) {
	// 此处的 `:plate` 应为版块的 `id`
	router.GET("/articles/:plate", controllers.GetArticles)
}

/** 注册创建文章帖子的路由 */
func registerCreateArticleRoute(router *gin.Engine) {
	router.POST("/article", controllers.CreateArticle)
}
