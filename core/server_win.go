// +build windows

package core

import "github.com/gofiber/fiber/v2"

type ServerImpl struct{}

func (*ServerImpl) ServeAsync(address string, app *fiber.App) error {
	return app.Listen(address)
}
