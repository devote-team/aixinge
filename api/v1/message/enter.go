package message

import "aixinge/api/service"

type ApiGroup struct {
	Application
	Channel
	ChannelTemplate
	MailTemplate
}

var applicationService = service.AppService.MessageService.ApplicationService
var channelService = service.AppService.MessageService.ChannelService
var channelTemplateService = service.AppService.MessageService.ChannelTemplateService
var mailTemplateService = service.AppService.MessageService.MailTemplateService
