package routers

import "github.com/gin-gonic/gin"

func InitRouter(engine *gin.Engine) *gin.Engine {
	// api 接口路由配置
	ApiV1(engine)
	ApiV2(engine)
	ApiV3(engine)
	return engine
}