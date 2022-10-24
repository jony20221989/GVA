package service

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/utils"
)

type UserService struct{}

func (userService *UserService) Login(u *entity.SysUser) (userInter *entity.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}

	var user entity.SysUser
	err = global.DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		//MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user entity.SysUser, err error) {
	var reqUser entity.SysUser
	err = global.DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	//MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	desc := info.Desc
	db := global.DB.Model(&entity.SysUser{})
	var userList []entity.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if desc {
		err = db.Order("id desc").Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	} else {
		err = db.Order("id asc").Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	}
	return userList, total, err
}
