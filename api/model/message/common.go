package message

// MsgType 消息类型
type MsgType int

const (
	Sms          MsgType = 1 // 短信
	Mail         MsgType = 2 // 邮件
	Inbox        MsgType = 3 // 站内信
	WebSocket    MsgType = 4 // WebSocket
	MiniPrograms MsgType = 5 // 小程序
)

// ChannelProvider 消息服务商
type ChannelProvider int

const (
	AliCloud     ChannelProvider = 1 // 阿里云
	TencentCloud ChannelProvider = 2 // 腾讯云
	BaiduCloud   ChannelProvider = 3 // 百度云
	HuaweiCloud  ChannelProvider = 4 // 华为云
)
