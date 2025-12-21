package jsoniter

import (
	"reflect"
	"testing"
)

// TestSerializer_Marshal 测试序列化功能
func TestSerializer_Marshal(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	serializer := NewSerializer()

	user := User{
		Name:  "张三",
		Age:   25,
		Email: "zhangsan@example.com",
	}

	data, err := serializer.Marshal(user)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	if len(data) == 0 {
		t.Error("序列化后的数据不能为空")
	}

	t.Logf("序列化结果: %s", string(data))
}

// TestSerializer_Unmarshal 测试反序列化功能
func TestSerializer_Unmarshal(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	serializer := NewSerializer()

	jsonData := `{"name":"李四","age":30,"email":"lisi@example.com"}`

	var user User
	err := serializer.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if user.Name != "李四" {
		t.Errorf("期望姓名: 李四, 实际: %s", user.Name)
	}

	if user.Age != 30 {
		t.Errorf("期望年龄: 30, 实际: %d", user.Age)
	}

	if user.Email != "lisi@example.com" {
		t.Errorf("期望邮箱: lisi@example.com, 实际: %s", user.Email)
	}
}

// TestSerializer_MarshalAndUnmarshal 测试序列化和反序列化组合
func TestSerializer_MarshalAndUnmarshal(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	serializer := NewSerializer()

	originalUser := User{
		Name:  "王五",
		Age:   28,
		Email: "wangwu@example.com",
	}

	// 序列化
	data, err := serializer.Marshal(originalUser)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	t.Logf("序列化结果: %s", string(data))

	// 反序列化
	var unmarshaledUser User
	err = serializer.Unmarshal(data, &unmarshaledUser)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	// 比较原始数据和解序列化后的数据
	if !reflect.DeepEqual(originalUser, unmarshaledUser) {
		t.Errorf("原始数据和解序列化后的数据不一致\n原始: %+v\n反序列化: %+v", originalUser, unmarshaledUser)
	}
}
