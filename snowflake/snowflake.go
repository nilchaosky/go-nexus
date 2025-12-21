package snowflake

import (
	"errors"
	"fmt"

	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/nilchaosky/go-nexus/serialize/variant"
)

var (
	sf *snowflake.Snowflake
)

// init 包加载时自动使用默认配置初始化Snowflake节点
func init() {
	config := Default()
	var err error
	sf, err = snowflake.NewSnowflake(config.DatacenterID, config.WorkerID)
	if err != nil {
		panic("Snowflake自动初始化失败: " + err.Error())
	}
}

// Register 注册Snowflake节点
func Register(config Config) error {
	var err error
	sf, err = snowflake.NewSnowflake(config.DatacenterID, config.WorkerID)
	if err != nil {
		return fmt.Errorf("节点初始化失败: %w", err)
	}
	return nil
}

// GenerateID 生成全局唯一ID
func GenerateID() (int64, error) {
	if sf == nil {
		return 0, errors.New("节点未初始化")
	}
	return sf.NextVal(), nil
}

// GenerateSerializeInt64 生成SerializeInt64类型ID
func GenerateSerializeInt64() (variant.SerializeInt64, error) {
	if sf == nil {
		return 0, errors.New("节点未初始化")
	}
	id, err := GenerateID()
	if err != nil {
		return 0, err
	}
	return variant.NewSerializeInt64(id), nil
}
