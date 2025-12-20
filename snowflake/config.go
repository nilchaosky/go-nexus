package snowflake

// Config Snowflake配置结构体
type Config struct {
	DatacenterID int64 `json:"datacenter_id"` // 数据中心ID，范围0-31
	WorkerID     int64 `json:"worker_id"`     // 工作节点ID，范围0-31
}

// Default 返回默认配置
// 默认数据中心ID为0，工作节点ID为0
func Default() Config {
	return Config{
		DatacenterID: 0,
		WorkerID:     0,
	}
}

// NewConfig 创建新的配置
// datacenterID 为数据中心ID，范围0-31
// workerID 为工作节点ID，范围0-31
func NewConfig(datacenterID, workerID int64) Config {
	return Config{
		DatacenterID: datacenterID,
		WorkerID:     workerID,
	}
}
