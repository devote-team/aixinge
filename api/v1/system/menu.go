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

type Menu struct {
}

// Create
// @Tags Menu
// @Summary 创建菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Menu true "创建菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/menu/create [post]
func (b *Menu) Create(c *fiber.Ctx) error {
	var menu system.Menu
	_ = c.BodyParser(&menu)
	err := menuService.Create(menu)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Any("err", err))
		return response.FailWithMessage("创建失败"+err.Error(), c)
	}
	return response.OkWithMessage("创建成功", c)
}

// Delete
// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/menu/delete [post]
func (b *Menu) Delete(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := menuService.Delete(idsReq); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags Menu
// @Summary 更新菜单信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Menu true "菜单信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /v1/menu/update [post]
func (b *Menu) Update(c *fiber.Ctx) error {
	var menu system.Menu
	_ = c.BodyParser(&menu)
	if err := validation.Verify(menu, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	err, menu := menuService.Update(menu)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败"+err.Error(), c)
	}

	return response.OkWithDetailed(systemRes.MenuResponse{Menu: menu}, "更新成功", c)
}

// Get
// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/menu/get [post]
func (b *Menu) Get(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, menu := menuService.GetById(idInfo.ID); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.MenuResponse{Menu: menu}, "获取成功", c)
	}
}

// Page
// @Tags Menu
// @Summary 分页获取菜单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/menu/page [post]
func (b *Menu) Page(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := validation.Verify(pageInfo, validation.PageInfo); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list, total := menuService.Page(pageInfo); err != nil {
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
// @Tags Menu
// @Summary 获取菜单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.MenuParams true "查询参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/menu/list [post]
func (b *Menu) List(c *fiber.Ctx) error {
	var params systemReq.MenuParams
	_ = c.BodyParser(&params)
	if err, list := menuService.List(params); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(list, "获取成功", c)
	}
}

// Auth
// @Tags Menu
// @Summary 获取当前登录用户授权菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/menu/auth [post]
func (b *Menu) Auth(c *fiber.Ctx) error {
	if err, list := menuService.AuthList(systemReq.GetUserInfo(c)); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(list, "获取成功", c)
	}
}

// ListTree
// @Tags Menu
// @Summary 获取菜单列表树
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/menu/list-tree [post]
func (b *Menu) ListTree(c *fiber.Ctx) error {
	if err, listTree := menuService.ListTree(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(listTree, "获取成功", c)
	}
}
