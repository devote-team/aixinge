package system

import (
	"aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(router fiber.Router) {
	userRouter := router.Group("user")
	var userApi = v1.AppApi.SystemApi.User
	{
		userRouter.Post("create", userApi.Create)                  // 创建
		userRouter.Post("delete", userApi.Delete)                  // 删除
		userRouter.Post("update", userApi.Update)                  // 更新
		userRouter.Post("change-password", userApi.ChangePassword) // 修改密码
		userRouter.Post("assign-role", userApi.AssignRole)         // 用户分配角色
		userRouter.Post("selected-roles", userApi.SelectedRoles)   // 用户已分配角色ID列表
		userRouter.Post("get", userApi.Get)                        // 根据id获取用户
		userRouter.Post("page", userApi.Page)                      // 分页获取用户列表
		userRouter.Post("list", userApi.List)                      // 获取用户列表
	}
}
