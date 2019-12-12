package repositories

import "fish.com/fish-common/models"

var SysConfigRepository = newSysconfigRepository()

func newSysconfigRepository() *sysconfigRepository {
	return &sysconfigRepository{}
}

type sysconfigRepository struct {
}
func (this *sysconfigRepository) Find() []models.SysConfig{
	list:=make([]models.SysConfig,0)
	models.GetDB().Find(&list)
	return list
}