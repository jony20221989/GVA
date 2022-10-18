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
