package system

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/common/response"
	"aixinge/api/model/system"
	systemReq "aixinge/api/model/system/request"
	systemRes "aixinge/api/model/system/response"
	"aixinge/api/model/validation"
	"aixinge/global"
	"aixinge/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type User struct {
}

// Login
// @Tags Base
// @Summary 用户登录
// @accept application/json
// @Produce application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /v1/login [post]
func (b *User) Login(c *fiber.Ctx) error {
	var l systemReq.Login
	_ = c.BodyParser(&l)
	if err := validation.Verify(l, validation.Login); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	u := &system.User{Username: l.Username, Password: l.Password}
	if err, user := userService.Login(u); err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		return response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		return b.tokenNext(c, *user)
	}
}

// 登录以后签发jwt
func (b *User) tokenNext(c *fiber.Ctx, user system.User) error {
	j := &middleware.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	var expiresTime = global.CONFIG.JWT.ExpiresTime
	claims := systemReq.TokenClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresTime) * time.Minute)), // 过期时间 30分钟 配置文件
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                               // 签发时间
		},
	}
	accessToken, err := j.CreateToken(claims)
	refreshToken, rtErr := j.CreateToken(systemReq.RefreshTokenClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)), // 过期时间 1小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                   // 签发时间
		},
	})
	if err != nil || rtErr != nil {
		global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
		return response.FailWithMessage("设置登录状态失败", c)
	}
	return response.OkWithDetailed(systemRes.LoginResponse{
		User:         user,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, "登录成功", c)
}

// RefreshToken
// @Tags Base
// @Summary 刷新Token
// @accept application/json
// @Produce application/json
// @Param data body systemReq.RefreshToken true "刷新票据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"刷新Token成功"}"
// @Router /v1/refresh-token [post]
func (b *User) RefreshToken(c *fiber.Ctx) error {
	var rt systemReq.RefreshToken
	err := c.BodyParser(&rt)
	if err != nil || rt.RefreshToken == "" {
		return response.FailWithMessage("未登录或非法访问", c)
	}
	j := middleware.NewJWT()
	claims, err := j.ParseToken(rt.RefreshToken)
	if err != nil {
		return response.Result(response.ExpireRefreshToken, fiber.Map{"reload": true}, err.Error(), c)
	}
	if err, user := userService.GetById(claims.ID); err != nil {
		return response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		return b.tokenNext(c, user)
	}
}

// Create
// @Tags User
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.UserCreate true "创建User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/user/create [post]
func (b *User) Create(c *fiber.Ctx) error {
	var uc systemReq.UserCreate
	_ = c.BodyParser(&uc)
	if err := validation.Verify(uc, validation.UserCreate); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	err, userReturn := userService.Create(uc)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Any("err", err))
		return response.FailWithDetailed(systemRes.UserResponse{User: userReturn}, "注册失败", c)
	} else {
		return response.OkWithDetailed(systemRes.UserResponse{User: userReturn}, "注册成功", c)
	}
}

// Delete
// @Tags User
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/user/delete [post]
func (b *User) Delete(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := userService.Delete(idsReq); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags User
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.User true "用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /v1/user/update [post]
func (b *User) Update(c *fiber.Ctx) error {
	var user system.User
	_ = c.BodyParser(&user)
	if err := validation.Verify(user, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	err, user := userService.Update(user)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败"+err.Error(), c)
	}

	return response.OkWithDetailed(systemRes.UserResponse{User: user}, "更新成功", c)
}

// ChangePassword
//
// @Tags User
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /v1/user/change-password [post]
func (b *User) ChangePassword(c *fiber.Ctx) error {
	var user systemReq.ChangePasswordStruct
	_ = c.BodyParser(&user)
	if err := validation.Verify(user, validation.ChangePassword); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	u := &system.User{Username: user.Username, Password: user.Password}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.LOG.Error("修改失败!", zap.Any("err", err))
		return response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		return response.OkWithMessage("修改成功", c)
	}
}

// AssignRole
// @Tags User
// @Summary 用户分配角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.UserRoleParams true "用户ID、角色ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/user/assign-role [post]
func (b *User) AssignRole(c *fiber.Ctx) error {
	var params systemReq.UserRoleParams
	_ = c.BodyParser(&params)
	if err := userService.AssignRole(params); err != nil {
		global.LOG.Error("角色分配菜单失败", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	}
	return response.OkWithMessage("获取成功", c)
}

// SelectedRoles
// @Tags User
// @Summary 根据id获取用户已分配的角色ID列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/user/selected-roles [post]
func (b *User) SelectedRoles(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, roleIds := userService.SelectedRoles(idInfo.ID); err != nil {
		global.LOG.Error("获取用户分配角色ID列表失败", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(roleIds, "获取成功", c)
	}
}

// Get
// @Tags User
// @Summary 根据id获取用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/user/get [post]
func (b *User) Get(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, user := userService.GetById(idInfo.ID); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.UserResponse{User: user}, "获取成功", c)
	}
}

// Page
// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/user/page [post]
func (b *User) Page(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := validation.Verify(pageInfo, validation.PageInfo); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list, total := userService.Page(pageInfo); err != nil {
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
// @Tags User
// @Summary 获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/user/list [post]
func (b *User) List(c *fiber.Ctx) error {
	if err, list := userService.List(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(list, "获取成功", c)
	}
}
