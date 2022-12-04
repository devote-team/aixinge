package system

import (
	"aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type FilesRouter struct {
}

func (s *FilesRouter) InitFileRouter(router fiber.Router) {
	fileRouter := router.Group("file")
	var fileApi = v1.AppApi.SystemApi.File
	{
		fileRouter.Post("upload", fileApi.Upload)    // 上传，记录操作日志
		fileRouter.Get("download", fileApi.Download) // 下载
	}
}
