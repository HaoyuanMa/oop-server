package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"oop-server/lib"
	"oop-server/models"
	"oop-server/protocol"
	"strconv"
	"strings"
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

func FaceValidate(c *gin.Context) {
	//var face protocol.FaceValidate
	//_ = c.ShouldBindJSON(&face)
	//user := getUserByFaceId(face.Feature)
	//if user.UserName == face.UserName{
	//	c.JSON(200, gin.H{
	//		"status":  200,
	//		"message": "ok",
	//	})
	//}
	//c.JSON(403, gin.H{
	//	"status":  403,
	//	"message": "failed",
	//})

	c.JSON(200, gin.H{
		"status":  200,
		"message": "ok",
	})
}

func getUserByFaceId(faceId string) models.User {
	var users []models.User
	db := lib.GetDBConn()
	_ = db.Find(&users).Error
	minDistance := -1.0
	resultStaff := users[0]
	for _, user := range users {
		curDistance := calDistance(faceId, user.Feature)
		if curDistance < minDistance {
			minDistance = curDistance
			resultStaff = user
		}
	}
	return resultStaff
}

func calDistance(id string, id2 string) float64 {
	v1 := strings.Fields(id)
	v2 := strings.Fields(id2)
	sum := 0.0
	for i := 0; i < 128; i++ {
		t1, _ := strconv.ParseFloat(v1[i], 64)
		t2, _ := strconv.ParseFloat(v2[i], 64)
		sum += math.Pow(t1-t2, 2)
	}
	return math.Pow(sum, 0.5)
}
