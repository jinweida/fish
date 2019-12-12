package services

import (
	"fish.com/fish-common/models"
	"fish.com/fish-front-api/repositories"
)

var SysAdvertService = newSysAdvertService()

func newSysAdvertService() *sysAdvertService {
	return &sysAdvertService{}
}

type sysAdvertService struct {
}

func(this *sysAdvertService) Find(location int) ([]models.SysAdvert) {
	return repositories.SysAdvertRepository.Find(location)
}
