package services

import (
	"fish.com/fish-common/models"
	"fish.com/fish-front-api/repositories"
)

var MerchantStoreService = newMerchantStoreService()

func newMerchantStoreService() *merchantStoreService {
	return &merchantStoreService{}
}

type merchantStoreService struct {
}

func(this *merchantStoreService) Find() ([]models.MerchantStore) {
	return repositories.MerchantStoreRepository.Find()
}
