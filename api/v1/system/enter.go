package system

import "aixinge/api/service"

type ApiGroup struct {
	User
	Role
	Menu
}

var menuService = service.AppService.SystemService.MenuService
var userService = service.AppService.SystemService.UserService
var roleService = service.AppService.SystemService.RoleService
