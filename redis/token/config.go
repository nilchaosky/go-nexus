package token

import "errors"

// Config Token配置
type Config struct {
	// Secret 密钥
	Secret string `json:"secret" mapstructure:"secret" yaml:"secret"`
	// Issuer 签发者
	Issuer string `json:"issuer" mapstructure:"issuer" yaml:"issuer"`
	// Expiration 过期时间（天）
	Expiration int `json:"expiration" mapstructure:"expiration" yaml:"expiration"`
	// RefreshTime 刷新时间（天）
	RefreshTime int `json:"refresh_time" mapstructure:"refresh_time" yaml:"refresh_time"`
}

// validate 验证配置
func (c *Config) validate() error {
	if c.Secret == "" {
		return errors.New("密钥不能为空")
	}
	if c.Issuer == "" {
		return errors.New("签发者不能为空")
	}
	if c.Expiration <= 0 {
		return errors.New("过期时间必须大于0")
	}
	if c.RefreshTime <= 0 {
		return errors.New("刷新时间必须大于0")
	}
	return nil
}
