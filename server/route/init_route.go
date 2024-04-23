package route

import (
	"xyserver/api"
	"xyserver/api/demo"

	"github.com/gin-gonic/gin"
)

// InitRoute 初始化路由
func InitRoute(aEngine *gin.Engine) {
	aEngine.GET("/", api.ShowHome)

	// 创建路由组
	mV1 := aEngine.Group("/v1")
	{
		mV1.GET("/hello1", demo.HellowDemo)
	}

	mV2 := aEngine.Group("/v2")
	{
		mV2.GET("/hello2", demo.HellowDemo)
	}
}
