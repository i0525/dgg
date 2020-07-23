package routers

import (
	"dgg/app/Controller/v1"
	"github.com/gin-gonic/gin"

	"dgg/middleware"
)

// Controller 接口v1版本路由
func UserApi(engine *gin.Engine) {
	g := engine.Group("/api/v1").Use(middleware.MyMiddleware())
	{
		// 简单验证示例
		//g.POST("/demo/validator/post", demo.Validator)
		g.GET("/", v1.Index)

		g.POST("/setAdmin", v1.SetAdmin)

	}
}
