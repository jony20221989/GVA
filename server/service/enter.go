package service

type ServiceGroup struct {
	UserService
	JwtService
}

var ServiceGroupApp = new(ServiceGroup)
