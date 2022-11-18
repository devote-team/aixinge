package service

import (
	"aixinge/api/service/message"
	"aixinge/api/service/system"
)

type Service struct {
	SystemService  system.ServiceGroup
	MessageService message.ServiceGroup
}

var AppService = new(Service)
