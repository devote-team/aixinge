package message

import "aixinge/api/service"

type ApiGroup struct {
	Application
}

var applicationService = service.AppService.MessageService.ApplicationService
