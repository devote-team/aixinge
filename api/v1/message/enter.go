package message

import "aixinge/api/service"

type ApiGroup struct {
	Application
	MailTemplate
}

var applicationService = service.AppService.MessageService.ApplicationService
var mailTemplateService = service.AppService.MessageService.MailTemplateService
