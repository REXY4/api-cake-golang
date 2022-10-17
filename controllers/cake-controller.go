package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"cake.com/cake/dto"
	"cake.com/cake/helper"
	"cake.com/cake/models"
	"cake.com/cake/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

type CakeController interface {
	All(context *gin.Context)
	FindCakeById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type cakeController struct {
	cakeService services.CakeService
}

func NewCakeController(cakeServ services.CakeService) CakeController {
	return &cakeController{
		cakeService: cakeServ,
	}
}

func (c *cakeController) All(context *gin.Context) {
	var cakes []models.Cake = c.cakeService.All()
	res := helper.BuildResponse(200, "Get All Data Cake Success!", cakes)
	context.JSON(http.StatusOK, res)
}

func (c *cakeController) Insert(context *gin.Context) {
	var cakeCreateDTO dto.CakeCreateDTO
	errDTO := context.ShouldBind(&cakeCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse(400, "Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		errBind := context.MustBindWith(&cakeCreateDTO, binding.FormMultipart)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is multipart",
			})
		}

		file, err := context.FormFile("imageFile")
		// The file cannot be received.
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
		}
		// // Retrieve file information
		extension := filepath.Ext(file.Filename)
		// // Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension
		// // // The file is received, so let's save it
		if err := context.SaveUploadedFile(file, `./uploads/`+newFileName); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
		}
		host := os.Getenv("HOST")
		port := os.Getenv("PORT")
		cakeCreateDTO.Image = host + ":" + port + "/uploads/" + newFileName
		result := c.cakeService.Insert(cakeCreateDTO)
		response := helper.BuildResponse(200, "create data cake success", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *cakeController) Update(context *gin.Context) {
	var cakeUpdateDTO dto.CakeCreateDTO
	errDTO := context.ShouldBind(&cakeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse(400, "Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		errBind := context.MustBindWith(&cakeUpdateDTO, binding.FormMultipart)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is multipart",
			})
		}

		file, err := context.FormFile("imageFile")
		// The file cannot be received.
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
		}
		// // Retrieve file information
		extension := filepath.Ext(file.Filename)
		// // Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension
		// // // The file is received, so let's save it
		if err := context.SaveUploadedFile(file, `./uploads/`+newFileName); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
		}
		host := os.Getenv("HOST")
		port := os.Getenv("PORT")
		//
		var id = context.Param("id")
		number, errConv := strconv.ParseUint(string(id), 10, 64)
		if errConv != nil {
			res := helper.BuildErrorResponse(400, "Failed to Convert Id", errDTO.Error(), helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)
		}
		cakeUpdateDTO.ID = number
		cakeUpdateDTO.Image = host + ":" + port + "/uploads/" + newFileName
		result := c.cakeService.Update(cakeUpdateDTO)
		response := helper.BuildResponse(200, "create data cake success", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *cakeController) FindCakeById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse(400, "No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var cake models.Cake = c.cakeService.FindCakeById(id)
	if (cake == models.Cake{}) {
		res := helper.BuildErrorResponse(400, "Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(200, "OK", cake)
		context.JSON(http.StatusOK, res)
	}
}

func (c *cakeController) Delete(context *gin.Context) {
	var cake models.Cake
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse(400, "Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	cake.ID = id
	c.cakeService.Delete(cake)
	res := helper.BuildResponse(200, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
