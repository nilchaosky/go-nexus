package snowflake

// Config Snowflake配置结构体
type Config struct {
	DatacenterID int64 `json:"datacenter_id" mapstructure:"datacenter_id" yaml:"datacenter_id"` // 数据中心ID，范围0-31
	WorkerID     int64 `json:"worker_id" mapstructure:"worker_id" yaml:"worker_id"`             // 工作节点ID，范围0-31
}

// Default 返回默认配置
func Default() Config {
	return Config{
		DatacenterID: 0,
		WorkerID:     0,
	}
}

// NewConfig 创建配置
func NewConfig(datacenterID, workerID int64) Config {
	return Config{
		DatacenterID: datacenterID,
		WorkerID:     workerID,
	}
}
