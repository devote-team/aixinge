# AiXinGe

Ai（爱）Xin（信）Ge（鸽） - 智能消息推送平台致力于解决大家在集成消息推送时的各种难题，力求将消息通知简单化、统一化，实现推送"All
in One"的效果。

> 目前项目刚刚启动，欢迎感兴趣的小伙伴 Star 插眼，我们会在后续的更新中逐步去实现我们的目标特性。

## 目标特性

- 易使用：一个 SDK/API 即可实现不同类型的消息推送，再也不用对接各种消息推送 SDK 了
- 易管理：集成市面上绝大部分推送渠道，实现统一管理，如需更改渠道，只需要进行相应配置并绑定到对应的消息模板即可。
- 易部署：可通过二进制文件或者 Docker 镜像实现一键启动，同时有可视化的引导式配置简化大家的配置难度。
- 高性能：依托于 Go 语言的特性，全程通过 Pipeline + Async 为推送平台提供强劲性能。

## 功能规划

V1 版本功能规划：
![](https://gitee.com/aixinge/aixinge/raw/master/wiki/img/Feature-V1.png)

后续功能计划中（小程序消息、OA 消息、订阅号等）

## 技术栈

- Fiber
- GORM
- Viper
- Casbin

## 依赖升级

```
go get -u github.com/gofiber/fiber/v2@latest

go get -u all

go mod tidy
```

# 文档更新

```
swag init
```

> 用户名 `admin` 密码 `123456`
http://127.0.0.1:8888/swagger/index.html

# 打包

> 交叉编译打包命令

```
goreleaser --snapshot --skip-publish --rm-dist
```
