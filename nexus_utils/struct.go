package nexus_utils

import (
	"errors"
	"reflect"
)

// IsPointer 判断值是否为指针类型，返回反射值
func IsPointer(v interface{}) (reflect.Value, error) {
	if v == nil {
		return reflect.Value{}, errors.New("值不能为nil")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return reflect.Value{}, errors.New("值必须是指针类型")
	}

	if rv.IsNil() {
		return reflect.Value{}, errors.New("指针不能为nil")
	}

	return rv, nil
}

// IsSlice 判断值是否为切片类型，返回反射值
func IsSlice(v interface{}) (reflect.Value, error) {
	if v == nil {
		return reflect.Value{}, errors.New("值不能为nil")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return reflect.Value{}, errors.New("值必须是指针类型")
	}

	if rv.IsNil() {
		return reflect.Value{}, errors.New("指针不能为nil")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Slice {
		return reflect.Value{}, errors.New("值必须是切片类型")
	}

	return rv, nil
}
