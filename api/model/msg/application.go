package msg

import "aixinge/global"

type Application struct {
	global.MODEL
	Name      string `json:"name"`      // 应用名称
	AppKey    string `json:"appKey"`    // 应用 ID
	AppSecret string `json:"appSecret"` // 应用密钥
	Remark    string `json:"Remark"`    // 备注
}
