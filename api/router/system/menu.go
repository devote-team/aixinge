package system

import (
	"aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(router fiber.Router) (R fiber.Router) {
	menuRouter := router.Group("menu")
	var menuApi = v1.AppApi.SystemApi.Menu
	{
		menuRouter.Post("create", menuApi.Create)      // 创建
		menuRouter.Post("delete", menuApi.Delete)      // 删除
		menuRouter.Post("update", menuApi.Update)      // 更新
		menuRouter.Post("get", menuApi.Get)            // 根据id获取
		menuRouter.Post("page", menuApi.Page)          // 分页获取列表
		menuRouter.Post("auth", menuApi.Auth)          // 授权菜单
		menuRouter.Post("list", menuApi.List)          // 列表
		menuRouter.Post("list-tree", menuApi.ListTree) // 列表树
	}
	return menuRouter
}
