package v1

import "server/service"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.UserService
	jwtService  = service.ServiceGroupApp.JwtService
)
