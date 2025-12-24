package snowflake

import (
	"testing"
)

// TestGenerateID_Single 测试生成单个ID功能
func TestGenerateID_Single(t *testing.T) {
	// 测试生成ID
	id := GenerateID()

	// 输出生成的ID
	t.Logf("生成的ID: %d", id)

	// 验证ID不为0
	if id == 0 {
		t.Error("生成的ID不能为0")
	}
}

// TestGenerateID_Multiple 测试生成多个ID功能
func TestGenerateID_Multiple(t *testing.T) {
	// 测试多次生成ID，验证不重复
	idMap := make(map[int64]bool, 100)
	count := 100

	for i := 0; i < count; i++ {
		newID := GenerateID()

		// 输出生成的ID
		t.Logf("第%d个生成的ID: %d", i+1, newID)

		// 验证ID不重复
		if idMap[newID] {
			t.Errorf("生成的ID重复: %d", newID)
		}
		idMap[newID] = true

		// 验证ID不为0
		if newID == 0 {
			t.Error("生成的ID不能为0")
		}
	}

	t.Logf("成功生成%d个不重复的ID", count)
}
