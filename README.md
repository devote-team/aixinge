# AiXinGe

Ai（爱）Xin（信）Ge（鸽） - 智能消息推送平台

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
