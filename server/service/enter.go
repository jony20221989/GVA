package service

type ServiceGroup struct {
	BaseService
	UserService
	JwtService
	InitDBService
}

var ServiceGroupApp = new(ServiceGroup)
