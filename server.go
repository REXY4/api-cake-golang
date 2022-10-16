package main

import (
	"cake.com/cake/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDataBaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.Run(":5000")
}
