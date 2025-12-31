package nexusres_types

const (
	// CodeSuccess 成功状态码
	CodeSuccess = 0
	// CodeError 默认错误状态码
	CodeError = -1
	// MessageSuccess 成功消息
	MessageSuccess = "success"
)

// Response 统一响应结构体
type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

// Page 分页结构体
type Page[T any] struct {
	Current int   `json:"current"`
	Size    int   `json:"size"`
	Total   int64 `json:"total"`
	Records []*T  `json:"records"`
}

// Success 创建成功响应
func Success[T any](data *T) Response[T] {
	return Response[T]{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    data,
	}
}

// SuccessWithNil 创建成功响应
func SuccessWithNil() Response[interface{}] {
	return Response[interface{}]{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    nil,
	}
}

// SuccessPage 创建分页成功响应
func SuccessPage[T any](data *Page[T]) Response[Page[T]] {
	return Response[Page[T]]{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    data,
	}
}

// Error 创建错误响应
func Error(message string) Response[interface{}] {
	return Response[interface{}]{
		Code:    CodeError,
		Message: message,
		Data:    nil,
	}
}

// ErrorWithCode 创建自定义错误码的错误响应
func ErrorWithCode(code int, message string) Response[interface{}] {
	return Response[interface{}]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
