package service

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"server/global"
	"server/model/entity"
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
