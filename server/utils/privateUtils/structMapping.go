package privateUtils

import (
	"fmt"
	"reflect"
)

// MapFields
// @author 林柏宇
// @description 结构体映射方法：将一个结构体的数据映射到另一个结构体中，两个参数一个源结构体，另一个接收结构体
func MapFields(source interface{}, dest interface{}) error {
	sourceValue := reflect.ValueOf(source).Elem()
	destValue := reflect.ValueOf(&dest).Elem()

	for i := 0; i < sourceValue.NumField(); i++ {
		sourceFieldName := sourceValue.Type().Field(i).Name
		destFieldName := destValue.Type().Field(i).Tag.Get("map")

		if destFieldName == "" {
			destFieldName = sourceFieldName
		}

		destFieldValue := destValue.FieldByName(destFieldName)
		if !destFieldValue.IsValid() {
			return fmt.Errorf("无效字段名: %s", destFieldName)
		}

		destFieldType := destFieldValue.Type()
		sourceFieldValue := sourceValue.FieldByName(sourceFieldName)

		if !sourceFieldValue.IsValid() {
			return fmt.Errorf("无效字段名: %s", sourceFieldName)
		}

		if sourceFieldValue.Type().ConvertibleTo(destFieldType) {
			destFieldValue.Set(sourceFieldValue.Convert(destFieldType))
		} else {
			return fmt.Errorf("类型不匹配: %s -> %s", sourceFieldValue.Type(), destFieldType)
		}
	}

	return nil
}
