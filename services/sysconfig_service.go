package services

import (
	"fish.com/fish-common/models"
	"fish.com/fish-front-api/repositories"
)

var SysConfigService = newSysConfigService()

func newSysConfigService() *sysconfigService {
	return &sysconfigService{}
}

type sysconfigService struct {
}

func(this *sysconfigService) Find() ([]models.SysConfig) {
	return repositories.SysConfigRepository.Find()
}