package service

type ServiceGroup struct {
	BaseService
	UserService
	JwtService
	InitDBService
	CasbinService
	ApiService
	OperationRecordService
}

var ServiceGroupApp = new(ServiceGroup)
