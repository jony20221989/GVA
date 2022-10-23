package service

type ServiceGroup struct {
	BaseService
	UserService
	JwtService
	InitDBService
	CasbinService
}

var ServiceGroupApp = new(ServiceGroup)
