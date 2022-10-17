package services

import (
	"log"

	"cake.com/cake/dto"
	"cake.com/cake/models"
	"cake.com/cake/repository"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type CakeService interface {
	All() []models.Cake
	Insert(c dto.CakeCreateDTO) models.Cake
	Update(c dto.CakeCreateDTO) models.Cake
	FindCakeById(cakeID uint64) models.Cake
	Delete(b models.Cake)
}

type cakeService struct {
	cakeRepository repository.CakeRepository
	context        *gin.Context
}

func NewCakeService(cakeRepo repository.CakeRepository) CakeService {
	return &cakeService{
		cakeRepository: cakeRepo,
	}
}

func (service *cakeService) All() []models.Cake {
	return service.cakeRepository.AllCake()
}

func (service *cakeService) Insert(c dto.CakeCreateDTO) models.Cake {
	cake := models.Cake{}
	errCake := smapping.FillStruct(&cake, smapping.MapFields(&c))
	if errCake != nil {
		log.Fatalf("Failed map %v: ", errCake)
	}
	res := service.cakeRepository.InsertCake(cake)
	return res
}

func (service *cakeService) Update(b dto.CakeCreateDTO) models.Cake {
	cake := models.Cake{}
	err := smapping.FillStruct(&cake, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.cakeRepository.UpdateCake(cake)
	return res
}

func (service *cakeService) FindCakeById(cakeID uint64) models.Cake {
	return service.cakeRepository.FindCakeById(cakeID)
}

func (service *cakeService) Delete(b models.Cake) {
	service.cakeRepository.DeleteCake(b)
}
