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
		router.POST("face_validate", api.FaceValidate)
		router.POST("Register_Pic_Upload", func(c *gin.Context) {
			c.String(200, "11")
		})
		router.GET("Register_Face_Feature", func(c *gin.Context) {
			c.String(200, "123456")
		})
	}

	_ = r.Run(":5001")
}
