package v1

import (
	"aixinge/api/v1/message"
	"aixinge/api/v1/system"
)

type ApiGroup struct {
	SystemApi  system.ApiGroup
	MessageApi message.ApiGroup
}

var AppApi = new(ApiGroup)
