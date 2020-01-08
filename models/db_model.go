package models

// 数据库orm基础类
// 所有子类必须实现IBaseModel接口
// 高级查询使用方法参考：https://beego.me/docs/mvc/model/query.md

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
	"time"
)

type IBaseModel interface {
	GetDBName() string
	GetTableName() string
}

// Pagination 数据模型分页对象.
type Pagination struct {
	Page       uint16      `json:"page"`
	PageSize   uint16      `json:"page_size"`
	TotalCount uint64      `json:"total_count"`
	List       interface{} `json:"list"`
}

type IBaseDao interface {
	Insert(md IBaseModel) (int64, error)
	InsertOrUpdate(md IBaseModel, colConflitAndArgs ...string) (int64, error)
	InsertMulti(mds []IBaseModel) (line int64, err error)
	Update(where map[string]interface{}, data map[string]interface{}) (line int64, err error)
	Delete(where map[string]interface{}) (line int64, err error)
	SelectOne(md IBaseModel, where map[string]interface{}, fields ...string) (err error)
	SelectOneOnWrite(md IBaseModel, where map[string]interface{}, fields ...string) (err error)
	SelectList(mds interface{}, where map[string]interface{}, fields []string) (err error)
	SelectAll(mds interface{}, where map[string]interface{}, fields []string, orderBy []string, groupBy []string) (err error)
	SelectPageList(mds interface{}, where map[string]interface{}, fields []string, orderBy []string, groupBy []string, page uint16, pageSize uint16) (pagination *Pagination, err error)
	NewOrmer(model IBaseModel) (orm.Ormer, error)
}

type BaseDao struct {
	EntityType reflect.Type
}

// Insert 插入单条记录.
func (b *BaseDao) Insert(md IBaseModel) (int64, error) {
	types := b.EntityType
	types.Name()
	ormer, err := b.NewOrmer(nil)
	if err != nil {
		return 0, err
	}
	return ormer.Insert(md)
}

// InsertOrUpdate 插入或者更新单条记录.
func (b *BaseDao) InsertOrUpdate(md IBaseModel, colConflitAndArgs ...string) (int64, error) {
	ormer, err := b.NewOrmer(nil)
	if err != nil {
		return 0, err
	}
	return ormer.InsertOrUpdate(md, colConflitAndArgs...)
}

// InsertMulti 批量插入多条记录.
func (b *BaseDao) InsertMulti(mds []IBaseModel) (line int64, err error) {
	ormer, err := b.NewOrmer(nil)
	if err != nil {
		return 0, err
	}
	return ormer.InsertMulti(len(mds), mds)
}

// Update 依据当前查询条件，进行批量更新操作.
func (b *BaseDao) Update(where map[string]interface{}, data map[string]interface{}) (line int64, err error) {
	query, err := b.NewQuerySeter(nil)
	if err != nil {
		return 0, err
	}
	return b.queryFilter(query, where).Update(data)
}

// Delete 依据当前条件，进行删除操作.
func (b *BaseDao) Delete(where map[string]interface{}) (line int64, err error) {
	query, err := b.NewQuerySeter(nil)
	if err != nil {
		return 0, err
	}
	return b.queryFilter(query, where).Delete()
}

// SelectOne 从读库获取单条记录.
func (b *BaseDao) SelectOne(md IBaseModel, where map[string]interface{}, fields ...string) (err error) {
	return b.queryOne(md, where, fields, false)
}

// SelectOneOnWrite 从写库获取单条记录.
func (b *BaseDao) SelectOneOnWrite(md IBaseModel, where map[string]interface{}, fields ...string) (err error) {
	return b.queryOne(md, where, fields, true)
}

// SelectList 根据过滤条件返回对应的结果集对象.
func (b *BaseDao) SelectList(mds interface{}, where map[string]interface{}, fields []string) (err error) {
	return b.SelectAll(mds, where, fields, nil, nil)
}

// SelectAll 根据过滤、排序、分组条件返回对应的结果集对象.
func (b *BaseDao) SelectAll(mds interface{}, where map[string]interface{}, fields []string, orderBy []string,
	groupBy []string) (err error) {
	query, err := b.NewQuerySeter(nil)
	if err != nil {
		return err
	}

	query = b.queryFilter(query, where)

	if len(orderBy) > 0 {
		query = query.OrderBy(orderBy...)
	}
	if len(groupBy) > 0 {
		query = query.GroupBy(groupBy...)
	}

	if _, err := query.All(mds, fields...); err != nil && err != orm.ErrNoRows {
		logs.Error("数据库返回对应的结果集对象异常， 异常信息 => ", err.Error())
	}

	return err
}

// SelectPageList 根据条件获取结果集分页对象.
func (b *BaseDao) SelectPageList(mds interface{}, where map[string]interface{}, fields []string, orderBy []string,
	groupBy []string, page uint16, pageSize uint16) (pagination *Pagination, err error) {
	offset := (page - 1) * pageSize
	pagination = &Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: 0,
		List:       [...]string{},
	}

	query, err := b.NewQuerySeter(nil)
	if err != nil {
		return pagination, err
	}

	query = b.queryFilter(query, where)

	if len(orderBy) > 0 {
		query = query.OrderBy(orderBy...)
	}
	if len(groupBy) > 0 {
		query = query.GroupBy(groupBy...)
	}
	if page > 0 && pageSize > 0 {
		query = query.Limit(pageSize, offset)
	}

	if _, err := query.All(mds, fields...); err != nil && err != orm.ErrNoRows {
		logs.Error("数据库返回对应的结果集对象异常， 异常信息 => ", err.Error())
		return pagination, err
	}
	pagination.List = mds

	count, err := query.Count()
	if err != nil {
		logs.Error("数据库获取结果行数异常 ->", err)
	}
	pagination.TotalCount = uint64(count)

	return pagination, err
}

// queryOne 根据过滤条件、查询列、是否主库查询，获取单条记录.
func (b *BaseDao) queryOne(container interface{}, where map[string]interface{}, fields []string, useWrite bool) (err error) {
	model := b.convertToInterface()
	ormer, err := b.NewOrmer(model)
	if err != nil {
		return err
	}
	query := ormer.QueryTable(model.GetTableName())

	query = b.queryFilter(query, where)

	if !useWrite {
		if err = query.One(container, fields...); err != nil {
			if err == orm.ErrNoRows {
				return err
			} else {
				logs.Error("数据库单条查询异常， 异常信息 => ", err.Error())
			}
		}
		return nil
	}

	if err := ormer.Begin(); err != nil {
		logs.Error(err.Error())
		return err
	}

	if err = query.One(container, fields...); err != nil {
		if err == orm.ErrNoRows {
			return err
		} else {
			logs.Error("数据库单条查询异常， 异常信息 => ", err.Error())
		}
		if rollBackErr := ormer.Rollback(); rollBackErr != nil {
			logs.Error("数据库事务回滚异常， 异常信息 => ", rollBackErr.Error())
		}
		return nil
	}

	if commitError := ormer.Commit(); commitError != nil {
		logs.Error("数据库事务提交异常， 异常信息 => ", commitError.Error())
	}

	return err
}

// reflectChildModel 获取当前结构体对象.
func (b *BaseDao) reflectChildModel() interface{} {
	return reflect.New(b.EntityType).Interface()
}

// convertToInterface 获取IBaseModel接口实现类.
func (b *BaseDao) convertToInterface() IBaseModel {
	if model, ok := b.reflectChildModel().(IBaseModel); ok {
		return model
	}

	panic("Model未实现IBaseModel接口")
}

// queryFilter querySeter条件过滤.
func (b *BaseDao) queryFilter(query orm.QuerySeter, where map[string]interface{}) orm.QuerySeter {
	for k, v := range where {
		query = query.Filter(k, v)
	}
	return query
}

// 创建querySeter对象
func (b *BaseDao) NewQuerySeter(model IBaseModel) (orm.QuerySeter, error) {
	if model == nil {
		model = b.convertToInterface()
	}
	ormer := orm.NewOrm()
	name := model.GetDBName()
	if name != "default" {
		if err := ormer.Using(name); err != nil {
			logs.Error("指定数据库异常， 异常信息 => ", err.Error())
			return nil, err
		}
	}
	return ormer.QueryTable(model.GetTableName()), nil
}

// NewOrmer 创建ormer对象.
func (b *BaseDao) NewOrmer(model IBaseModel) (orm.Ormer, error) {
	if model == nil {
		model = b.convertToInterface()
	}
	ormer := orm.NewOrm()
	name := model.GetDBName()
	if name != "default" {
		if err := ormer.Using(name); err != nil {
			logs.Error("指定数据库异常， 异常信息 => ", err.Error())
			return nil, err
		}
	}
	return ormer, nil
}

// ResultHandle 列表结果集处理.
func ResultHandle(result []interface{}, fields []string,
	callback ...func(remark string, dataValue interface{}) interface{}) []interface{} {
	var ml []interface{}
	if len(result) == 0 || len(callback) > 1 {
		return ml
	}

	if len(fields) == 0 {
		for _, v := range result {
			ml = append(ml, v)
		}
		return ml
	}

	for _, v := range result {
		m := make(map[string]interface{})
		val := reflect.ValueOf(v)
		t := reflect.TypeOf(v)
		for _, fName := range fields {
			field, _ := t.FieldByName(fName)
			m = fieldFormat(m, fName, field.Tag.Get("json"), val.FieldByName(fName).Interface(), callback...)
		}
		ml = append(ml, m)
	}

	return ml
}

// OneResultHandle 单个结果集处理.
func OneResultHandle(result interface{}, fields []string,
	callback ...func(remark string, dataValue interface{}) interface{}) map[string]interface{} {
	var ml = make(map[string]interface{})
	if result == nil || len(callback) > 1 {
		return ml
	}

	val := reflect.ValueOf(result)
	t := reflect.TypeOf(result)
	for _, fName := range fields {
		field, _ := t.FieldByName(fName)
		ml = fieldFormat(ml, fName, field.Tag.Get("json"), val.FieldByName(fName).Interface(), callback...)
	}

	return ml
}

// fieldFormat 结果集列格式化处理.
func fieldFormat(m map[string]interface{}, fName string, remark string,
	dataValue interface{}, callback ...func(remark string, dataValue interface{}) interface{}) map[string]interface{} {
	switch dataValue.(type) {
	case time.Time:
		dataValue = dataValue.(time.Time).Format("2006-01-02 15:04:05")
	case float64:
		dataValue, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", dataValue), 64)
	}

	if remark == "" {
		m[fName] = dataValue
	} else {
		if len(callback) > 0 {
			dataValue = callback[0](remark, dataValue)
		}
		m[remark] = dataValue
	}

	return m
}
