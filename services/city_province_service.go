package services

import (
	"beego/models"
)

type ICityProvinceService interface {
	FindByProvincename(provincename string) map[string]interface{}
	findOne(where map[string]interface{}) map[string]interface{}
}

var cityProvinceServiceImpl ICityProvinceService

type CityProvinceService struct {
}

func NewCityProvinceService() ICityProvinceService {
	if cityProvinceServiceImpl == nil {
		cityProvinceServiceImpl = &CityProvinceService{}
	}

	return cityProvinceServiceImpl
}

func (s *CityProvinceService) FindByProvincename(provincename string) map[string]interface{} {
	return s.findOne(map[string]interface{}{
		"provincename": provincename,
	})
}

func (s *CityProvinceService) findOne(where map[string]interface{}) map[string]interface{} {
	var (
		cityProvinceProfile map[string]interface{}
		err                 error
	)
	cityProvinceModel := models.CityProvince{}
	fields := models.CityProvinceReturnFields
	err = models.NewCityProvinceDao().SelectOne(&cityProvinceModel, where, fields...)

	if err != nil {
		return cityProvinceProfile
	}

	return models.OneResultHandle(cityProvinceModel, fields)
}
