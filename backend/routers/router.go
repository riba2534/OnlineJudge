package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/riba2534/OnlineJudge/backend/middleware/jwt"
	"github.com/riba2534/OnlineJudge/backend/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	// 中间件用来校验接口
	apiv1.Use(jwt.JWT())
	{
	}
	return r
}
