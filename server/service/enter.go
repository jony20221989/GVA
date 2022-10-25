package service

type ServiceGroup struct {
	BaseService
	UserService
	JwtService
	InitDBService
	CasbinService
	ApiService
	OperationRecordService
	MenuService
	BaseMenuService
}

var ServiceGroupApp = new(ServiceGroup)
