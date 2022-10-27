import request from "@/utils/request";


// 获取验证码
export function getCaptcha() {
    return request({
        url: '/base/captcha',
     //   url: '/captchaImage',
        method: 'post',
        timeout: 10000
    })
}
