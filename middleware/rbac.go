package middleware

import (
	"aixinge/api/model/common/response"
	"aixinge/api/model/system/request"
	"github.com/gofiber/fiber/v2"
)

func RbacHandler() fiber.Handler {
	// rbac 权限处理
	return func(c *fiber.Ctx) error {
		var success = false
		claims := c.Locals("claims")
		waitUse := claims.(*request.TokenClaims)
		uid := waitUse.ID
		if uid == 1 {
			// 管理员
			success = true
		} else {
			// 获取请求的URI
			url := c.OriginalURL()
			// 获取请求方法
			method := c.Method()
			// 未来要做 RBAC 权限认证
			print("获取请求的URI = " + url + ", method= " + method)
		}
		if success {
			return c.Next()
		} else {
			return response.FailWithDetailed(fiber.Map{}, "权限不足", c)
		}
	}
}
