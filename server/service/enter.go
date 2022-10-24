package service

type ServiceGroup struct {
	BaseService
	UserService
	JwtService
	InitDBService
	CasbinService
	ApiService
}

var ServiceGroupApp = new(ServiceGroup)
