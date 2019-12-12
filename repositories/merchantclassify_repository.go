package repositories

import "fish.com/fish-common/models"

var MerchantClassifyRepository = newMerchantClassifyRepository()

func newMerchantClassifyRepository() *merchantClassifyRepository {
	return &merchantClassifyRepository{}
}

type merchantClassifyRepository struct {
}
func (this *merchantClassifyRepository) Find() []models.MerchantClassify{
	list:=make([]models.MerchantClassify,0)
	models.GetDB().Where("isstop = 0").Order("sort asc").Find(&list)
	return list
}
