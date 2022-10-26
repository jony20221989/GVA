package service

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
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

func (userService *UserService) Register(u entity.SysUser) (userInter entity.SysUser, err error) {
	var user entity.SysUser
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return u, err
}

func (userService *UserService) ChangePassword(u *entity.SysUser, newPassword string) (userInter *entity.SysUser, err error) {
	var user entity.SysUser
	if err = global.DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.DB.Save(&user).Error
	return &user, err
}

func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
	assignErr := global.DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&entity.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.DB.Where("id = ?", id).First(&entity.SysUser{}).Update("authority_id", authorityId).Error
	return err
}
func (userService *UserService) DeleteUser(id int) (err error) {
	var user entity.SysUser
	err = global.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.DB.Delete(&[]entity.SysUserAuthority{}, "sys_user_id = ?", id).Error
	return err
}
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]entity.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []entity.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, entity.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&entity.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (userService *UserService) SetUserInfo(req entity.SysUser) error {
	return global.DB.Updates(&req).Error
}

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.DB.Model(&entity.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}
