package response

import (
	"server/model/entity"
)

type CaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength""`
}

type SysUserResponse struct {
	User entity.SysUser `json:"user"`
}

type LoginResponse struct {
	User      entity.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type SysAPIResponse struct {
	Api entity.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []entity.SysApi `json:"apis"`
}
type SysAuthorityResponse struct {
	Authority entity.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      entity.SysAuthority `json:"authority"`
	OldAuthorityId uint                `json:"oldAuthorityId"` // 旧角色ID
}

type SysMenusResponse struct {
	Menus []entity.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []entity.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu entity.SysBaseMenu `json:"menu"`
}
