package main

import (
	"aixinge/core"
	"aixinge/global"
	"aixinge/initialize"
	"aixinge/utils"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title AiXinGe API
// @version 1.0.0
// @description artificial intelligence message push service
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.VP = core.Viper()      // 初始化Viper
	global.LOG = core.Zap()       // 初始化zap日志库
	global.DB = initialize.Gorm() // gorm连接数据库
	utils.Open()                  // 打开首页
	core.RunServer()
}
