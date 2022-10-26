import axios from 'axios'

axios.defaults.headers['Content-Type'] = 'application/json;charset=utf-8'

// 创建axios实例
const request = axios.create({
    // axios中请求配置有baseURL选项，表示请求URL公共部分
    baseURL: "http://localhost:8765",//process.env.VUE_APP_BASE_API,
    // 超时
    timeout: 10000
})

// request拦截器
request.interceptors.request.use(config => {


    return config
}, error => {
    console.log(error)
    Promise.reject(error)
})

// 响应拦截器
request.interceptors.response.use(res => {
    // 未设置状态码则默认成功状态

    return res.data
})

export default request