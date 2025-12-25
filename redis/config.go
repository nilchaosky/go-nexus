package redis

// Config 配置结构体
type Config struct {
	Address  string `json:"address" mapstructure:"address" yaml:"address"`    // 地址（默认：localhost）
	Port     int    `json:"port" mapstructure:"port" yaml:"port"`             // 端口（默认：6379）
	Password string `json:"password" mapstructure:"password" yaml:"password"` // 密码（默认：空）
	DB       []int  `json:"db" mapstructure:"db" yaml:"db"`                   // 数据库编号列表（默认：[0]）
}
