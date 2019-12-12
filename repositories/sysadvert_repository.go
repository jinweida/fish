package repositories

import "fish.com/fish-common/models"

var SysAdvertRepository = newSysAdvertRepository()

func newSysAdvertRepository() *sysAdvertRepository {
	return &sysAdvertRepository{}
}

type sysAdvertRepository struct {
}
func (this *sysAdvertRepository) Find(location int) []models.SysAdvert{
	list:=make([]models.SysAdvert,0)
	models.GetDB().Find(&list).Order("sort asc").Where("location = ?",location)
	return list
}
