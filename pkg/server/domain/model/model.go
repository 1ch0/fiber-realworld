package model

import (
	"fmt"
	"reflect"
	"time"
)

var tableNamePrefix = "fiber_"

var registeredModels = map[string]Interface{}

// Interface model interface
type Interface interface {
	TableName() string
	ShortTableName() string
}

// RegisterModel register model
func RegisterModel(models ...Interface) {
	for _, model := range models {
		if _, exist := registeredModels[model.TableName()]; exist {
			panic(fmt.Errorf("model table name %s conflict", model.TableName()))
		}
		registeredModels[model.TableName()] = model
	}
}

// GetRegisterModels will return the register models
func GetRegisterModels() map[string]Interface {
	return registeredModels
}

// BaseModel common model
type BaseModel struct {
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

// SetCreateTime set create time
func (m *BaseModel) SetCreateTime(time time.Time) {
	m.CreateTime = time
}

// SetUpdateTime set update time
func (m *BaseModel) SetUpdateTime(time time.Time) {
	m.UpdateTime = time
}

func deepCopy(src interface{}) interface{} {
	dst := reflect.New(reflect.TypeOf(src).Elem())

	val := reflect.ValueOf(src).Elem()
	nVal := dst.Elem()
	for i := 0; i < val.NumField(); i++ {
		nvField := nVal.Field(i)
		nvField.Set(val.Field(i))
	}

	return dst.Interface()
}
