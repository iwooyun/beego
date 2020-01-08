package controllers

// @TagName City
// @Description CityController operations for City
import (
	"beego/app/recover"
	"beego/library/utils/base"
	"beego/models"
	"beego/services"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// @TagName 城市扩展
// @Description  业务城市模块 有关城市的业务属性字段
type CityController struct {
	BaseController
}

// URLMapping ...
func (c *CityController) URLMapping() {
	c.Mapping("GetProfileById", c.GetProfileById)
	c.Mapping("GetListByAreaId", c.GetListByAreaId)
	c.Mapping("GetListByProvinceId", c.GetListByProvinceId)
	c.Mapping("GetListByEname", c.GetListByEname)
	c.Mapping("GetAllPageList", c.GetAllPageList)
	c.Mapping("Add", c.Add)
	c.Mapping("Update", c.Update)
}

// @Title GetProfileById
// @Summary 【单个】获取城市基本信息(城市id)
// @Description 使用单个城市id获取城市的基本信息列表
// @Param cityid query int true "城市id"
// @Success 200 {object} vo.CityVo
// @Failure 500 system error
// @router /get-profile-by-id [get]
func (c *CityController) GetProfileById() {
	var where = make(map[string]interface{})
	var cityId = c.GetString("cityid")
	if cityId == "" {
		c.Failure(recover.ParamError, "参数不能为空", [...]string{})
	}
	var cityNum, _ = strconv.Atoi(cityId)
	where["cityid"] = cityNum
	var cityService = services.CityService{}
	resultCityProfile := cityService.FindProfile(where)
	c.Success(resultCityProfile)
}

// @Title GetProfileByName
// @Summary 【单个】获取城市基本信息(城市名)
// @Description 使用城市中文名获取城市的基本信息
// @Param name query string true "城市中文名"
// @Success 200 {object} vo.CityVo
// @Failure 500 system error
// @router /get-profile-by-name [get]
func (c *CityController) GetProfileByName() {
	var where = make(map[string]interface{})
	var cityName = c.GetString("name")
	if cityName == "" {
		c.Failure(recover.ParamError, "参数不能为空", [...]string{})
	}
	where["cityname"] = cityName
	var cityService = services.CityService{}
	resultCityProfile := cityService.FindProfile(where)
	c.Success(resultCityProfile)
}

// @Title GetProfileByEname
// @Summary 【单个】获取城市基本信息(城市拼音名)
// @Description 使用城市拼音名获取城市的基本信息
// @Param ename query string true "城市拼音名"
// @Success 200 {object} vo.CityVo
// @Failure 500 system error
// @router /get-profile-by-ename [get]
func (c *CityController) GetProfileByEname() {
	var where = make(map[string]interface{})
	var cityEName = c.GetString("ename")
	if cityEName == "" {
		c.Failure(recover.ParamError, "参数不能为空", [...]string{})
	}
	where["ename"] = cityEName
	var cityService = services.CityService{}
	resultCityProfile := cityService.FindProfile(where)
	c.Success(resultCityProfile)
}

// @Title GetListByAreaId
// @Summary 【批量】获取城市基本信息列表(大区)
// @Description 使用单个大区id获取城市的基本信息列表
// @Param area_id query int true "大区id"
// @Success 200 {object} vo.CityVo
// @Failure 500 system error
// @router /get-list-by-area-id [get]
func (c *CityController) GetListByAreaId() {
	var where = make(map[string]interface{})
	var areaId = c.GetString("area_id")
	if areaId == "" {
		c.Failure(recover.ParamError, "参数不能为空", [...]string{})
	}
	var areaNum, _ = strconv.Atoi(areaId)
	where["areaid"] = areaNum
	var cityService = services.CityService{}
	resultCityList := cityService.FindList(where)
	c.Success(resultCityList)
}

// @Title GetListByProvinceId
// @Summary 【批量】获取城市基本信息列表(省份)
// @Description 使用单个省份id获取城市的基本信息列表
// @Param ProvinceId query int true "省份id"
// @Success 200 {object} vo.CityVo
// @Failure 500 system error
// @router /get-list-by-province-id [get]
func (c *CityController) GetListByProvinceId() {
	var where = make(map[string]interface{})
	var provinceId = c.GetString("province_id")
	if provinceId == "" {
		c.Failure(recover.ParamError, "参数不能为空", [...]string{})
	}
	var provinceIdNum, _ = strconv.Atoi(provinceId)
	where["provinceid"] = provinceIdNum
	var cityService = services.CityService{}
	resultCityList := cityService.FindList(where)
	c.Success(resultCityList)
}

// @Title GetListByEname
// @Summary 【批量】获取城市基本信息列表(拼音)
// @Description 使用城市拼音首字母获取城市的基本信息列表
// @Param ename query string true "拼音名首字母"
// @Success 200 {object} vo.CityVo
// @Failure 500 system error
// @router /get-list-by-ename [get]
func (c *CityController) GetListByEname() {
	var where = make(map[string]interface{})
	var ename = c.GetString("ename")
	if ename == "" {
		c.Failure(recover.ParamError, "参数不能为空", [...]string{})
	}
	where["ename__istartswith"] = ename
	var cityService = services.CityService{}
	resultCityList := cityService.FindList(where)
	c.Success(resultCityList)
}

// @Title GetAllPageList
// @Summary 【分页】获取城市基本信息列表
// @Description 分页获取全部城市数据
// @Param query query string false "查询条件"
// @Param page query uint16 true "当前页码"
// @Param limit query uint16 true "每页数量"
// @Success 200 {object} vo.CityPageListVo
// @Failure 500 system error
// @router /get-all-page-list [get]
func (c *CityController) GetAllPageList() {
	var cityService = services.CityService{}
	var where map[string]interface{}
	var groupBy []string
	orderBy := []string{"-cityid"}
	page, _ := c.GetUint16("page")
	limit, _ := c.GetUint16("limit")
	query := c.GetString("query")
	if query != ""{
		if err := json.Unmarshal([]byte(query), &where); err != nil {
			c.Failure(recover.ParamError, "参数有误", [...]string{})
		}
	}
	resultCityList := cityService.FindPageList(where, orderBy, groupBy, page, limit)
	c.Success(resultCityList)
}


// @Title 新增城市
// @Summary 添加城市
// @Description 添加城市
// @Param areaid  query  int true "大区id"
// @Param province_name query  string true "省份名称"
// @Param cityname query  string true "城市名称"
// @Param ename query  string true "城市拼音"
// @Param shortname query  int true "城市缩写"
// @Param longitude query  string true "径度"
// @Param latitude query 	 string true "维度"
// @Param city_code query 	 string true "城市编码"
// @Success 200 {object} vo.ApiResponse
// @Failure 500 system error
// @router /add [post]
func (c *CityController) Add() {
	areaid, _ := c.GetInt8("areaid")
	province_name := c.GetString("province_name")
	cityname := c.GetString("cityname")
	ename := c.GetString("ename")
	shortname := c.GetString("shortname")
	longitude, err := strconv.ParseFloat(c.GetString("longitude"), 64)
	if err != nil {
		c.Failure(recover.ParamError, "longitude:"+err.Error(), [...]string{})
		return
	}
	latitude, err := strconv.ParseFloat(c.GetString("latitude"), 64)
	if err != nil {
		c.Failure(recover.ParamError, "latitude:"+err.Error(), [...]string{})
		return
	}
	city_code := c.GetString("city_code")

	city := models.City{}
	cityAll := models.CityAll{}

	if areaid == 0 {
		c.Failure(recover.ParamError, "请输入大区id:areaid", [...]string{})
		return
	}

	if ename == "" {
		c.Failure(recover.ParamError, "请输入城市拼音名:ename", [...]string{})
		return
	}

	if cityname == "" {
		c.Failure(recover.ParamError, "请输入城市名:cityname", [...]string{})
		return
	}
	if city_code == "" {
		c.Failure(recover.ParamError, "请输入编码:city_code", [...]string{})
		return
	}
	if shortname == "" {
		c.Failure(recover.ParamError, "请输入城市简称:shortname", [...]string{})
		return
	}

	if province_name == "" {
		c.Failure(recover.ParamError, "请输入省份名:province_name", [...]string{})
		return
	} else {
		//查找省份信息
		//provinceInfo := models.CityProvince{}
		province_name = strings.Replace(province_name, "省", "", -1)
		province_name = strings.Replace(province_name, "市", "", -1)
		cityProvinceService := services.CityProvinceService{}
		provinceInfo := cityProvinceService.FindByProvincename(province_name)

		if provinceInfo == nil {
			c.Failure(recover.ParamError, "输入省份错误:province_name", [...]string{})
			return
		}

		city.Provinceid = provinceInfo["provinceid"].(int)
		city.BigAreaid = provinceInfo["big_areaid"].(int8)
	}

	city.Areaid = areaid
	city.Ename = ename
	city.Cityname = cityname
	city.Shortname = shortname
	city.Longitude = longitude
	city.Latitude = latitude
	city.Updatetime = time.Now()
	cityAll.Areaid = int(areaid)
	cityAll.BigAreaid = uint(city.BigAreaid)
	cityAll.Provinceid = uint(city.Provinceid)
	cityAll.Cityname = cityname
	cityAll.Ename = ename
	cityAll.Longitude = longitude
	cityAll.Latitude = latitude
	cityAll.CityCode = city_code

	ormer, _ := models.NewCityDao().NewOrmer(&city)
	ormer.Begin()
	cityId, err := services.NewCityService().Add(city)
	if err != nil {
		ormer.Rollback()
		c.Failure(recover.Failed, err.Error(), "")
		return
	}
	cityAll.Id = base.Int64ToInt(cityId)
	cityId, cityAllErr := services.NewCityAllService().Add(cityAll)
	if cityAllErr != nil {
		ormer.Rollback()
		c.Failure(recover.Failed, cityAllErr.Error(), "")
		return
	}

	ormer.Commit()
	c.Success(cityId)
}

// @Title 修改城市
// @Summary 修改城市
// @Description 修改城市
// @Param cityid  query  int true "城市id"
// @Param areaid  query  int false "大区id"
// @Param province_name query  string false "省份名称"
// @Param cityname query  string false "城市名称"
// @Param ename query  string false "城市拼音"
// @Param shortname query  int false "城市缩写"
// @Param longitude query  string false "径度"
// @Param latitude query 	 string false "维度"
// @Param city_code query 	 string false "城市编码"
// @Success 200 {object} vo.ApiResponse
// @Failure 500 system error
// @router /update [post]
func (c *CityController) Update() {
	var cityId, city_err = c.GetInt64("cityid")
	province_name := c.GetString("province_name")
	cityname := c.GetString("cityname")
	ename := c.GetString("ename")
	shortname := c.GetString("shortname")
	longitude, _ := c.GetFloat("longitude")
	latitude, _ := c.GetFloat("latitude")
	city_code := c.GetString("city_code")
	if city_err != nil {
		c.Failure(recover.ParamError, "请输入城市ID：cityid", [...]string{})
	}
	cityInfo := make(map[string]interface{})
	cityAllInfo := make(map[string]interface{})
	var areaId, _ = c.GetInt("areaid")
	if areaId > 0 {
		cityInfo["areaid"] = areaId
		cityAllInfo["areaid"] = areaId
	}

	// 省份名称的校验
	if province_name != "" {
		//查找省份信息
		province_name = strings.Replace(province_name, "省", "", -1)
		province_name = strings.Replace(province_name, "市", "", -1)
		cityProvinceService := services.CityProvinceService{}
		provinceInfo := cityProvinceService.FindByProvincename(province_name)

		if provinceInfo != nil {
			cityInfo["provinceid"] = provinceInfo["provinceid"].(int)
			cityAllInfo["provinceid"] = provinceInfo["provinceid"].(int)
			cityInfo["big_areaid"] = provinceInfo["big_areaid"].(int8)
			cityAllInfo["big_areaid"] = provinceInfo["big_areaid"].(int8)
		}
	}

	if cityname != "" {
		cityInfo["cityname"] = cityname
		cityAllInfo["cityname"] = cityname
	}

	if ename != "" {
		cityInfo["ename"] = ename
		cityAllInfo["ename"] = ename
	}

	if shortname != "" {
		cityInfo["shortname"] = shortname
	}

	if longitude > 0 {
		cityInfo["longitude"] = longitude
		cityAllInfo["longitude"] = longitude
	}

	if latitude > 0 {
		cityInfo["latitude"] = latitude
		cityAllInfo["latitude"] = latitude
	}

	if city_code != "" {
		cityAllInfo["city_code"] = city_code
	}

	if cityInfo == nil {
		c.Failure(recover.ParamError, "请输入要修改的值", [...]string{})
	}

	ormer, _ := models.NewCityDao().NewOrmer(&models.City{})

	cityService := services.NewCityService()
	_, error := cityService.UpdateById(cityId, cityInfo)
	ormer.Begin()
	if error != nil {
		ormer.Rollback()
		c.Failure(recover.Failed, error.Error(), "")
		return
	}

	cityAllService := services.CityAllService{}
	_, error = cityAllService.UpdateById(cityId, cityAllInfo)
	if error != nil {
		ormer.Rollback()
		c.Failure(recover.Failed, error.Error(), "")
		return
	}
	ormer.Commit()
	c.Success(true)
}
