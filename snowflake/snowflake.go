package snowflake

import (
	"errors"

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

// Register 使用自定义配置注册Snowflake节点
// config 为Snowflake配置
func Register(config Config) error {
	var err error
	sf, err = snowflake.NewSnowflake(config.DatacenterID, config.WorkerID)
	if err != nil {
		return errors.New("Snowflake节点初始化失败: " + err.Error())
	}
	return nil
}

// GenerateID 生成全局唯一ID
// 返回int64类型的唯一ID
// 包加载时已自动初始化，通常不会返回错误
func GenerateID() (int64, error) {
	if sf == nil {
		return 0, errors.New("Snowflake节点未初始化")
	}
	return sf.NextVal(), nil
}

// GenerateSerializeInt64 生成全局唯一ID（SerializeInt64类型）
// 返回SerializeInt64类型的唯一ID，支持Gin框架的JSON序列化
// 包加载时已自动初始化，通常不会返回错误
func GenerateSerializeInt64() (variant.SerializeInt64, error) {
	if sf == nil {
		return 0, errors.New("Snowflake节点未初始化")
	}
	id, err := GenerateID()
	if err != nil {
		return 0, err
	}
	return variant.NewSerializeInt64(id), nil
}
