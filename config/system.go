package config

type System struct {
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`                        // 端口
	Node        int64  `mapstructure:"node" json:"node" yaml:"node"`                        // 节点
	DbType      string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	ContextPath string `mapstructure:"context-path" json:"contextPath" yaml:"context-path"` // 请求上下文路径

}
