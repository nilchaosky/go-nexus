package snowflake

import (
	"errors"
	"sync"

	"github.com/GUAIK-ORG/go-snowflake/snowflake"
)

var (
	sf      *snowflake.Snowflake
	sfOnce  sync.Once
	initErr error
)

// init 包加载时自动使用默认配置初始化Snowflake节点
func init() {
	config := Default()
	err := Register(config)
	if err != nil {
		panic("Snowflake自动初始化失败: " + err.Error())
	}
}

// Register 使用自定义配置注册Snowflake节点
// config 为Snowflake配置
func Register(config Config) error {
	initErr = nil
	sfOnce.Do(func() {
		var err error
		sf, err = snowflake.NewSnowflake(config.DatacenterID, config.WorkerID)
		if err != nil {
			initErr = errors.New("Snowflake节点初始化失败: " + err.Error())
		}
	})
	return initErr
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

// MustGenerateID 生成全局唯一ID，如果失败会panic
// 适用于已确保节点已初始化的场景
func MustGenerateID() int64 {
	id, err := GenerateID()
	if err != nil {
		panic(err)
	}
	return id
}
