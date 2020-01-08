package services

import (
	"beego/models"
	"fmt"
)

var cityServiceImpl ICityService

func NewCityService() ICityService {
	if cityServiceImpl == nil {
		cityServiceImpl = &CityService{}
	}

	return cityServiceImpl
}

var CityServiceImpl ICityService

type ICityService interface {
	Add(city models.City) (int64, error)
	FindProfile(where map[string]interface{}) map[string]interface{}
	UpdateById(cityid int64, cityInfo map[string]interface{}) (int64, error)
}

type CityService struct {
}

func NewCItyService() ICityService {
	if CityServiceImpl == nil {
		CityServiceImpl = &CityService{}
	}

	return CityServiceImpl
}

// GetCityProfile 根据where条件中的条件获取城市数据信息
func (s *CityService) FindProfile(where map[string]interface{}) map[string]interface{} {
	var cityInfo = models.City{}
	var oneCityInfo map[string]interface{}
	// city表数据处理
	var cityFields = models.CityReturnFields
	err := models.NewCityDao().SelectOne(&cityInfo, where, cityFields...)
	if err != nil {
		return oneCityInfo
	}
	oneCityInfo = models.OneResultHandle(cityInfo, cityFields)
	return oneCityInfo
}

// FindList 根据where条件获取城市信息列表
func (s *CityService) FindList(where map[string]interface{}) []interface{} {
	var returnList []interface{}
	var cityInfoList []models.City
	// city表数据处理
	var cityFields = models.CityReturnFields
	_ = models.NewCityDao().SelectList(&cityInfoList, where, cityFields)
	fmt.Println(cityInfoList)
	if len(cityInfoList) <= 0 {
		return returnList
	}
	return models.ResultHandle(s.modelTransInterface(cityInfoList, returnList), cityFields)
}

// FindPageList 根据where条件分页获取城市信息列表
func (s *CityService) FindPageList(where map[string]interface{}, orderBy, groupBy []string, page, limit uint16) *models.Pagination {
	var resultList []interface{}
	var cityInfoList []models.City
	var pageCityList *models.Pagination
	// city表数据处理
	var cityFields = models.CityReturnFields
	pageCityList, _ = models.NewCityDao().SelectPageList(&cityInfoList, where, cityFields, orderBy, groupBy, page, limit)
	cityHandleInterface := s.modelTransInterface(cityInfoList, resultList)
	if len(cityHandleInterface) == 0 {
		return pageCityList
	}
	pageCityList.List = models.ResultHandle(cityHandleInterface, cityFields)
	return pageCityList
}

// modelTransInterface 结构体切片转interface切片.
func (s *CityService) modelTransInterface(cityModelList []models.City, cityInterface []interface{}) []interface{} {
	for _, city := range cityModelList {
		cityInterface = append(cityInterface, city)
	}
	return cityInterface
}

func (s *CityService) Add(city models.City) (int64, error) {
	return models.NewCityDao().Insert(&city)
}

func (s *CityService) UpdateById(cityid int64, cityInfo map[string]interface{}) (int64, error) {
	model := models.NewCityDao()
	where := make(map[string]interface{})
	where["cityid"] = cityid
	return model.Update(where, cityInfo)
}
