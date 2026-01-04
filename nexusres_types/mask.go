package nexusres_types

import (
	"reflect"
)

// maskFields 检测结构体字段的 nexusmask 标签，将有该标签的字段设置为零值
// 功能说明：
// 1. 递归遍历结构体的所有字段
// 2. 如果字段有 nexusmask 标签，将其设置为零值
// 3. 如果字段没有 nexusmask 标签，递归处理嵌套的结构体、指针、切片/数组中的结构体
// 注意：可以配合 JSON 标签的 omitempty 使用，将字段设置为零值后，omitempty 会使其在 JSON 中不出现
//      例如：Password string `json:"password,omitempty" nexusmask` 或 Password *string `json:"password,omitempty" nexusmask`
func maskFields(data interface{}) {
	if data == nil {
		return
	}

	rv := reflect.ValueOf(data)
	// 如果是指针，获取指向的值
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return
		}
		rv = rv.Elem()
	}

	// 必须是结构体类型
	if rv.Kind() != reflect.Struct {
		return
	}

	rt := rv.Type()
	// 遍历所有字段
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldValue := rv.Field(i)

		// 检查是否有 nexusmask 标签
		if _, hasMask := field.Tag.Lookup("nexusmask"); hasMask {
			// 设置为零值
			if fieldValue.CanSet() {
				zeroValue := reflect.Zero(fieldValue.Type())
				fieldValue.Set(zeroValue)
			}
			continue
		}

		// 递归处理嵌套结构体
		maskNestedValue(fieldValue)
	}
}

// maskNestedValue 递归处理嵌套的值（结构体、指针、切片/数组中的结构体）
func maskNestedValue(fieldValue reflect.Value) {
	switch fieldValue.Kind() {
	case reflect.Ptr:
		// 指针类型：如果不是 nil 且指向结构体，递归处理
		if !fieldValue.IsNil() && fieldValue.Elem().Kind() == reflect.Struct {
			maskFields(fieldValue.Interface())
		}
	case reflect.Struct:
		// 值类型结构体：需要获取可寻址的值
		if fieldValue.CanAddr() {
			maskFields(fieldValue.Addr().Interface())
		}
	case reflect.Slice, reflect.Array:
		// 切片/数组：遍历每个元素，递归处理结构体
		for j := 0; j < fieldValue.Len(); j++ {
			elem := fieldValue.Index(j)
			maskNestedValue(elem)
		}
	}
}
