package system

import "aixinge/api/service"

type ApiGroup struct {
	User
	Role
	Menu
	File
}

var menuService = service.AppService.SystemService.MenuService
var userService = service.AppService.SystemService.UserService
var roleService = service.AppService.SystemService.RoleService
var fileService = service.AppService.SystemService.FileService
