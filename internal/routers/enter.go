package routers

import (
	"Go/internal/routers/manage"
	"Go/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)