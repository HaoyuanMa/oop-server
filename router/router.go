package router

import (
	"github.com/gin-gonic/gin"
	"oop-server/api"
)

func InitRouter() {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(gin.Recovery())

	var router = r.Group("api")
	{
		router.POST("login", api.Login)
	}

	_ = r.Run(":5001")
}
