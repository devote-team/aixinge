package message

import "aixinge/global"

type Application struct {
	global.MODEL
	Name      string `json:"name"`      // 应用名称
	AppKey    string `json:"appKey"`    // 应用 ID
	AppSecret string `json:"appSecret"` // 应用密钥
	Remark    string `json:"remark"`    // 备注
	Status    int    `json:"status"`    // 状态，1、正常 2、禁用
}
