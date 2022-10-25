package request

import (
	"gorm.io/gorm"
	"server/model/entity"
)

type AddMenuAuthorityInfo struct {
	Menus       []entity.SysBaseMenu `json:"menus"`
	AuthorityId uint                 `json:"authorityId"` // 角色ID
}

func DefaultMenu() []entity.SysBaseMenu {
	return []entity.SysBaseMenu{{
		Model:     gorm.Model{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: entity.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
