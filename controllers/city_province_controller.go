package controllers

// @TagName CityProvince
// @Description CityProvinceController operations for CityProvince
type CityProvinceController struct {
	BaseController
}

// URLMapping ...
func (c *CityProvinceController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// @Title Index
// @Summary Index
// @Description Index function
// @Param debug query bool false "debug"
// @Success 200 {object} vo.ApiResponse
// @Failure 500 system error
// @router /index [get]
func (c *CityProvinceController) Index() {
	c.SuccessWithoutData()
}
