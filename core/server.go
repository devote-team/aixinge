package core

import (
	"aixinge/global"
	"aixinge/initialize"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Server interface {
	ServeAsync(string, *fiber.App) error
}

func RunServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)

	time.Sleep(10 * time.Microsecond)
	global.LOG.Error(newServer().ServeAsync(address, Router).Error())
}

func newServer() Server {
	return &ServerImpl{}
}
