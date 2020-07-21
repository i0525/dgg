package routers

import (
	"github.com/gin-gonic/gin"
	"lidongheDemo/controllers/ApiV1Controllers"
	"lidongheDemo/middleware"
)


// ApiV1Controllers 接口v1版本路由
func ApiV1(engine *gin.Engine) {
	g := engine.Group("/api/v1").Use(middleware.MyMiddleware())
	{

		// 简单验证示例
		//g.POST("/demo/validator/post", demo.Validator)
		g.GET("/Index", ApiV1Controllers.Index )
	}
}

// ApiV2 接口v2版本路由
func ApiV2(engine *gin.Engine) {
	//g := engine.Group("/api/v2")
	//{
	//	g.GET("/demo/json")
	//}
}

// ApiV3 接口v3版本路由
func ApiV3(engine *gin.Engine) {
	//g := engine.Group("/api/v3")
	//{
	//	g.GET("/demo/json")
	//}
}