package system

import (
	"aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type RoleRouter struct {
}

func (s *RoleRouter) InitRoleRouter(router fiber.Router) {
	roleRouter := router.Group("role")
	var roleApi = v1.AppApi.SystemApi.Role
	{
		roleRouter.Post("create", roleApi.Create)                // 创建
		roleRouter.Post("delete", roleApi.Delete)                // 删除
		roleRouter.Post("update", roleApi.Update)                // 更新
		roleRouter.Post("assign-user", roleApi.AssignUser)       // 角色分配用户
		roleRouter.Post("selected-users", roleApi.SelectedUsers) // 角色已分配用户ID列表
		roleRouter.Post("assign-menu", roleApi.AssignMenu)       // 角色分配菜单
		roleRouter.Post("selected-menus", roleApi.SelectedMenus) // 角色已分配菜单ID列表
		roleRouter.Post("get", roleApi.Get)                      // 根据id获取角色
		roleRouter.Post("batch-get", roleApi.BatchGet)           // 批量根据id获取角色
		roleRouter.Post("page", roleApi.Page)                    // 分页获取角色列表
		roleRouter.Post("list", roleApi.List)                    // 获取角色列表
	}
}
