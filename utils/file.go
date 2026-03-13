package nexus_utils

import (
	"os"
)

// IsDir 判断路径是否为目录
// 如果路径是目录返回 true，如果是文件或不存在返回 false
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
