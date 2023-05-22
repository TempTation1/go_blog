package router

import (
	v1 "github.com/Temptation1/blog-service/go-programming-tour-book/blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery()) //先注册自有的recovery和logger中间件

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("api/v1") //handler bashpath engine
	{
		apiv1.POST("/tags", tag.Create)            //新增标签
		apiv1.DELETE("/tags/:id", tag.Delete)      //删除指定id标签
		apiv1.PUT("/tags/:id", tag.Update)         //更新指定标签
		apiv1.PATCH("/tags/:id/state", tag.Update) //更新指定标签的一个小状态
		apiv1.GET("/tags/", tag.List)              //获取标签列表

		apiv1.POST("/articles", article.Create)            //新增文章
		apiv1.DELETE("/articles/:id", article.Delete)      //删除指定文章
		apiv1.PUT("/articles/:id", article.Update)         //更新指定文章
		apiv1.PATCH("/articles/:id/state", article.Update) //更新指定文章小状态
		apiv1.GET("/articles/:id", article.Get)            //获取指定文章
		apiv1.GET("/articles", article.List)               //获取文章列表
	}

	return r
}