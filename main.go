package main

import (
	"oop-server/lib"
	"oop-server/models"
	"oop-server/router"
)

func main() {
	lib.InitDb()
	_ = lib.GetDBConn().AutoMigrate(&models.User{})

	router.InitRouter()

}
