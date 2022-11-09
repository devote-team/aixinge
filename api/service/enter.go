package service

import (
	"aixinge/api/service/system"
)

type Service struct {
	SystemService system.ServiceGroup
}

var AppService = new(Service)
