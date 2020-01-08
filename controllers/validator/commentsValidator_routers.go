package validator

const (
	CityAllValidator = "CityAllValidator"
	TestValidator    = "TestValidator"
	CityValidator    = "CityValidator"
)

func init() {
	GlobalControllerValidator["/api/city-base/get-profile-by-ename"] = ValidComments{
		Validator: CityAllValidator,
		Method:    "GetProfileByEname",
	}
	GlobalControllerValidator["/api/city-base/get-list-by-provinceid"] = ValidComments{
		Validator: CityAllValidator,
		Method:    "GetListByProvinceId",
	}
	GlobalControllerValidator["/api/test/log"] = ValidComments{
		Validator: TestValidator,
		Method:    "Log",
	}
	GlobalControllerValidator["/api/city-base/search-list-by-ename"] = ValidComments{
		Validator: CityAllValidator,
		Method:    "SearchListByEname",
	}
	GlobalControllerValidator["/api/city/get-list-by-area-id"] = ValidComments{
		Validator: CityValidator,
		Method:    "GetListByAreaId",
	}
	GlobalControllerValidator["/api/city/get-list-by-province-id"] = ValidComments{
		Validator: CityValidator,
		Method:    "GetListByProvinceId",
	}
	GlobalControllerValidator["/api/city/get-list-by-ename"] = ValidComments{
		Validator: CityValidator,
		Method:    "GetListByEname",
	}
	GlobalControllerValidator["/api/city/add"] = ValidComments{
		Validator: CityValidator,
		Method:    "Add",
	}
	GlobalControllerValidator["/api/test/mysql"] = ValidComments{
		Validator: TestValidator,
		Method:    "Mysql",
	}
	GlobalControllerValidator["/api/test/redis"] = ValidComments{
		Validator: TestValidator,
		Method:    "Redis",
	}
	GlobalControllerValidator["/api/test/alarm"] = ValidComments{
		Validator: TestValidator,
		Method:    "Alarm",
	}
	GlobalControllerValidator["/api/city-base/get-profile-by-code"] = ValidComments{
		Validator: CityAllValidator,
		Method:    "GetProfileByCode",
	}
	GlobalControllerValidator["/api/city-base/get-list-by-bigareaid"] = ValidComments{
		Validator: CityAllValidator,
		Method:    "GetListByBigAreaId",
	}
	GlobalControllerValidator["/api/city/get-profile-by-ename"] = ValidComments{
		Validator: CityValidator,
		Method:    "GetProfileByEname",
	}
	GlobalControllerValidator["/api/test/mail"] = ValidComments{
		Validator: TestValidator,
		Method:    "Mail",
	}
	GlobalControllerValidator["/api/city-base/get-profile"] = ValidComments{
		Validator: CityAllValidator,
		Method:    "GetProfile",
	}
	GlobalControllerValidator["/api/city/get-profile-by-id"] = ValidComments{
		Validator: CityValidator,
		Method:    "GetProfileById",
	}
	GlobalControllerValidator["/api/city/get-profile-by-name"] = ValidComments{
		Validator: CityValidator,
		Method:    "GetProfileByName",
	}
	GlobalControllerValidator["/api/city/update"] = ValidComments{
		Validator: CityValidator,
		Method:    "Update",
	}
	GlobalControllerValidator["/api/city/get-all-page-list"] = ValidComments{
		Validator: CityValidator,
		Method:    "GetAllPageList",
	}
}
