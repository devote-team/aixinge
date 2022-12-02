package message

import "aixinge/api/service"

type ApiGroup struct {
	Application
	Channel
	ChannelTemplate
	MailLog
	MailTemplate
}

var applicationService = service.AppService.MessageService.ApplicationService
var channelService = service.AppService.MessageService.ChannelService
var channelTemplateService = service.AppService.MessageService.ChannelTemplateService
var mailLogService = service.AppService.MessageService.MailLogService
var mailTemplateService = service.AppService.MessageService.MailTemplateService
