package token

import "errors"

// Config Token配置
type Config struct {
	// Secret 密钥
	Secret string `json:"secret" mapstructure:"secret" yaml:"secret"`
	// Issuer 签发者
	Issuer string `json:"issuer" mapstructure:"issuer" yaml:"issuer"`
	// Duration 过期时间（小时）
	Duration int `json:"duration" mapstructure:"duration" yaml:"duration"`
	// RefreshDuration 刷新时间（小时）
	RefreshDuration int `json:"refresh_duration" mapstructure:"refresh_duration" yaml:"refresh_duration"`
}

// validate 验证配置
func (c *Config) validate() error {
	if c.Secret == "" {
		return errors.New("密钥不能为空")
	}
	if c.Issuer == "" {
		return errors.New("签发者不能为空")
	}
	if c.Duration <= 0 {
		return errors.New("过期时间必须大于0")
	}
	return nil
}
