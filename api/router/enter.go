package router

import (
	"aixinge/api/router/message"
	"aixinge/api/router/system"
)

type Router struct {
	System  system.RouterGroup
	Message message.RouterGroup
}

var AppRouter = new(Router)
