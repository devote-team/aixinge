package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime uint   `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
}
