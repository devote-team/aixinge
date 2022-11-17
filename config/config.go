package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// 数据库
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	// 上传文件
	Upload Upload `mapstructure:"upload" json:"upload" yaml:"upload"`
}
