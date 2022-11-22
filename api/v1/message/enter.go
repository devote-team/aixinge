package message

import "aixinge/api/service"

type ApiGroup struct {
	Application
	Channel
	MailTemplate
}

var applicationService = service.AppService.MessageService.ApplicationService
var channelService = service.AppService.MessageService.ChannelService
var mailTemplateService = service.AppService.MessageService.MailTemplateService
