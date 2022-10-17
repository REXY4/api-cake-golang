package repository

import (
	"cake.com/cake/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CakeRepository interface {
	AllCake() []models.Cake
	InsertCake(c models.Cake) models.Cake
	UpdateCake(c models.Cake) models.Cake
	FindCakeById(cakeID uint64) models.Cake
	DeleteCake(b models.Cake)
}

type cakeConnection struct {
	connection *gorm.DB
	context    *gin.Context
}

func NewCakeRepository(dbConn *gorm.DB) CakeRepository {
	return &cakeConnection{
		connection: dbConn,
	}
}

func (db *cakeConnection) AllCake() []models.Cake {
	var cake []models.Cake
	db.connection.Preload("Cake").Find(&cake)
	return cake
}

func (db *cakeConnection) InsertCake(c models.Cake) models.Cake {
	db.connection.Save(&c)
	return c
}

func (db *cakeConnection) UpdateCake(c models.Cake) models.Cake {
	db.connection.Save(&c)
	return c
}

func (db *cakeConnection) FindCakeById(cakeID uint64) models.Cake {
	var cake models.Cake
	db.connection.Preload("User").Find(&cake, cakeID)
	return cake
}

func (db *cakeConnection) DeleteCake(b models.Cake) {
	db.connection.Delete(&b)
}
