package service

import (
	"errors"
	"server/model/request"
)

type InitDBService struct{}

// InitDB 创建数据库并初始化 总入口
func (initDBService *InitDBService) InitDB(conf request.InitDB) (err error) {

	return errors.New("fail")
}
