// @Time : 2019/9/10 16:05
// @Author : duanqiangwen
// @File : city_all_service
// @Software: GoLand
package services

import "beego/models"

var cityAllServiceImpl ICityAllService

type ICityAllService interface {
	FindByPrimaryKey(cityId int) map[string]interface{}
	FindByEname(ename string) map[string]interface{}
	FindByCityCode(cityCode string) map[string]interface{}
	FindListByBigAreaId(bigAreaId int) []interface{}
	FindListByProvinceId(provinceId int) []interface{}
	FindListByStartWithEname(ename string) []interface{}
	Add(md models.CityAll) (int64, error)
	UpdateById(cityid int64, cityAllInfo map[string]interface{}) (int64, error)
}

type CityAllService struct {
}

func NewCityAllService() ICityAllService {
	if cityAllServiceImpl == nil {
		cityAllServiceImpl = &CityAllService{}
	}

	return cityAllServiceImpl
}

// FindByPrimaryKey 根据主键城市ID获取城市基本信息.
func (s *CityAllService) FindByPrimaryKey(cityId int) map[string]interface{} {
	return s.findOne(map[string]interface{}{
		"cityid": cityId,
	})
}

// FindByEname 根据城市拼音获取城市基本信息
func (s *CityAllService) FindByEname(ename string) map[string]interface{} {
	return s.findOne(map[string]interface{}{
		"ename": ename,
	})
}

// FindByEname 根据城市代码获取城市基本信息
func (s *CityAllService) FindByCityCode(cityCode string) map[string]interface{} {
	return s.findOne(map[string]interface{}{
		"city_code": cityCode,
	})
}

// GetList 根据大区ID获取城市列表.
func (s *CityAllService) FindListByBigAreaId(bigAreaId int) []interface{} {
	return s.findList(map[string]interface{}{
		"big_areaid": bigAreaId,
	})
}

// GetList 根据省份ID获取城市列表.
func (s *CityAllService) FindListByProvinceId(provinceId int) []interface{} {
	return s.findList(map[string]interface{}{
		"provinceid": provinceId,
	})
}

// GetList 获取以城市拼音开头的城市列表.
func (s *CityAllService) FindListByStartWithEname(ename string) []interface{} {
	return s.findList(map[string]interface{}{
		"ename__istartswith": ename,
	})
}

// findList 根据条件获取城市基本信息.
func (s *CityAllService) findOne(where map[string]interface{}) map[string]interface{} {
	var (
		cityAllProfile map[string]interface{}
		err            error
	)
	cityAllModel := models.CityAll{}
	fields := models.CityAllReturnFields
	err = models.NewCityAllDao().SelectOne(&cityAllModel, where, fields...)
	if err != nil {
		return cityAllProfile
	}

	return models.OneResultHandle(cityAllModel, fields)
}

// findList 根据条件获取城市列表.
func (s *CityAllService) findList(where map[string]interface{}) []interface{} {
	var (
		cityAllModelList []models.CityAll
		cityAllInterface []interface{}
	)
	fields := models.CityAllReturnFields
	_ = models.NewCityAllDao().SelectList(&cityAllModelList, where, fields)
	if len(cityAllModelList) <= 0 {
		return cityAllInterface
	}

	return models.ResultHandle(s.modelTransInterface(cityAllModelList, cityAllInterface), fields)
}

// modelTransInterface 结构体切片转interface切片.
func (s *CityAllService) modelTransInterface(cityAllModelList []models.CityAll, cityAllInterface []interface{}) []interface{} {
	for _, cityAll := range cityAllModelList {
		cityAllInterface = append(cityAllInterface, cityAll)
	}
	return cityAllInterface
}

func (s *CityAllService) Add(md models.CityAll) (int64, error) {
	return models.NewCityAllDao().Insert(&md)
}

func (s *CityAllService) UpdateById(cityid int64, cityAllInfo map[string]interface{}) (int64, error) {
	model := models.NewCityAllDao()
	where := make(map[string]interface{})
	where["cityid"] = cityid
	return model.Update(where, cityAllInfo)
}
