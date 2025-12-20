package validator

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

// TestFormatFieldErrors_Required 测试字段是否为空
func TestFormatFieldErrors_Required(t *testing.T) {
	// 定义测试结构体
	type User struct {
		Name     string `validate:"required" label:"姓名"`
		Password string `validate:"required" label:"密码"`
	}

	validate := validator.New()

	// 测试字段为空的情况
	user := User{
		Name: "",
	}

	err := validate.Struct(user)
	if err == nil {
		t.Error("期望校验失败，但校验通过了")
		return
	}

	// 格式化错误信息
	errorMsg := FormatFieldErrors(user, err)
	if errorMsg == "" {
		t.Error("期望返回错误信息，但返回为空")
		return
	}

	// 输出格式化后的错误信息
	t.Logf("格式化后的错误信息: %s", errorMsg)

	// 验证错误信息包含必填项提示
	if !strings.Contains(errorMsg, "必填") && !strings.Contains(errorMsg, "required") {
		t.Errorf("错误信息应该包含必填提示，实际: %s", errorMsg)
	}

	// 验证错误信息包含label标签的值（姓名）而不是字段名（Name）
	if !strings.Contains(errorMsg, "姓名") {
		t.Errorf("错误信息应该包含label标签的值'姓名'，实际: %s", errorMsg)
	}

	// 验证错误信息不包含字段名（因为使用了label）
	if strings.Contains(errorMsg, "Name") {
		t.Errorf("错误信息不应该包含字段名Name（应该使用label），实际: %s", errorMsg)
	}
}

// TestFormatFieldErrors_List 测试列表校验功能
func TestFormatFieldErrors_List(t *testing.T) {
	// 定义测试结构体
	type User struct {
		Name  string `validate:"required" label:"姓名"`
		Email string `validate:"required,email" label:"邮箱"`
	}

	type Request struct {
		Users []User `validate:"required,dive" label:"用户列表"`
	}

	validate := validator.New()

	// 测试列表为空的情况
	t.Run("列表为空", func(t *testing.T) {
		req := Request{
			Users: nil,
		}

		err := validate.Struct(req)
		if err == nil {
			t.Error("期望校验失败，但校验通过了")
			return
		}

		errorMsg := FormatFieldErrors(req, err)
		if errorMsg == "" {
			t.Error("期望返回错误信息，但返回为空")
			return
		}

		t.Logf("格式化后的错误信息: %s", errorMsg)

		// 验证错误信息包含列表字段
		if !strings.Contains(errorMsg, "用户列表") && !strings.Contains(errorMsg, "Users") {
			t.Errorf("错误信息应该包含列表字段，实际: %s", errorMsg)
		}
	})

	// 测试列表元素校验
	t.Run("列表元素校验", func(t *testing.T) {
		req := Request{
			Users: []User{
				{
					Name:  "",
					Email: "",
				},
				{
					Name:  "张三",
					Email: "invalid-email",
				},
			},
		}

		err := validate.Struct(req)
		if err == nil {
			t.Error("期望校验失败，但校验通过了")
			return
		}

		errorMsg := FormatFieldErrors(req, err)
		if errorMsg == "" {
			t.Error("期望返回错误信息，但返回为空")
			return
		}

		t.Logf("格式化后的错误信息: %s", errorMsg)

		// 验证错误信息包含"第一项"而不是[0]
		if !strings.Contains(errorMsg, "第一项") {
			t.Errorf("错误信息应该包含'第一项'，实际: %s", errorMsg)
		}

		// 验证错误信息包含"第二项"
		if !strings.Contains(errorMsg, "第二项") {
			t.Errorf("错误信息应该包含'第二项'，实际: %s", errorMsg)
		}

		// 验证错误信息包含label（姓名、邮箱）
		if !strings.Contains(errorMsg, "姓名") && !strings.Contains(errorMsg, "邮箱") {
			t.Errorf("错误信息应该包含label（姓名或邮箱），实际: %s", errorMsg)
		}
	})

	// 测试列表元素部分字段为空
	t.Run("列表元素部分字段为空", func(t *testing.T) {
		req := Request{
			Users: []User{
				{
					Name:  "李四",
					Email: "",
				},
			},
		}

		err := validate.Struct(req)
		if err == nil {
			t.Error("期望校验失败，但校验通过了")
			return
		}

		errorMsg := FormatFieldErrors(req, err)
		if errorMsg == "" {
			t.Error("期望返回错误信息，但返回为空")
			return
		}

		t.Logf("格式化后的错误信息: %s", errorMsg)

		// 验证错误信息包含"第一项"而不是[0]
		if !strings.Contains(errorMsg, "第一项") {
			t.Errorf("错误信息应该包含'第一项'，实际: %s", errorMsg)
		}

		// 验证错误信息包含邮箱字段的错误
		if !strings.Contains(errorMsg, "邮箱") {
			t.Errorf("错误信息应该包含邮箱字段，实际: %s", errorMsg)
		}
	})
}
