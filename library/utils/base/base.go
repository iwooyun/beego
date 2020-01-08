package base

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strconv"
)

//	Contains 搜索数组中是否存在指定的值.
func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

//	StructToMap struct转map.
func StructToMap(object interface{}) map[string]interface{} {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)

	var data = make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Name] = objValue.Field(i).Interface()
	}
	return data
}

//	MapMerge 合并map.
func MapMerge(map1 map[string]interface{}, map2 map[string]interface{}) map[string]interface{} {
	if len(map1) == 0 {
		return map2
	}
	if len(map2) == 0 {
		return map1
	}

	for k, v := range map2 {
		map1[k] = v
	}

	return map1
}

// DebugLogsFormat json格式化输出.
func DebugLogsFormat(v interface{}) {
	msg, _ := json.MarshalIndent(v, "", "\t")
	logs.Informational("debug message Content = > %s", msg)
}

// SliceValues 获取指定下标的字段值的map.
func SliceValues(slice []interface{}, fieldName string) []interface{} {
	var valueData []interface{}
	for _, val := range slice {
		valueData = append(valueData, val.(map[string]interface{})[fieldName])
	}
	return valueData
}

// SliceColumnData 指定字段作为map下标   下标类型为string
func SliceColumnData(m []interface{}, fieldName string, result map[string]interface{}) {
	for _, value := range m {
		_index := value.(map[string]interface{})[fieldName].(string)
		result[_index] = value
	}
}

// SliceIntColumnData 指定字段作为map下标  下标类型为int
func SliceIntColumnData(m []interface{}, fieldName string, result map[int]interface{}) {
	for _, value := range m {
		_index := value.(map[string]interface{})[fieldName].(int)
		result[_index] = value
	}
}

// IsNil 判断interface是否为空
func IsNil(i interface{}) bool {
	defer func() bool {
		recover()
		return false
	}()

	vi := reflect.ValueOf(i)

	if !vi.IsValid() || vi.IsNil() {
		return true
	}

	return false
}

func Int64ToInt(input int64) int {
	string := strconv.FormatInt(input, 10)
	output, _ := strconv.Atoi(string)
	return output
}
