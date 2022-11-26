package system

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/common/response"
	"aixinge/api/model/system"
	systemReq "aixinge/api/model/system/request"
	systemRes "aixinge/api/model/system/response"
	"aixinge/api/model/validation"
	"aixinge/global"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Role struct {
}

// Create
// @Tags Role
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Role true "创建角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/create [post]
func (b *Role) Create(c *fiber.Ctx) error {
	var role system.Role
	_ = c.BodyParser(&role)

	if err := roleService.Create(role); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		return response.FailWithMessage("创建失败", c)
	}

	return response.OkWithMessage("创建成功", c)
}

// Delete
// @Tags Role
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/role/delete [post]
func (b *Role) Delete(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := roleService.Delete(idsReq); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags Role
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Role true "角色信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /v1/role/update [post]
func (b *Role) Update(c *fiber.Ctx) error {
	var role system.Role
	_ = c.BodyParser(&role)
	if err := validation.Verify(role, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	err, role := roleService.Update(role)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败"+err.Error(), c)
	}

	return response.OkWithDetailed(systemRes.RoleResponse{Role: role}, "更新成功", c)
}

// AssignUser
// @Tags Role
// @Summary 角色分配用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.RoleUserParams true "角色ID、用户ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/assign-user [post]
func (b *Role) AssignUser(c *fiber.Ctx) error {
	var params systemReq.RoleUserParams
	_ = c.BodyParser(&params)
	if err := roleService.AssignUser(params); err != nil {
		global.LOG.Error("角色分配用户失败!", zap.Any("err", err))
		return response.FailWithMessage(err.Error(), c)
	}
	return response.OkWithMessage("获取成功", c)
}

// SelectedUsers
// @Tags Role
// @Summary 根据id获取角色已分配的用户ID列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/selected-users [post]
func (b *Role) SelectedUsers(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, userIds := roleService.SelectedUsers(idInfo.ID); err != nil {
		global.LOG.Error("获取角色分配用户ID列表失败", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(userIds, "获取成功", c)
	}
}

// AssignMenu
// @Tags Role
// @Summary 角色分配菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.RoleMenuParams true "角色ID、菜单ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/assign-menu [post]
func (b *Role) AssignMenu(c *fiber.Ctx) error {
	var params systemReq.RoleMenuParams
	_ = c.BodyParser(&params)
	if err := roleService.AssignMenu(params); err != nil {
		global.LOG.Error("角色分配菜单失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	}
	return response.OkWithMessage("获取成功", c)
}

// SelectedMenus
// @Tags Role
// @Summary 根据id获取角色已分配的菜单ID列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/selected-menus [post]
func (b *Role) SelectedMenus(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, menuIds := roleService.SelectedMenus(idInfo.ID); err != nil {
		global.LOG.Error("获取角色分配菜单ID列表失败", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(menuIds, "获取成功", c)
	}
}

// SelectedMenus
// @Tags Role
// @Summary 根据id获取角色已分配的菜单详细信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/selected-menus [post]
func (b *Role) SelectedMenusDetail(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, menuList := roleService.SelectedMenusDetail(idInfo.ID); err != nil {
		global.LOG.Error("获取角色分配菜单ID列表失败", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(menuList, "获取成功", c)
	}
}

// Get
// @Tags Role
// @Summary 根据id获取角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/get [post]
func (b *Role) Get(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, role := roleService.GetById(idInfo.ID); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.RoleResponse{Role: role}, "获取成功", c)
	}
}

// Get
// @Tags Role
// @Summary 批量根据id获取角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "角色ID列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/batch-get [post]
func (b *Role) BatchGet(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list := roleService.GetByIds(idsReq); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(list, "获取成功", c)
	}
}

// Page
// @Tags Role
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/page [post]
func (b *Role) Page(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := validation.Verify(pageInfo, validation.PageInfo); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list, total := roleService.Page(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// List
// @Tags Role
// @Summary 获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/role/list [post]
func (b *Role) List(c *fiber.Ctx) error {
	if err, list := roleService.List(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(list, "获取成功", c)
	}
}
