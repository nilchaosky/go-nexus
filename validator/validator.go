package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatFieldErrors 格式化字段校验错误
// req 为校验的结构体，err 为validator返回的错误
// 返回格式化后的错误信息字符串
func FormatFieldErrors(req interface{}, err error) string {
	if err == nil {
		return ""
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return err.Error()
	}

	// 预分配切片容量，减少内存重新分配
	errorCount := len(validationErrors)
	errorMessages := make([]string, 0, errorCount)
	for _, fieldError := range validationErrors {
		fieldName := fieldError.Field()
		namespace := fieldError.Namespace()
		tag := fieldError.Tag()
		param := fieldError.Param()

		// 解析字段名，处理列表索引
		// 使用 Namespace 获取完整路径（包含索引信息）
		displayName := formatFieldName(req, namespace, fieldName)

		// 获取标签模板
		template := GetTagTemplate(tag)

		// 格式化错误消息
		message := formatMessage(template, displayName, param, fieldError)
		errorMessages = append(errorMessages, message)
	}

	if len(errorMessages) == 0 {
		return ""
	}

	return strings.Join(errorMessages, ";")
}

// formatMessage 格式化错误消息
// template 为模板字符串，fieldName 为字段名，param 为参数，fieldError 为字段错误
func formatMessage(template, fieldName, param string, fieldError validator.FieldError) string {
	// 如果模板不包含占位符，直接返回
	if !strings.Contains(template, "{0}") && !strings.Contains(template, "{1}") {
		return template
	}

	message := template

	// 替换 {0} 为字段名
	if strings.Contains(message, "{0}") {
		message = strings.ReplaceAll(message, "{0}", fieldName)
	}

	// 替换 {1} 为参数值
	if strings.Contains(message, "{1}") {
		if param != "" {
			message = strings.ReplaceAll(message, "{1}", param)
		} else {
			// 如果没有参数，尝试从错误中获取其他信息
			if fieldError.Value() != nil {
				message = strings.ReplaceAll(message, "{1}", fmt.Sprintf("%v", fieldError.Value()))
			}
		}
	}

	return message
}

// formatFieldName 格式化字段名，支持列表索引和label标签
// req 为结构体实例，namespace 为完整命名空间（如 Request.Users[0].Name），fieldName 为字段名（如 Name）
// 返回格式化后的字段显示名称
func formatFieldName(req interface{}, namespace, fieldName string) string {
	if req == nil {
		return fieldName
	}

	// 使用 namespace 解析完整路径（包含索引信息）
	// 例如：Request.Users[0].Name -> Request, Users[0], Name
	parts := strings.Split(namespace, ".")
	if len(parts) == 0 {
		// 如果没有命名空间，使用字段名
		displayName := getFieldLabel(req, fieldName)
		if displayName == "" {
			return fieldName
		}
		return displayName
	}

	// 预分配切片容量，减少内存重新分配
	resultParts := make([]string, 0, len(parts)-1)
	currentReq := req

	// 跳过第一个部分（通常是结构体类型名）
	for i := 1; i < len(parts); i++ {
		part := parts[i]

		// 检查是否包含索引 [0]
		leftBracket := strings.IndexByte(part, '[')
		if leftBracket >= 0 {
			rightBracket := strings.IndexByte(part, ']')
			if rightBracket > leftBracket {
				// 提取字段名和索引
				fieldPart := part[:leftBracket]
				indexPart := part[leftBracket : rightBracket+1]

				// 获取字段的label或使用字段名
				displayName := getFieldLabel(currentReq, fieldPart)
				if displayName == "" {
					displayName = fieldPart
				}

				// 将索引转换为中文（[0] -> 第一项，[1] -> 第二项）
				indexText := formatIndex(indexPart)

				// 组合显示名称和索引
				resultParts = append(resultParts, displayName+indexText)

				// 获取嵌套字段的值（用于后续处理）
				currentReq = getNestedFieldValue(currentReq, fieldPart, indexPart)
			} else {
				// 格式错误，使用原始part
				displayName := getFieldLabel(currentReq, part)
				if displayName == "" {
					displayName = part
				}
				resultParts = append(resultParts, displayName)
				currentReq = getNestedFieldValue(currentReq, part, "")
			}
		} else {
			// 普通字段，获取label或使用字段名
			displayName := getFieldLabel(currentReq, part)
			if displayName == "" {
				displayName = part
			}
			resultParts = append(resultParts, displayName)

			// 获取嵌套字段的值
			currentReq = getNestedFieldValue(currentReq, part, "")
		}
	}

	if len(resultParts) == 0 {
		// 如果没有解析到任何部分，使用字段名
		displayName := getFieldLabel(req, fieldName)
		if displayName == "" {
			return fieldName
		}
		return displayName
	}

	return strings.Join(resultParts, ".")
}

// formatIndex 将索引转换为中文表示
// index 为索引字符串，如 "[0]", "[1]"
// 返回中文表示，如 "第一项", "第二项"（索引从0开始，显示时+1）
func formatIndex(index string) string {
	// 提取索引数字（去掉首尾的 [ 和 ]）
	if len(index) < 3 || index[0] != '[' || index[len(index)-1] != ']' {
		return index
	}
	indexStr := index[1 : len(index)-1]
	if indexStr == "" {
		return index
	}

	// 使用 strconv 转换为数字，性能优于 fmt.Sscanf
	indexNum, err := strconv.Atoi(indexStr)
	if err != nil {
		return index
	}

	// 索引+1（因为索引从0开始，但显示时从1开始）
	displayNum := indexNum + 1

	// 转换为中文
	chineseNumbers := []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	if displayNum > 0 && displayNum <= len(chineseNumbers) {
		return "第" + chineseNumbers[displayNum-1] + "项"
	}

	// 如果超过10，使用数字（使用 strconv 性能更好）
	return "第" + strconv.Itoa(displayNum) + "项"
}

// getFieldLabel 从结构体字段获取label标签
// req 为结构体实例，fieldName 为字段名
// 返回label标签的值，如果没有则返回空字符串
func getFieldLabel(req interface{}, fieldName string) string {
	if req == nil {
		return ""
	}

	// 获取结构体类型
	rt := reflect.TypeOf(req)
	if rt == nil {
		return ""
	}

	// 如果是指针类型，获取元素类型
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	// 必须是结构体类型
	if rt.Kind() != reflect.Struct {
		return ""
	}

	// 查找字段
	field, found := rt.FieldByName(fieldName)
	if !found {
		return ""
	}

	// 获取label标签
	label := field.Tag.Get("label")
	return label
}

// getNestedFieldValue 获取嵌套字段的值
// req 为结构体实例，fieldName 为字段名，index 为索引（如 "[0]"）
func getNestedFieldValue(req interface{}, fieldName, index string) interface{} {
	if req == nil {
		return nil
	}

	rv := reflect.ValueOf(req)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return nil
	}

	field := rv.FieldByName(fieldName)
	if !field.IsValid() {
		return nil
	}

	// 如果有索引，获取切片/数组元素
	if index != "" {
		if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
			// 返回切片元素类型的新实例（用于获取label）
			if field.Len() > 0 {
				return field.Index(0).Interface()
			}
			// 如果切片为空，尝试从类型创建零值
			elemType := field.Type().Elem()
			if elemType.Kind() == reflect.Ptr {
				elemType = elemType.Elem()
			}
			return reflect.New(elemType).Interface()
		}
	}

	return field.Interface()
}
