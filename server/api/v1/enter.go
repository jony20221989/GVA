package v1

import "server/service"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	BaseApi
	DBApi
	UserApi
	JwtApi
	SystemApiApi
}

var (
	baseService   = service.ServiceGroupApp.BaseService
	userService   = service.ServiceGroupApp.UserService
	jwtService    = service.ServiceGroupApp.JwtService
	initDBService = service.ServiceGroupApp.InitDBService
	apiService    = service.ServiceGroupApp.ApiService
)
