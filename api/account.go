package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oop-server/lib"
	"oop-server/models"
	"oop-server/protocol"
)

func Login(c *gin.Context) {
	var login protocol.Login
	_ = c.ShouldBindJSON(&login)
	db := lib.GetDBConn()
	var user models.User
	err := db.Where("user_name=?", login.UserName).First(&user).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(403, gin.H{
			"status":  403,
			"message": "failed",
		})
		return
	}
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
