package routers

import (
	"github.com/gin-gonic/gin"

	"dgg/controllers/ApiV1Controllers"
	"dgg/middleware"
)

// ApiV1Controllers 接口v1版本路由
func ApiV1(engine *gin.Engine) {
	g := engine.Group("/api/v1").Use(middleware.MyMiddleware())
	{
		// 简单验证示例
		//g.POST("/demo/validator/post", demo.Validator)
		g.GET("/", ApiV1Controllers.Index)

		g.POST("/setAdmin", ApiV1Controllers.SetAdmin)

	}
}
