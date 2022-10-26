const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,


  devServer:{
    host:'0.0.0.0',
    port:'8080',
    // https:false,
    open:false,
    //以上的ip和端口是我们本机的;下面为需要跨域的
    proxy:{ //配置跨域
      '/api':{
        target:'http://localhost:8765',
        ws:true,
        changeOrigin:true,//允许跨域
        pathRewrite:{
          '^/api':''   //请求的时候使用这个api就可以
        }
      }
    }
  }

})
