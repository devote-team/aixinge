package core

import (
	"aixinge/global"
	"aixinge/initialize"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	ServeAsync(string, *fiber.App) error
}

func RunServer() {
	// init routers
	app := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)

	// kill daemon exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-quit
		fmt.Println("Shutdown Server ...")
		if err := app.Shutdown(); err != nil {
			fmt.Println(err)
			log.Fatalf("Server Shutdown: %s", err)
		}
		fmt.Println("Server exit")
	}()

	// start app
	time.Sleep(10 * time.Microsecond)
	global.LOG.Error(app.Listen(address).Error())
}
