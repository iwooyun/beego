// @Time : 2019/9/16 14:10
// @Author : duanqiangwen
// @File : base_validator
// @Software: GoLand
package validator

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"reflect"
	"strconv"
)

type IValidator interface {
	ErrorHandle(valid validation.Validation)
}

type Instance func() IValidator

var (
	adapters                  = make(map[string]Instance)
	GlobalControllerValidator = make(map[string]ValidComments)
)

type ValidComments struct {
	Validator string
	Method    string
}

type Validator struct {
	Input *context.BeegoInput
}

func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("validator: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("validator: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func NewValidator(validComments ValidComments, input *context.BeegoInput) (err error) {
	defer func() {
		errStr := recover()
		if errStr != nil {
			err = errors.New(fmt.Sprintf("%s", errStr))
		}
	}()

	instanceFunc, ok := adapters[validComments.Validator]
	if !ok {
		err := fmt.Errorf("validator: unknown adapter name %q (forgot to import?)", validComments.Validator)
		logs.Error(err)
		panic(err)
	}
	adapter := instanceFunc()
	v := reflect.ValueOf(adapter)
	args := []reflect.Value{reflect.ValueOf(input)}
	v.MethodByName(validComments.Method).Call(args)

	return
}

// GetString returns the input value by key string or the default value while it's present and input is blank
func (c *Validator) GetString(key string, def ...string) string {
	if v := c.Input.Query(key); v != "" {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetInt returns input as an int or the default value while it's present and input is blank
func (c *Validator) GetInt(key string, def ...int) (int, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.Atoi(strv)
}

// GetInt8 return input as an int8 or the default value while it's present and input is blank
func (c *Validator) GetInt8(key string, def ...int8) (int8, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	i64, err := strconv.ParseInt(strv, 10, 8)
	return int8(i64), err
}

// GetUint8 return input as an uint8 or the default value while it's present and input is blank
func (c *Validator) GetUint8(key string, def ...uint8) (uint8, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	u64, err := strconv.ParseUint(strv, 10, 8)
	return uint8(u64), err
}

// GetInt16 returns input as an int16 or the default value while it's present and input is blank
func (c *Validator) GetInt16(key string, def ...int16) (int16, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	i64, err := strconv.ParseInt(strv, 10, 16)
	return int16(i64), err
}

// GetUint16 returns input as an uint16 or the default value while it's present and input is blank
func (c *Validator) GetUint16(key string, def ...uint16) (uint16, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	u64, err := strconv.ParseUint(strv, 10, 16)
	return uint16(u64), err
}

// GetInt32 returns input as an int32 or the default value while it's present and input is blank
func (c *Validator) GetInt32(key string, def ...int32) (int32, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	i64, err := strconv.ParseInt(strv, 10, 32)
	return int32(i64), err
}

// GetUint32 returns input as an uint32 or the default value while it's present and input is blank
func (c *Validator) GetUint32(key string, def ...uint32) (uint32, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	u64, err := strconv.ParseUint(strv, 10, 32)
	return uint32(u64), err
}

// GetInt64 returns input value as int64 or the default value while it's present and input is blank.
func (c *Validator) GetInt64(key string, def ...int64) (int64, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseInt(strv, 10, 64)
}

// GetUint64 returns input value as uint64 or the default value while it's present and input is blank.
func (c *Validator) GetUint64(key string, def ...uint64) (uint64, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseUint(strv, 10, 64)
}

// GetBool returns input value as bool or the default value while it's present and input is blank.
func (c *Validator) GetBool(key string, def ...bool) (bool, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseBool(strv)
}

// GetFloat returns input value as float64 or the default value while it's present and input is blank.
func (c *Validator) GetFloat(key string, def ...float64) (float64, error) {
	strv := c.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseFloat(strv, 64)
}

func (c *Validator) ErrorHandle(valid validation.Validation) {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			panic(fmt.Errorf(":%s %s", err.Key, err.Message))
		}
	}
}
