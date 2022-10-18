package service

type ServiceGroup struct {
	UserService
	JwtService
	InitDBService
}

var ServiceGroupApp = new(ServiceGroup)
