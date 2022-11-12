package message

// 消息类型
const (
	Sms          = 1 // 短信
	Mail         = 2 // 邮件
	Inbox        = 3 // 站内信
	WebSocket    = 4 // WebSocket
	MiniPrograms = 5 // 小程序
)

// 消息服务商
const (
	AliCloud     = 1 // 阿里云
	TencentCloud = 2 // 腾讯云
	BaiduCloud   = 3 // 百度云
	HuaweiCloud  = 4 // 华为云
)
