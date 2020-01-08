package controllers

import (
	e "beego/app/recover"
	"beego/services"
)

// @TagName 城市基本信息
// @Description 城市基础模块
type CityAllController struct {
	BaseController
}

// URLMapping ...
func (c *CityAllController) URLMapping() {
	c.Mapping("GetByBigAreaId", c.GetListByBigAreaId)
	c.Mapping("GetByProvinceId", c.GetListByProvinceId)
	c.Mapping("SearchListByEname", c.SearchListByEname)
	c.Mapping("GetProfile", c.GetProfile)
	c.Mapping("GetProfileByEname", c.GetProfileByEname)
	c.Mapping("GetProfileByCode", c.GetProfileByCode)
}

// @Title GetProfile
// @Summary 城市ID查城市基本信息
// @Description 根据城市ID获取城市基本信息
// @Param cityid query int true "城市ID"
// @Success 200 {object} vo.CityAllVo
// @Failure 500 system error
// @router /get-profile [get]
func (c *CityAllController) GetProfile() {
	cityId, _ := c.GetInt("cityid", 0)
	if cityId <= 0 {
		c.Failure(e.ParamError, "param error", "参数错误")
	}
	profile := services.NewCityAllService().FindByPrimaryKey(cityId)

	c.Success(profile)
}

// @Title GetProfileByEname
// @Summary 城市拼音查城市基本信息
// @Description 根据城市拼音获取城市基本信息
// @Param ename query string true "城市拼音"
// @Success 200 {object} vo.CityAllVo
// @Failure 500 system error
// @router /get-profile-by-ename [get]
func (c *CityAllController) GetProfileByEname() {
	ename := c.GetString("ename", "")
	if ename == "" {
		c.Failure(e.ParamError, "param error", "参数错误")
	}
	profile := services.NewCityAllService().FindByEname(ename)

	c.Success(profile)
}

// @Title GetProfileByEname
// @Summary 城市代码查城市基本信息
// @Description 根据城市代码获取城市基本信息
// @Param city_code query string true "城市代码"
// @Success 200 {object} vo.CityAllVo
// @Failure 500 system error
// @router /get-profile-by-code [get]
func (c *CityAllController) GetProfileByCode() {
	cityCode := c.GetString("city_code", "")
	if cityCode == "" {
		c.Failure(e.ParamError, "param error", "参数错误")
	}
	profile := services.NewCityAllService().FindByCityCode(cityCode)

	c.Success(profile)
}

// @Title GetListByBigAreaId
// @Summary 大区ID查城市列表
// @Description 根据大区ID获取城市列表
// @Param big_areaid query int true "大区ID"
// @Success 200 {object} vo.CityAllListVo
// @Failure 500 system error
// @router /get-list-by-bigareaid [get]
func (c *CityAllController) GetListByBigAreaId() {
	bigAreaId, _ := c.GetInt("big_areaid", 0)
	if bigAreaId <= 0 {
		c.Failure(e.ParamError, "param error", "参数错误")
	}
	list := services.NewCityAllService().FindListByBigAreaId(bigAreaId)

	c.Success(list)
}

// @Title GetByProvinceId
// @Summary 省份ID查城市列表
// @Description 根据所属省份ID获取城市列表
// @Param provinceid query int true "省份ID"
// @Success 200 {object} vo.CityAllListVo
// @Failure 500 system error
// @router /get-list-by-provinceid [get]
func (c *CityAllController) GetListByProvinceId() {
	provinceId, _ := c.GetInt("provinceid", 0)
	if provinceId <= 0 {
		c.Failure(e.ParamError, "param error", "参数错误")
	}
	list := services.NewCityAllService().FindListByProvinceId(provinceId)

	c.Success(list)
}

// @Title SearchListByEname
// @Summary 搜索城市拼音前匹配列表
// @Description 根据城市拼音获取符合前匹配的城市列表
// @Param ename query string true "城市拼音"
// @Success 200 {object} vo.CityAllListVo
// @Failure 500 system error
// @router /search-list-by-ename [get]
func (c *CityAllController) SearchListByEname() {
	ename := c.GetString("ename", "")
	if ename == "" {
		c.Failure(e.ParamError, "param error", "参数错误")
	}
	list := services.NewCityAllService().FindListByStartWithEname(ename)

	c.Success(list)
}
