package request

import (
	"aixinge/global"
	"github.com/gofiber/fiber/v2"
)

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *fiber.Ctx) *TokenClaims {
	if claims := c.Locals("claims"); claims == nil {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return nil
	} else {
		waitUse := claims.(*TokenClaims)
		return waitUse
	}
}
