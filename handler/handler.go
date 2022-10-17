package handler

import (
	"cake.com/cake/config"
	"cake.com/cake/controllers"
	"cake.com/cake/repository"
	"cake.com/cake/services"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	cakeRepository repository.CakeRepository  = repository.NewCakeRepository(db)
	cakeService    services.CakeService       = services.NewCakeService(cakeRepository)
	cakeController controllers.CakeController = controllers.NewCakeController(cakeService)
)

func NewHandler(c *Config) {
	// defer config.CloseDatabaseConnection(db)

	cakeRoutes := c.R.Group("/api/v1")
	{
		cakeRoutes.GET("/cake", cakeController.All)
		cakeRoutes.POST("/cake", cakeController.Insert)
		cakeRoutes.PATCH("/cake/:id", cakeController.Update)
		cakeRoutes.GET("/cake/:id", cakeController.FindCakeById)
		cakeRoutes.DELETE("/cake/:id", cakeController.Delete)
	}

}
