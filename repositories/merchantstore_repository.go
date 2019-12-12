package repositories

import "fish.com/fish-common/models"

var MerchantStoreRepository = newMerchantStoreRepository()

func newMerchantStoreRepository() *merchantStoreRepository {
	return &merchantStoreRepository{}
}

type merchantStoreRepository struct {
}
func (this *merchantStoreRepository) Find() []models.MerchantStore{
	list:=make([]models.MerchantStore,0)
	models.GetDB().Find(&list)
	return list
}