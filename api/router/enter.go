package router

import (
	"aixinge/api/router/system"
)

type Router struct {
	System system.RouterGroup
}

var AppRouter = new(Router)
