package v1

import (
	"aixinge/api/v1/system"
)

type ApiGroup struct {
	SystemApi system.ApiGroup
}

var AppApi = new(ApiGroup)
