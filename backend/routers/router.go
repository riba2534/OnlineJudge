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

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://127.0.0.1:8080"}
	// config.AllowCredentials = true
	// r.Use(cors.New(config))

	apiv1 := r.Group("/api")

	apiv1.POST("/userdata", api.UserData)
	apiv1.POST("/register", api.Register)
	apiv1.POST("/login", api.Login)
	// 中间件用来校验接口
	apiv1.Use(jwt.JWT())
	{
	}
	return r
}
