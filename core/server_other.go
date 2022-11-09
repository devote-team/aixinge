// +build !windows

package core

import (
	"github.com/cloudflare/tableflip"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerImpl struct{}

func (*ServerImpl) ServeAsync(address string, app *fiber.App) error {
	return gracefulUpgrade(address, app)
}

func gracefulUpgrade(address string, router *fiber.App) error {
	upg, err := tableflip.New(tableflip.Options{})
	if err != nil {
		return err
	}
	defer upg.Stop()

	// 监听系统的 SIGHUP 信号，以此信号触发进程重启
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)
		for range sig {
			// 核心的 Upgrade 调用
			err := upg.Upgrade()
			if err != nil {
				log.Println("Upgrade failed:", err)
			}
		}
	}()

	// 注意必须使用 upg.Listen 对端口进行监听
	ln, err := upg.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Can't listen:", err)
		return err
	}

	// 启动server
	go func() {
		err = router.Listener(ln)
		if err != nil {
			panic(err)
		}
	}()

	if err = upg.Ready(); err != nil {
		panic(err)
		return err
	}
	<-upg.Exit()

	// 给老进程的退出设置一个 30s 的超时时间，保证老进程的退出
	time.AfterFunc(30*time.Second, func() {
		log.Println("Graceful shutdown timed out")
		os.Exit(1)
	})
	return nil
}
