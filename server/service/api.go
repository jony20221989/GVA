package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"server/global"
	"server/model/entity"
	"server/model/request"
)

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api entity.SysApi) (err error) {
	if !errors.Is(global.DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&entity.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.DB.Create(&api).Error
}

func (apiService *ApiService) DeleteApi(api entity.SysApi) (err error) {
	var entity entity.SysApi
	err = global.DB.Where("id = ?", api.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                  // api记录不存在
		return err
	}
	err = global.DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	success := CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if !success {
		return errors.New(entity.Path + ":" + entity.Method + "casbin同步清理失败")
	}
	e := CasbinServiceApp.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

func (apiService *ApiService) GetAPIInfoList(api entity.SysApi, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := global.DB.Model(&entity.SysApi{})
	var apiList []entity.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't matched any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return apiList, total, err
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return apiList, total, err
}

func (apiService *ApiService) GetAllApis() (apis []entity.SysApi, err error) {
	err = global.DB.Find(&apis).Error
	return
}

func (apiService *ApiService) GetApiById(id int) (api entity.SysApi, err error) {
	err = global.DB.Where("id = ?", id).First(&api).Error
	return
}
func (apiService *ApiService) UpdateApi(api entity.SysApi) (err error) {
	var oldA entity.SysApi
	err = global.DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&entity.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.DB.Save(&api).Error
		}
	}
	return err
}

func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	var apis []entity.SysApi
	err = global.DB.Find(&apis, "id in ?", ids.Ids).Delete(&apis).Error
	if err != nil {
		return err
	} else {
		for _, sysApi := range apis {
			success := CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
			if !success {
				return errors.New(sysApi.Path + ":" + sysApi.Method + "casbin同步清理失败")
			}
		}
		e := CasbinServiceApp.Casbin()
		err = e.InvalidateCache()
		if err != nil {
			return err
		}
	}
	return err
}
