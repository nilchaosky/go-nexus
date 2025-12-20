package response

const (
	// CodeSuccess 成功状态码
	CodeSuccess = 0
	// CodeError 默认错误状态码
	CodeError = -1
	// MessageSuccess 成功消息
	MessageSuccess = "success"
)

// Response 统一响应结构体
// T 为泛型类型，表示响应数据的类型
type Response[T any] struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 响应消息
	Data    T      `json:"data"`    // 响应数据，使用泛型
}

// Page 分页结构体
// T 为泛型类型，表示记录项的类型
type Page[T any] struct {
	Current int `json:"current"` // 当前页码
	Size    int `json:"size"`    // 每页大小
	Total   int `json:"total"`   // 总记录数
	Records []T `json:"records"` // 记录列表，使用泛型切片
}

// Success 创建成功响应
// data 为响应数据，使用泛型类型
// 支持值类型和指针类型，例如：
//   - Success(user) 返回 Response[User]
//   - Success(&user) 返回 Response[*User]
func Success[T any](data T) Response[T] {
	return Response[T]{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    data,
	}
}

// SuccessWithNil 创建成功响应，data为nil
// 返回Response[any]类型，适用于不需要类型约束的场景
func SuccessWithNil() Response[any] {
	return Response[any]{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    nil,
	}
}

// SuccessWithNilData 创建成功响应，data为nil指针
// 返回Response[*T]类型，提供类型安全的nil响应
// 适用于需要明确数据类型但数据为空的场景
func SuccessWithNilData[T any]() Response[*T] {
	return Response[*T]{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    nil,
	}
}

// SuccessPage 创建分页成功响应
// data 为分页数据，使用泛型类型
func SuccessPage[T any](data Page[T]) Response[Page[T]] {
	return Response[Page[T]]{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    data,
	}
}

// Error 创建错误响应
// message 为错误消息，data为nil
// 返回Response[any]类型，适用于不需要类型约束的场景
func Error(message string) Response[any] {
	return Response[any]{
		Code:    CodeError,
		Message: message,
		Data:    nil,
	}
}

// ErrorWithCode 创建自定义错误码的错误响应
// code 为自定义错误码，message 为错误消息，data为nil
// 返回Response[any]类型，适用于不需要类型约束的场景
func ErrorWithCode(code int, message string) Response[any] {
	return Response[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// ErrorWithNilData 创建错误响应，data为nil指针
// message 为错误消息，返回Response[*T]类型，提供类型安全的nil响应
func ErrorWithNilData[T any](message string) Response[*T] {
	return Response[*T]{
		Code:    CodeError,
		Message: message,
		Data:    nil,
	}
}

// ErrorWithCodeAndNilData 创建自定义错误码的错误响应，data为nil指针
// code 为自定义错误码，message 为错误消息，返回Response[*T]类型，提供类型安全的nil响应
func ErrorWithCodeAndNilData[T any](code int, message string) Response[*T] {
	return Response[*T]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
