package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	Error              = -1 // 异常
	Success            = 0  // 正常
	ExpireToken        = 1  // 登录 Token 过期
	ExpireRefreshToken = 2  // 刷新 RefreshToken 过期
)

func Result(code int, data interface{}, msg string, c *fiber.Ctx) error {
	// 开始时间
	return c.JSON(Response{
		code,
		data,
		msg,
	})
}

func Ok(c *fiber.Ctx) error {
	return Result(Success, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *fiber.Ctx) error {
	return Result(Success, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *fiber.Ctx) error {
	return Result(Success, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *fiber.Ctx) error {
	return Result(Success, data, message, c)
}

func Fail(c *fiber.Ctx) error {
	return Result(Error, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *fiber.Ctx) error {
	return Result(Error, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *fiber.Ctx) error {
	return Result(Error, data, message, c)
}
