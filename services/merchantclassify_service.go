package services

import (
	"fish.com/fish-common/models"
	"fish.com/fish-front-api/repositories"
)

var MerchantClassifyService = newMerchantClassifyService()

func newMerchantClassifyService() *merchantClassifyService {
	return &merchantClassifyService{}
}

type merchantClassifyService struct {
}

func(this *merchantClassifyService) Find() ([]models.MerchantClassify) {
	return repositories.MerchantClassifyRepository.Find()
}
