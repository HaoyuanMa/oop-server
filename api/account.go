package api

import (
	"github.com/gin-gonic/gin"
	"oop-server/lib"
	"oop-server/models"
	"oop-server/protocol"
)

func Login(c *gin.Context) {
	var login protocol.Login
	_ = c.ShouldBindJSON(login)
	db := lib.GetDBConn()
	var user models.User
	db.Where("user_name = ?", user.UserName).Find(user)
	if user.Password == login.PassWord {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "ok",
		})
	} else {
		c.JSON(403, gin.H{
			"status":  403,
			"message": "failed",
		})
	}
}
